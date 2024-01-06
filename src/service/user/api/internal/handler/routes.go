// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"nichebox/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/checkemailexists",
				Handler: CheckEmailExistsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/refreshtoken",
				Handler: RefreshTokenHandler(serverCtx),
			},
		},
	)
}
