package main

import (
	"github.com/astaxie/beego"
	_ "github.com/huanghuangzym/moshopserver/models"
	_ "github.com/huanghuangzym/moshopserver/routers"
	"github.com/huanghuangzym/moshopserver/services"
	_ "github.com/huanghuangzym/moshopserver/utils"
)

func main() {

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.Listen.HTTPAddr = "0.0.0.0"
	beego.BConfig.Listen.HTTPPort = 8080

	beego.InsertFilter("/api/cart/*", beego.BeforeExec, services.FilterFunc, true, true)
	beego.InsertFilter("/api/order/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.InsertFilter("/api/collect/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.InsertFilter("/api/address/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8080

}
