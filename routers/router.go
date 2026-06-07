package routers

import (
	"travelSphere/controllers"
	"travelSphere/controllers/api"

	"github.com/beego/beego/v2/server/web"
)

func Init() {
	// Home page route
	web.Router("/", &controllers.HomeController{})

	// Countries page route
	web.Router("/countries", &controllers.CountriesController{})

	// API Endpoint for Country Search
	web.Router("/api/countries", &api.CountryAPIController{})
}
