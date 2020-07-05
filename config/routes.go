package config

import (
	"fmt"
	"review-server/controllers"
	"review-server/core/http/router"
)

func RegisterRouters() {
	fmt.Println("register routers")
	router.POST("/api/login", &controllers.LoginController{}, "Login", map[string]interface{}{})
}
