package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	web.Controller
}

func (c *IndexController) Get() {
	c.TplName = "index.tpl"
}
