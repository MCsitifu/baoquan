package controllers

import (
	"DateCertproject/db_mysql"
	"github.com/astaxie/beego"
	"hallo/models"
)

type LoginController struct {
	beego.Controller
}


func (l *LoginConroller) Get()  {
    l.TplName = "login.html"
}

func (l *LoginController) Get() {
	var user models.User
	err := l.ParseFrom(&user)
	if err !=nil{

		l.Ctx.WriteString("抱歉，用户信息解析失败，请重试")
		return
	}
	u,err:=user.QueryUser()
	if err!= nil{
	l.Ctx.WriteString("抱歉，用户登陆失败，请重试")
		return
	}
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"
}