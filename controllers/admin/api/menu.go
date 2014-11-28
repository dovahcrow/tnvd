package api

import (
	"tnvd/controllers/lib"
	"tnvd/controllers/lib/menu"
)

func init() {
	ApiNS.Router("mainMenu", &MainMenu{})
}

type MainMenu struct {
	lib.AuthController
}

func (this *MainMenu) Get() {
	menu := menu.NewMenu()
	menu.SetBrand("TNVDadmin", this.UrlFor("AdminController.Main"))
	menu.AddSinMenu("热点关注", "#/hot")
	menu.AddSinMenu("漏洞列表", "#/leak")
	menu.AddSinMenu("补丁信息", "#/patch")
	menu.AddSinMenu("安全公告", "#/bulletin")
	menu.AddSinMenu("统计数据", "#/statistic")
	this.Data[`json`] = menu
	this.ServeJson()
}
