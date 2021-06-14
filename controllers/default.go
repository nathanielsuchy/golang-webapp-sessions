package controllers

type MainController struct {
	PrivateController
}

func (c *MainController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}
