package conf

import (
	"sixbyte/framework"
)

func init() {
	framework.SetConfig("APP_Name", "myfirst frame for go")
	framework.SetConfig("APP_PATH", "./src/project")
	framework.SetConfig("HTTP_PORT", "8080")
	framework.SetConfig("RUN_MODE", "dev")

	framework.SetConfig("ROUTE_STATIC_DIR", framework.GetConfig("APP_PATH")+"/asset")
	framework.SetConfig("ROUTE_STATIC_PATH", "/asset")
	framework.SetConfig("VIEW_PATH", framework.GetConfig("APP_PATH")+"/views")
}
