package config

import (
	"fmt"
	servicesIns "review-server/core/services"
	"review-server/middlewares"
	"review-server/security"
)

func RegisterServices() {
	fmt.Println("register services")
	services := servicesIns.GetServices()

	services.HttpServices.RegisterSecurityFirewalls("request.sender", security.NewRequestSenderProvider(`[^\/api\/login(.*)]`))
	services.HttpServices.RegisterMiddleware(middlewares.NewViewMiddleware())

}
