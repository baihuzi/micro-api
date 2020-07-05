package main

import (
	_ "review-server/config"
	"review-server/core"
)

func main() {
	core.Run()
}
