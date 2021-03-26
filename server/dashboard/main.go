// dashboard example

package main

import (
	_ "github.com/admpub/frp/assets/frps/statik"

	"github.com/admpub/frp/pkg/config"
	"github.com/admpub/frp/server"

	"github.com/webx-top/echo"
	"github.com/webx-top/echo/engine/standard"
	"github.com/webx-top/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Log())
	e.RouteDebug = true
	g := e.Group(`/frp`)
	svr, err := server.NewService(config.GetDefaultServerConf())
	if err != nil {
		panic(err)
	}
	svr.RegisterTo(g)
	// svr.Run()
	e.Run(standard.New(`:8080`))
}
