package admin

import (
	"github.com/astaxie/beego"
	"tnvd/controllers/lib"
	. "tnvd/models"
)

type AdminController struct {
	lib.AuthController
}

func (this *AdminController) Main() {
	this.TplNames = "admin.html"
}
func (this *AdminController) Exit() {
	this.DelSession("userId")
	this.Redirect(this.UrlFor("LoginController.Login"), 302)
}

type LoginController struct {
	lib.BaseController
}

func (this *LoginController) Login() {
	if this.GetSession("userId") != nil {
		this.Redirect("/admin", 302)
		return
	}
	if this.Ctx.Input.Method() == "GET" {
		beego.ReadFromRequest(&this.Controller)
		this.Layout = ``
		this.TplNames = `login.html`

	} else if this.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		username := this.GetString("username")
		password := this.GetString("password")

		userid, err := AdminCollection.ValidateUser(username, password)
		if err != nil {
			flash.Error("用户名或密码错误")
			flash.Store(&this.Controller)
			this.Redirect(this.UrlFor(".Login"), 302)
			return
		}
		this.SetSession("userId", userid)
		this.Redirect(`/admin`, 302)

		return
	}

}
