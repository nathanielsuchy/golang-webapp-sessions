package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}
