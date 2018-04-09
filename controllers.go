package main

import (
	"strings"

	"github.com/astaxie/beego"
)

type Agent struct {
	IpAddr   string
	IsAccess bool
	beego.Controller
}

// Bundan sonraki controller'ler için Prepare+Yeni Kurallar
type NestPreparer interface {
	NestPrepare()
}

// Bundan sonraki controller'ler için Finish+Yeni Kurallar
type NestFinisher interface {
	NestFinish()
}

func (c *Agent) Prepare() {
	c.IpAddr = strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	c.EnableRender = false

	if err := Auth(c.GetString("secret"), c.IpAddr); err != nil {
		c.Data["json"] = JsonData(false, err.Error())
		c.ServeJSON()
		return
	}

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (c *Agent) Finish() {
	// Tum kuralları NestFinish fonksiyonununa aktarmaktadır.
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}

func (c *Agent) Index() {
	c.Data["json"] = JsonData(false, "Test API", nil)
	c.ServeJSON()
}

func (c *Agent) Info() {
	c.Data["json"] = JsonData(false, "Test API", nil)
	c.ServeJSON()
}

func (c *Agent) Push() {
	c.Data["json"] = JsonData(false, "Test API", nil)
	c.ServeJSON()
}

/*
  Error Controller
*/

type Error struct {
	Agent
}

func (c *Error) NestPrepare() {
}

func (c *Error) Error404() {
	c.Data["json"] = JsonData(false, "404 not found")
	c.ServeJSON()
}

func (c *Error) Error500() {
	c.Data["json"] = JsonData(false, "500 internal server error")
	c.ServeJSON()
}

func (c *Error) Error504() {
	c.Data["json"] = JsonData(false, "504 https zorunludur")
	c.ServeJSON()
}

func (c *Error) Error401() {
	c.Data["json"] = JsonData(false, "401 not authorized")
	c.ServeJSON()
}
