package services

import (
	"review-server/core/http/middlewares"
	"review-server/core/http/security"
)

type HttpServices struct {
	//Routers     []router.Route
	Middlewares []middlewares.MiddlewareInterface
	Security    *security.Security
	//RB          *request.RequestBodyInterface
}

func NewHttpServices() *HttpServices {
	return &HttpServices{
		Middlewares: make([]middlewares.MiddlewareInterface, 0),
		Security: &security.Security{
			Firewalls:            map[string]security.FirewallInterface{},
			AccessTokenGenerator: map[string]security.AccessTokenGeneratorInterface{},
		},
	}
}
func (h *HttpServices) RegisterMiddleware(m middlewares.MiddlewareInterface) {
	h.Middlewares = append(h.Middlewares, m)
}

func (h *HttpServices) RegisterSecurityFirewalls(key string, f security.FirewallInterface) {
	h.Security.Firewalls[key] = f
}

func (h *HttpServices) RegisterSecurityAccessTokenGenerator(key string, a security.AccessTokenGeneratorInterface) {
	h.Security.AccessTokenGenerator[key] = a
}
