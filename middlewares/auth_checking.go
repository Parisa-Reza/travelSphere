package middlewares

import (
	beegoContext "github.com/beego/beego/v2/server/web/context"
)

// AuthCheckFilter checks if a user session exists. if not, it redirects them to the landing page.
func AuthCheckFilter(ctx *beegoContext.Context) {
	if ctx.Input.Session("user_id") == nil {
		ctx.Redirect(302, "/")
		return
	}
}
