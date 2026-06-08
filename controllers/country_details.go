package controllers

import (
	"travelSphere/services"

	"github.com/beego/beego/v2/server/web"
)

type CountryDetailController struct {
	web.Controller
}

func (c *CountryDetailController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "countryInfo.tpl"

	slug := c.Ctx.Input.Param(":slug")
	svc := &services.CountryDetailService{}

	detail, err := svc.FindBySlug(slug)
	if err != nil {

		// forcing Beego to stop rendering 'destination.tpl' and jump directly to  ErrorController's Error404 handler

		c.Abort("404")
		return
	}

	c.Layout = "layout.tpl"
	c.TplName = "countryInfo.tpl"
	c.Data["Country"] = detail
}
