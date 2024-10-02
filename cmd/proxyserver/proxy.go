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
	"time"

	service "github.com/coderkhushal/proxabay/cmd/services"
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
		log.Fatal(err)
		return err
	}
	proxyhandler := httputil.NewSingleHostReverseProxy(url)

	// proxyhandler.ModifyResponse = func(r *http.Response) error {
	// }
	p.Server.Addr = p.HttpPort

	p.Server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		existingcache, err := service.GetCacheForProxy(p.Origin, p.HttpPort)
		if err != nil {
			fmt.Println("some error occured while fetching cache")

		} else if existingcache.Origin == "" {
			fmt.Println("Cache not found , hitting main server")
		} else {
			// write response from cache
			var headers http.Header
			json.Unmarshal(existingcache.Headers, headers)
			for key, values := range headers {
				for _, value := range values {
					w.Header().Add(key, value)
				}
			}
			w.Header().Add("Cache", "hit")
			w.Write(existingcache.Body)
			return
		}
		proxyhandler.ServeHTTP(w, r)
	})
	proxyhandler.ModifyResponse = func(r *http.Response) error {
		// var response map[string]interface{}
		responsejson, _ := io.ReadAll(r.Body)
		headerjson, _ := json.Marshal(r.Request.Header) // Convert http.Header to []byte
		r.Body.Close()

		if err != nil {
			fmt.Println(err)
			return err
		}
		err = service.CreateNewCache(p.Origin, p.HttpPort, headerjson, responsejson)
		r.Body = io.NopCloser(bytes.NewBuffer(responsejson))
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
	}()
	select {
	case err := <-errChan:
		if err != nil {
			fmt.Println(err)
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
