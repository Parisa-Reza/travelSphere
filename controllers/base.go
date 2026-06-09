package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

// // Prepare runs automatically before the specific GET/POST method executes
// func (c *BaseController) Prepare() {
// 	// Check if a username exists in the active session
// 	username := c.GetSession("username")

// 	if username != nil {
// 		c.Data["IsLoggedIn"] = true
// 		c.Data["CurrentUserName"] = username.(string)
// 	} else {
// 		c.Data["IsLoggedIn"] = false
// 	}
// }

func (c *BaseController) Prepare() {
	username := c.GetSession("username")
	userID := c.GetSession("user_id")

	if username != nil && userID != nil {
		c.Data["IsLoggedIn"] = true
		c.Data["CurrentUserName"] = username.(string)
		c.Data["CurrentUserID"] = userID.(string) // Now exposed inside your template scopes safely
	} else {
		c.Data["IsLoggedIn"] = false
	}
}
