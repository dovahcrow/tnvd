package lib

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type more []string

func (m *more) Add(s string) string {
	*m = append(*m, s)
	return ``
}
func (c *BaseController) Prepare() {
	c.Data[`moreStyles`] = &more{}
	c.Data[`moreScripts`] = &more{}
	c.Layout = "layout.html"
	c.Data["position"] = ``

}

type AuthController struct {
	BaseController
}

func (c *AuthController) Prepare() {
	c.BaseController.Prepare()
	if uid := c.GetSession("userId"); uid == nil {
		c.Abort("403")
	}
}
