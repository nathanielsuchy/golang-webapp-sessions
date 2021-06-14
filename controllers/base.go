package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

// Runs after Init before request function execution
func (c *BaseController) Prepare() {
	c.Data["IsSignedIn"] = (c.GetSession("email") != nil)
	c.Data["IsNotSignedIn"] = (c.GetSession("email") == nil)
}

// Runs after request function execution
func (c *BaseController) Finish() {
	// Any cleanup logic common to all requests goes here. Logging or metrics, for example.
}
