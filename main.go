package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "tnvd/routers"
)

func main() {
	beego.TemplateLeft = "[["
	beego.TemplateRight = "]]"
	beego.Run()
}
