package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	web.Controller
}

func (c *ErrorController) Error404() {
	c.Layout = "layout.tpl"
	c.TplName = "404.tpl"
	c.Data["ErrorMessage"] = "The requested destination could not found"
}
