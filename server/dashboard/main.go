// dashboard example

package main

import (
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
	server.RegisterTo(g)
	e.Run(standard.New(`:4444`))
}
