package main

import (
	"travelSphere/middlewares"
	"travelSphere/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Initialize routes
	routers.Init()

	// registering logging filter for all routes
	web.InsertFilter("/*", web.BeforeRouter, middlewares.RequestLogger())

	// Start the Beego server
	web.Run()
}
