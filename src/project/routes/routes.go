package routes

import (
	"project/controllers"
	"sixbyte/framework"
)

func init() {
	// react
	framework.App.Router.RegisterFunc("GET", "/react/nihao", controllers.React.Nihao)
	framework.App.Router.RegisterFunc("GET", "/react/comment", controllers.React.GetComment)
	framework.App.Router.RegisterFunc("POST", "/react/comment", controllers.React.PostComment)
	framework.App.Router.RegisterFunc("GET", "/react/animation", controllers.React.AddonsAnimation)
	framework.App.Router.RegisterFunc("GET", "/react/amaze", controllers.React.AmazeUi)

	// websocket
	framework.App.Router.RegisterWebsocket("GET", "/websocket/server", controllers.Websocket.Server)
	framework.App.Router.RegisterFunc("GET", "/websocket/client", controllers.Websocket.Client)
	// framework.App.Router.RegisterFunc("GET", "/websocket/player", controllers.Websocket.GetPlayers)
}
