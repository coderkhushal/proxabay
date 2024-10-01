package server

import (
	"fmt"
)

type ProxyManger struct {
	Proxies []*Proxy
}

func NewProxyManger() *ProxyManger {
	return &ProxyManger{}
}

func (m *ProxyManger) StartNewProxy(origin string, port string) error {
	newproxy := NewProxy(origin, port)
	err := newproxy.Start()
	if err != nil {
		fmt.Println(err)
		return err
	}
	m.Proxies = append(m.Proxies, newproxy)
	return nil

}

func (m *ProxyManger) StopAllProxies() error {
	var response error = nil
	for _, value := range m.Proxies {
		err := value.Stop()
		if err != nil {
			response = err
			fmt.Println(err)
		} else {
			fmt.Printf("Stopped Origin : %s , Port : %s ", value.Origin, value.HttpPort)

		}
	}
	return response
}
