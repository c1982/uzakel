package main

import (
	"errors"

	"github.com/astaxie/beego"
)

type Master struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

// Master isimli struct kullanılarak basit standart json nesleleri üretiminde kullanılır
func JsonData(status bool, message string, data ...interface{}) *Master {
	IsData := data != nil
	d := &Master{}
	d.Status = status
	if status {
		d.Message = message
	} else {
		d.Error = message
	}
	if IsData {
		d.Data = data
	}
	return d
}

// Auth için kullanılır string izin verilen IP adresini vermektedir!
func Auth(key, addr string) error {
	IpRange := beego.AppConfig.Strings("IpAccsess")
	var IpMatch bool
	for i, _ := range IpRange {
		if IpRange[i] == addr {
			IpMatch = true
		}
	}
	if !IpMatch {
		beego.Error("Yetkisiz IP adresinden erişim yapan :", addr)
		return errors.New(addr + " IP adresinden erişim yapamazsınız!")
	}
	if beego.AppConfig.String("ApiSecret") != key {
		beego.Error("Hatalı Secret Key Gönderen :", addr)
		return errors.New("Secret Key hatalı gönderildi!")
	}
	return nil
}
