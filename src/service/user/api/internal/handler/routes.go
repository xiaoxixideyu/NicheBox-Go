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
				Path:    "/api/user/checkemailexists",
				Handler: CheckEmailExistsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/forgetpassword",
				Handler: ForgetPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/sendverificationcodepwd",
				Handler: SendVerificationCodePWDHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/sendverificationcoderegister",
				Handler: SendVerificationCodeRegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/checkverificationcodecriticaluserinfo",
				Handler: CheckVerificationCodeCriticalUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/refreshtoken",
				Handler: RefreshTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/sendverificationcodecriticaluserinfo",
				Handler: SendVerificationCodeCriticalUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/setcriticaluserinfo",
				Handler: SetCriticalUserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
