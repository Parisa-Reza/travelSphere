package routers

import (
	"travelSphere/controllers"
	"travelSphere/controllers/api"
	"travelSphere/middlewares"

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

	// Simple Auth Handlers
	web.Router("/login", &controllers.AuthController{}, "post:LoginPost")
	web.Router("/logout", &controllers.AuthController{}, "get:Logout")

	// Enable reading Raw request body parameters for JSON APIs
	web.BConfig.CopyRequestBody = true

	// whishlist page route
	web.Router("/wishlist", &controllers.WishlistController{})

	// Wishlist APis
	web.Router("/api/wishlist", &controllers.WishlistController{}, "get:GetAPI;post:PostAPI")
	web.Router("/api/wishlist/:id", &controllers.WishlistController{}, "put:PutAPI;delete:DeleteAPI")

	web.InsertFilter("/wishlist", web.BeforeRouter, middlewares.AuthCheckFilter)
	web.InsertFilter("/api/wishlist", web.BeforeRouter, middlewares.AuthCheckFilter)
	web.InsertFilter("/api/wishlist/*", web.BeforeRouter, middlewares.AuthCheckFilter)

}
