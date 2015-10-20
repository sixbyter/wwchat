package framework

import (
	"fmt"
	"github.com/realint/dbgutil"
	"net/http"
	"sixbyte/framework/configer"
	"sixbyte/framework/router"
	"sixbyte/framework/viewer"
)

type app struct {
	Router   *router.Router
	Configer *configer.Configer
	Viewer   *viewer.Viewer
}

var App *app

func init() {
	App = &app{}
	fmt.Println("创建一个新的实例")
	App.Configer = configer.NewConfiger()
	App.Router = router.NewRouter()
	App.Viewer = viewer.NewViewer()
}

func Run() {
	// 路由配置
	App.Router.StaticDir = GetConfig("ROUTE_STATIC_DIR")
	App.Router.StaticPath = GetConfig("ROUTE_STATIC_PATH")

	// 视图配置
	App.Viewer.ViewDir = GetConfig("VIEW_PATH")

	dbgutil.FormatDisplay("App", App)

	port := GetConfig("HTTP_PORT")
	http.ListenAndServe(":"+port, App.Router)
}
