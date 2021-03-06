package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/huanghuangzym/moshopserver/models"
	"github.com/huanghuangzym/moshopserver/services"
	"github.com/huanghuangzym/moshopserver/utils"
)

type AuthController struct {
	beego.Controller
}

type AuthLoginBody struct {
	Code     string               `json:"code"`
	UserInfo services.ResUserInfo `json:"userInfo"`
}

func (this *AuthController) Auth_LoginByWeixin() {

	var alb AuthLoginBody
	body := this.Ctx.Input.RequestBody

	err := json.Unmarshal(body, &alb)
	//fmt.Print(alb)
	clientIP := this.Ctx.Input.IP()

	userInfo := services.Login(alb.Code, alb.UserInfo)
	if userInfo == nil {

	}
	fmt.Printf("Auth_LoginByWeixin  get userinfo %v \n", userInfo)

	o := orm.NewOrm()

	var user models.NideshopUser
	usertable := new(models.NideshopUser)
	err = o.QueryTable(usertable).Filter("weixin_openid", userInfo.OpenID).One(&user)
	if err == orm.ErrNoRows {
		newuser := models.NideshopUser{Username: utils.GetUUID(), Password: "", RegisterTime: utils.GetTimestamp(),
			RegisterIp: clientIP, Mobile: "", WeixinOpenid: userInfo.OpenID, Avatar: userInfo.AvatarUrl, Gender: userInfo.Gender,
			Nickname: userInfo.NickName}
		o.Insert(&newuser)
		o.QueryTable(usertable).Filter("weixin_openid", userInfo.OpenID).One(&user)
	}

	userinfo := make(map[string]interface{})
	userinfo["id"] = user.Id
	userinfo["username"] = user.Username
	userinfo["nickname"] = user.Nickname
	userinfo["gender"] = user.Gender
	userinfo["avatar"] = user.Avatar
	userinfo["birthday"] = user.Birthday

	user.LastLoginIp = clientIP
	user.LastLoginTime = utils.GetTimestamp()

	if _, err := o.Update(&user); err == nil {

	}

	sessionKey := services.Create(utils.Int2String(user.Id))
	//fmt.Println("sessionkey==" + sessionKey)

	rtnInfo := make(map[string]interface{})
	rtnInfo["token"] = sessionKey
	rtnInfo["userInfo"] = userinfo

	utils.ReturnHTTPSuccess(&this.Controller, rtnInfo)
	this.ServeJSON()

}
