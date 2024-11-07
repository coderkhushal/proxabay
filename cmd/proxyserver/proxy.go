package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"syscall"
	"time"

	service "github.com/coderkhushal/proxabay/cmd/services"
	"github.com/pterm/pterm"
)

type Proxy struct {
	HttpPort string
	Server   http.Server
	Origin   string
}

func NewProxy(origin string, port string) *Proxy {
	return &Proxy{
		Server:   http.Server{},
		HttpPort: port,
		Origin:   origin,
	}
}

func (p *Proxy) Start() error {
	url, err := url.Parse(p.Origin)
	if err != nil {
		pterm.Error.Println(err)
		return err
	}
	proxyhandler := httputil.NewSingleHostReverseProxy(url)
	proxyhandler.Director = func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.Host = url.Host
		req.Header.Set("User-Agent", "curl/18.0.1")
		req.Header.Set("Accept", "*/*")
	}

	p.Server.Addr = p.HttpPort

	p.Server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		existingcache, err := service.GetCacheForProxy(p.Origin+r.URL.String(), p.HttpPort)

		if err != nil {
			pterm.Error.Println(err)

		} else if existingcache.Origin == "" {
			pterm.Info.Println("Cache not found , hitting main server")
		} else {
			// write response from cache
			var headers http.Header

			json.Unmarshal(existingcache.Headers, &headers)

			for key, values := range headers {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}
			w.Header().Add("Cache", "hit")
			w.WriteHeader(existingcache.Status)
			w.Write(existingcache.Body)
			switch existingcache.Status {
			case 200:
				s := pterm.FgYellow.Sprintf("http://localhost%s -> %s", existingcache.Port+r.URL.String(), existingcache.Origin)

				pterm.FgLightGreen.Printf("%d Cache: HIT ", existingcache.Status)
				pterm.FgYellow.Printfln(s)
				break
			default:
				s := pterm.FgYellow.Sprintf("http://localhost%s -> %s", existingcache.Port+r.URL.String(), existingcache.Origin)

				pterm.FgLightBlue.Printf("%d Cache: HIT ", existingcache.Status)
				pterm.FgYellow.Printfln(s)
			}
			return
		}
		proxyhandler.ServeHTTP(w, r)
	})
	proxyhandler.ModifyResponse = func(r *http.Response) error {
		// var response map[string]interface{}
		responsejson, _ := io.ReadAll(r.Body)
		headerjson, _ := json.Marshal(r.Header) // Convert http.Header to []byte
		r.Body.Close()

		if err != nil {
			pterm.Error.Println(err)
			return err
		}
		err = service.CreateNewCache(r.Request.URL.String(), p.HttpPort, headerjson, responsejson, r.StatusCode)
		r.Body = io.NopCloser(bytes.NewBuffer(responsejson))

		switch r.StatusCode {
		case 200:

			s := pterm.FgYellow.Sprintf("http://localhost%s -> %s", p.HttpPort, r.Request.URL)
			pterm.FgLightGreen.Printf("%d Cache: MISS ", r.StatusCode)
			fmt.Println(s)

			break
		default:

			s := pterm.FgYellow.Sprintf("http://localhost%s -> %s", p.HttpPort, r.Request.URL)
			pterm.FgLightBlue.Printf("%d Cache: MISS ", r.StatusCode)
			fmt.Println(s)
		}
		return nil
	}
	errChan := make(chan error)

	go func() {
		// This runs the server in a goroutine
		log.Printf("Starting proxy on port %s\n", p.HttpPort)
		if err := p.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}

		close(errChan) // Close the channel when done
		service.Sigch <- syscall.SIGINT
	}()
	select {
	case err := <-errChan:
		if err != nil {
			pterm.Error.Println(err)
			return err
		}
		return nil
	case <-time.After(time.Second * 1):
		return nil
	}
}

func (p *Proxy) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := p.Server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil

}
