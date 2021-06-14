package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)
	if flash != nil {
	}
	c.Layout = "layout.tpl"
	c.TplName = "index.tpl"
}
