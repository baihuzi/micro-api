package core

import (
	"fmt"
	"review-server/core/http/router"
)

func init()  {
	fmt.Println("load config.ini")
	loadConfig()
}

func Run() {
	fmt.Println("app run")
	InitDB()
	InitHttpServer()
}

func InitDB() {
	config := GetConfig()
	config.Mysql.Conn()
}

func InitHttpServer() {
	config := GetConfig()
	config.Http.Start(router.GetAll())
}
