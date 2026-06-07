package main

import (
	"travelSphere/middlewares/filters"
	"travelSphere/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// Initialize routes
	routers.Init()

	// registering global request logging filter for all routes
	web.InsertFilter("/*", web.BeforeRouter, filters.RequestLogger())

	// Start the Beego server
	web.Run()
}
