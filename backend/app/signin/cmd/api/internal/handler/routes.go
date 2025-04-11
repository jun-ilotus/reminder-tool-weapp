// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	recode "looklook/app/signin/cmd/api/internal/handler/recode"
	"looklook/app/signin/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// recode addToday
				Method:  http.MethodPost,
				Path:    "/recode/addToday",
				Handler: recode.AddTodayHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/signin/v1"),
	)
}
