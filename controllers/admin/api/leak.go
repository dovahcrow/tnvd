package api

import (
	"tnvd/controllers/lib"
	"tnvd/models"
)

func init() {
	ApiNS.Router("/leak/totalNum", &ApiLeakController{}, "get:Num")
}

type ApiLeakController struct {
	lib.AuthController
}

func (this *ApiLeakController) Num() {
	leaknum, err := models.LeakCollection.GetLeakNums()
	if err != nil {
		this.Abort(`500`)
		return
	}
	this.Data[`json`] = leaknum
	this.ServeJson()
}
