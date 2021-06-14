package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type SessionController struct {
	BaseController
}

type SessionFormSubmission struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (c *SessionController) Get() {

	v := c.GetSession("email")
	if v == nil {
		c.XSRFExpire = 7200
		c.Data["xsrfdata"] = c.XSRFFormHTML()

		c.Data["Website"] = "beego.me"
		c.Data["Email"] = "astaxie@gmail.com"
		c.Layout = "layout.tpl"
		c.TplName = "signin.tpl"
	} else {
		flash := beego.NewFlash()
		flash.Error("You are already signed in!")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/")
	}
}

func (c *SessionController) Post() {
	sessionFormSubmission := SessionFormSubmission{}

	if err := c.ParseForm(&sessionFormSubmission); err != nil {
		println(err.Error())
	}

	sess := c.StartSession()
	sess.Set(c.Ctx.Request.Context(), "email", sessionFormSubmission.Email)

	flash := beego.NewFlash()

	flash.Notice("You are now signed in!")
	flash.Store(&c.Controller)

	c.Ctx.Redirect(302, "/")
}
