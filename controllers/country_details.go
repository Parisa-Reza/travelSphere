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
		c.Ctx.Output.SetStatus(404)
		return
	}

	c.Data["Country"] = detail
}