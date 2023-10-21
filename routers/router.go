// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/kilianp07/movieDB/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/film", &controllers.FilmController{}, "post:Post")
	beego.Router("/film/:id", &controllers.FilmController{}, "get:GetOne")
	beego.Router("/film", &controllers.FilmController{}, "get:GetAll")
	beego.Router("/film/:id", &controllers.FilmController{}, "put:Put")
	beego.Router("/film/:id", &controllers.FilmController{}, "delete:Delete")

	beego.Router("/actor", &controllers.ActorController{}, "post:Post")
	beego.Router("/actor/:id", &controllers.ActorController{}, "get:GetOne")
	beego.Router("/actor", &controllers.ActorController{}, "get:GetAll")
	beego.Router("/actor/:id", &controllers.ActorController{}, "put:Put")
	beego.Router("/actor/:id", &controllers.ActorController{}, "delete:Delete")
}
