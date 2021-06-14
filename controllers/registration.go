package controllers

import (
	"fmt"
	"webapp/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationController struct {
	BaseController
}

type RegistrationFormSubmission struct {
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmpassword"`
}

func (c *RegistrationController) Get() {
	if c.Data["IsSignedIn"] == false {
		c.XSRFExpire = 7200
		c.Data["xsrfdata"] = c.XSRFFormHTML()
		c.Layout = "layout.tpl"
		c.TplName = "signup.tpl"
	} else {
		flash := beego.NewFlash()
		flash.Error("You are already signed in!")
		flash.Store(&c.Controller)
		c.Ctx.Redirect(302, "/")
	}
}

func (c *RegistrationController) Post() {
	registrationFormSubmission := RegistrationFormSubmission{}

	if err := c.ParseForm(&registrationFormSubmission); err != nil {
		println(err.Error())
	}

	if registrationFormSubmission.Password == registrationFormSubmission.ConfirmPassword {
		o := orm.NewOrm()
		user := models.User{Email: registrationFormSubmission.Email}
		err := o.Read(&user)

		if err == orm.ErrNoRows {
			// Email is not in use. Register the account.
			new_user := new(models.User)
			new_user.Email = registrationFormSubmission.Email

			passwordHash, hashErr := bcrypt.GenerateFromPassword([]byte(registrationFormSubmission.Password), 12)

			if hashErr != nil {
				fmt.Println(err)

				flash := beego.NewFlash()

				flash.Warning("An error occured generating a password hash, please try again later!")
				flash.Store(&c.Controller)

				c.Ctx.Redirect(302, "/sign-up")
			}

			new_user.PasswordHash = string(passwordHash)
			_, insertErr := o.Insert(new_user)

			if insertErr != nil {
				fmt.Println(err)

				flash := beego.NewFlash()

				flash.Warning("An error occured saving your account, please try again later!")
				flash.Store(&c.Controller)

				c.Ctx.Redirect(302, "/sign-up")
			} else {
				flash := beego.NewFlash()

				flash.Success("Your account was successfully created and is ready for you to sign-in!")
				flash.Store(&c.Controller)

				c.Ctx.Redirect(302, "/sign-in")
			}
		} else {
			// Email is already in use. Throw an error.
			flash := beego.NewFlash()

			flash.Warning("The Email you entered is already in use. Please sign-in instead!")
			flash.Store(&c.Controller)

			c.Ctx.Redirect(302, "/sign-in")
		}
	} else {
		flash := beego.NewFlash()

		flash.Warning("The Password and Confirm Password do not match. Please check the spelling and try again!")
		flash.Store(&c.Controller)

		c.Ctx.Redirect(302, "/sign-up")
	}
}
