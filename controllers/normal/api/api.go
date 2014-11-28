package api

import (
	"github.com/astaxie/beego"
)

var ApiNS = beego.NewNamespace("/api")

func init() {
	ApiNS.Router("/mainMenu", &MainMenu{})
}
