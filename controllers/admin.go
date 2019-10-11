package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sinksmell/lanblog/models"
)

type AdminController struct {
	beego.Controller
}

func (u *AdminController) URLMapping() {
	u.Mapping("Login", u.Login)
	u.Mapping("Options", u.Options)
}

// @Title AdminLogin
// @Description admin login
// @Param	body		body 	models.Admin	true		"body for Admin content"
// @Success 200
// @Failure 403 body is empty
// @router /login [post]
func (u *AdminController) Login() {
	result := &models.LoginResult{}
	var admin models.Admin
	json.Unmarshal(u.Ctx.Input.RequestBody, &admin)
	adminName := beego.AppConfig.String("adminName")
	adminPWD := beego.AppConfig.String("adminPWD")
	if admin.Name == adminName && admin.Password == adminPWD {
		result.Code = 0
		result.Msg = "OK"
		result.Data.User = admin
		result.Data.Token = models.GenToken()
		result.Data.Name = "管理员"
		result.Data.UUId = "sinksmell"
	} else {
		result.Code = 100
		result.Msg = "用户名或密码错误!"
	}
	u.Data["json"] = result
	u.ServeJSON()
}

// @Title BeforeAdminLogin
// @Description cros before admin login
// @router /login [options]
func (u *AdminController) Options() {
	//u.Data["json"] = "test"
	//u.ServeJSON()
}
