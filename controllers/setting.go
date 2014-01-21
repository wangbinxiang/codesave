package controllers

import (
	"github.com/wangbinxiang/codesave/libs"
)

type SettingController struct {
	libs.BaseController
}

func (this *SettingController) Prepare() {
	this.BaseController.Prepare()
	this.LoginJump(true)
}

func (this *SettingController) Get() {

}

func (this *SettingController) Post() {

}
