package api

import (
	"tnvd/controllers/lib"
	"tnvd/controllers/lib/menu"
)

type MainMenu struct {
	lib.BaseController
}

func (this *MainMenu) Get() {
	menu := menu.NewMenu()
	menu.SetBrand("TNVD", this.UrlFor("MainController.Main"))
	menu.AddSinMenu("首页", this.UrlFor("MainController.Main"))
	menu.AddSinMenu("热点关注", this.UrlFor("HotController.Main"))
	menu.AddSinMenu("漏洞列表", this.UrlFor("LeakListController"))
	menu.AddSinMenu("补丁信息", this.UrlFor("PatchController.Main"))
	menu.AddSinMenu("安全公告", this.UrlFor("BulletinController.Main"))
	menu.AddSinMenu("统计数据", this.UrlFor("StatisticController.Main"))

	// map[string]string{
	// 	"行业漏洞": this.UrlFor("LeakIndustryController.Main"),
	// 	"应用漏洞": this.UrlFor("LeakApplicationController.Main"),
	// })

	this.Data[`json`] = menu
	this.ServeJson()
}
