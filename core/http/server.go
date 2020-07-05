package http

import (
	"fmt"
	"log"
	"net/http"
	"review-server/core/http/router"
)

type Http struct {
	host string
	port string
}

func (h *Http) Init(host, port string) {
	h.host = host
	h.port = port
}

func (h *Http) Start(routes []*router.Route) {
	for _, route := range routes {
		http.HandleFunc(route.Path, route.Handler)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", h.host, h.port), nil))
}
