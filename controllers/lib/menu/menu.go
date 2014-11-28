package menu

type Menu struct {
	Brand      string
	BrandUrl   string
	SinMenuBar []map[string]string
	MulMenuBar []map[string]interface{}
}

func NewMenu() *Menu {
	menu := Menu{}
	menu.SinMenuBar = []map[string]string{}
	menu.MulMenuBar = []map[string]interface{}{}
	return &menu
}

func (this *Menu) AddSinMenu(title string, url string) *Menu {
	this.SinMenuBar = append(this.SinMenuBar, map[string]string{
		"title": title,
		"url":   url,
	})
	return this
}

func (this *Menu) SetBrand(brand string, url string) *Menu {
	this.Brand = brand
	this.BrandUrl = url
	return this
}

func (this *Menu) AddMulMenu(title string, subs map[string]string) *Menu {
	m := make(map[string]interface{})
	m["title"] = title
	m["submenu"] = subs
	this.MulMenuBar = append(this.MulMenuBar, m)
	return this
}
