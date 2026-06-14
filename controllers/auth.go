package controllers

import (
	"travelSphere/services"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) LoginPost() {
	username := c.GetString("username")
	if username == "" {
		c.Redirect("/", 302)
		return
	}

	authService := &services.AuthService{}
	user, err := authService.AuthenticateUser(username)
	if err != nil {
		c.Redirect("/", 302)
		return
	}

	c.SetSession("user_id", user.ID)
	c.SetSession("username", user.Username)

	c.Redirect(c.Ctx.Input.Referer(), 302) //  back  to the active page view
}

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Redirect("/", 302)
}
