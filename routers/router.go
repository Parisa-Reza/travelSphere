package routers

import (
	"travelSphere/controllers"
	"travelSphere/controllers/api"

	"github.com/beego/beego/v2/server/web"
)

func Init() {
	// Home page route (ssr)
	web.Router("/", &controllers.HomeController{})

	// Countries page route (ssr)
	web.Router("/countries", &controllers.CountriesController{})

	//  Single Country Detail Page (ssr)
	web.Router("/countries/:slug", &controllers.CountryDetailController{})

	// API Endpoint for Country Search (json api)
	web.Router("/api/countries", &api.CountryAPIController{})

	// Error handling route
	web.ErrorController(&controllers.ErrorController{})
}
