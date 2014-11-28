package normal

import (
	"tnvd/controllers/lib"
)

type MainController struct {
	lib.BaseController
}

func (this *MainController) Main() {
	this.TplNames = "main.html"
}
