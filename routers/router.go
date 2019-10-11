// @APIVersion 1.0.0
// @Title beego Test API
// @Description 配置路由地址
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/sinksmell/lanblog/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/topic",
			beego.NSInclude(
				&controllers.TopicController{},
			),
		),
		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/label",
			beego.NSInclude(
				&controllers.LabelController{},
			),
		),
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.AdminController{},
			),
		),
	)
	beego.AddNamespace(ns)

}
