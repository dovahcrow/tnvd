package routers

import (
	"github.com/astaxie/beego"
	"tnvd/controllers/admin"
	adminapi "tnvd/controllers/admin/api"
	"tnvd/controllers/normal"
	normalapi "tnvd/controllers/normal/api"
)

func init() {
	beego.AddNamespace(normalapi.ApiNS)
	beego.Router("/", &normal.MainController{}, "get:Main")
	beego.AddNamespace(adminapi.ApiNS)
	beego.Router("/admin", &admin.AdminController{}, "get:Main")
	beego.Router("/admin/login", &admin.LoginController{}, "get,post:Login")
	beego.Router("/admin/exit", &admin.AdminController{}, "get:Exit")
}
