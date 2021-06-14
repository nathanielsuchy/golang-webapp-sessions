package controllers

import (
	"webapp/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

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
	} else {
		o := orm.NewOrm()
		var user models.User
		err := o.QueryTable("user").Filter("id", c.GetSession("user_id")).One(&user)
		if err == orm.ErrNoRows {
			// The user is signed into a deleted account. Destroy their session.
			c.DestroySession()
			flash := beego.NewFlash()
			flash.Warning("The account you were logged into no longer exists!")
			flash.Store(&c.Controller)
			c.Ctx.Redirect(302, "/sign-in")
		} else {
			if user.IsDisabled == true {
				// The user is signed into a disabled account. Destroy their session.
				c.DestroySession()
				flash := beego.NewFlash()
				flash.Warning("The account you were logged is disabled. You are being signed out!")
				flash.Store(&c.Controller)
				c.Ctx.Redirect(302, "/sign-in")
			}
		}
	}
}

// Runs after request function execution
func (c *PrivateController) Finish() {
	// Any cleanup logic common to all requests goes here. Logging or metrics, for example.
}
