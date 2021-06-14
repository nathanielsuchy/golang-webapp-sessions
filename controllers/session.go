package controllers

import (
	"webapp/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type SessionController struct {
	BaseController
}

type SessionFormSubmission struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (c *SessionController) Get() {
	if c.Data["IsSignedIn"] == false {
		c.XSRFExpire = 7200
		c.Data["xsrfdata"] = c.XSRFFormHTML()
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

	o := orm.NewOrm()

	var user models.User
	userErr := o.QueryTable("user").Filter("email", sessionFormSubmission.Email).One(&user)

	if userErr == orm.ErrNoRows {
		// Account does not exist throw an error!
		flash := beego.NewFlash()

		flash.Warning("The account you tried to access does not exist. Did you mean to create an account?")
		flash.Store(&c.Controller)

		c.Ctx.Redirect(302, "/sign-in")
	} else {
		// Account exists attempt to login
		if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(sessionFormSubmission.Password)) == nil {
			// If the credentials are correct start the session
			sess := c.StartSession()
			sess.Set(c.Ctx.Request.Context(), "user_id", user.Id)

			flash := beego.NewFlash()

			flash.Notice("Welcome back, " + user.Email + " You are now signed in!")
			flash.Store(&c.Controller)

			c.Ctx.Redirect(302, "/")
		} else {
			// If the credentials are incorrect show an error
			flash := beego.NewFlash()

			flash.Warning("The username and password you entered is incorrect. Please check the spelling and try again!")
			flash.Store(&c.Controller)

			c.Ctx.Redirect(302, "/sign-in")
		}
	}
}
