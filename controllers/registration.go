package controllers

type RegistrationController struct {
	BaseController
}

type RegistrationFormSubmission struct {
	Email           string `form:"email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirmpassword"`
}

func (c *RegistrationController) Get() {

}
