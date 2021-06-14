package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type SignoutController struct {
	BaseController
}

func (c *SignoutController) Get() {
	c.DestroySession()

	flash := beego.NewFlash()

	flash.Warning("You have been signed out!")
	flash.Store(&c.Controller)

	c.Ctx.Redirect(302, "/")
}
