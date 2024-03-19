package main

import (
	"secondkill/routers"
	"secondkill/service"
)

func main() {
	service.Mysqlinit()
	service.Redisinit()
	routers.InitRouters()
}
