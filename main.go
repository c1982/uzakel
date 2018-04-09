package main

import (
	"os"
	"os/user"

	"github.com/astaxie/beego"
)

func init() {
	user, err := user.Current()
	if err != nil {
		beego.Error(err)
		os.Exit(1)
	}
	beego.LoadAppConfig("ini", user.HomeDir+"/.uzakel")
	beego.BConfig.AppName = "Uzak El"
	beego.BConfig.RunMode = "dev"
	beego.BConfig.ServerName = "Uzak El Api"

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.WebConfig.EnableDocs = false
	beego.BConfig.WebConfig.DirectoryIndex = false

	beego.BConfig.Listen.HTTPAddr = "localhost"
	beego.BConfig.Listen.HTTPPort = 5000
}

func main() {
	beego.ErrorController(&Error{})
	beego.Router("/", &Agent{}, "get:Index")
	beego.Router("/push", &Agent{}, "post:Push")
	beego.Router("/info", &Agent{}, "post:Info")
	beego.Run()
}
