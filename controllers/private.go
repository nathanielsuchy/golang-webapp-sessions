package controllers

import beego "github.com/beego/beego/v2/server/web"

type PrivateController struct {
	BaseController
}

// Runs after Init before request function execution
func (c *PrivateController) Prepare() {
	if c.GetSession("user_id") == nil {
		flash := beego.NewFlash()
		flash.Warning("You must be signed in to view this page!")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/sign-in")
	}
}

// Runs after request function execution
func (c *PrivateController) Finish() {
	// Any cleanup logic common to all requests goes here. Logging or metrics, for example.
}
