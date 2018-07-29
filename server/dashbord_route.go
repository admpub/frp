package server

import (
	"context"
	"net/http"

	"github.com/admpub/frp/assets"
	"github.com/webx-top/echo"
)

// NewRouteGroup 为echo框架创建路由组
func NewRouteGroup(prefix string, e *echo.Echo) *echo.Group {
	g := e.Group(prefix)
	// api, see dashboard_api.go
	g.Get("/api/serverinfo", apiServerInfo)
	g.Get("/api/proxy/:type", func(ctx echo.Context) error {
		r := ctx.Request().StdRequest()
		v := map[string]string{
			`type`: ctx.Param(`type`),
		}
		r.WithContext(context.WithValue(r.Context(), 0, v))
		apiProxyByType(ctx.Response().StdResponseWriter(), r)
		return nil
	})
	g.Get("/api/proxy/:type/:name", func(ctx echo.Context) error {
		r := ctx.Request().StdRequest()
		v := map[string]string{
			`type`: ctx.Param(`type`),
			`name`: ctx.Param(`name`),
		}
		r.WithContext(context.WithValue(r.Context(), 0, v))
		apiProxyByTypeAndName(ctx.Response().StdResponseWriter(), r)
		return nil
	})
	g.Get("/api/traffic/:name", func(ctx echo.Context) error {
		r := ctx.Request().StdRequest()
		v := map[string]string{
			`name`: ctx.Param(`name`),
		}
		r.WithContext(context.WithValue(r.Context(), 0, v))
		apiProxyTraffic(ctx.Response().StdResponseWriter(), r)
		return nil
	})

	// view
	g.Get("/static/", http.StripPrefix(prefix+"/static/", http.FileServer(assets.FileSystem)))

	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "./static/", http.StatusMovedPermanently)
	})
	return g
}
