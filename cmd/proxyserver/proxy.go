package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
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
	proxyhandler.ModifyResponse = func(r *http.Response) error {
		var response map[string]interface{}

		json.NewDecoder(r.Body).Decode(&response)

		responsejson, err := json.Marshal(response)

		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf(string(responsejson))

		return nil
	}
	p.Server.Addr = p.HttpPort
	p.Server.Handler = proxyhandler
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
