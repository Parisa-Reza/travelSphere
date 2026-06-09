package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Layout = "layout.tpl"
	c.TplName = "404.tpl"
	c.Data["ErrorMessage"] = "The requested destination could not found"
}
