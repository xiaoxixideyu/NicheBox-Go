package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"nichebox/service/user/api/internal/logic"
	"nichebox/service/user/api/internal/svc"
	"nichebox/service/user/api/internal/types"
)

func SendVerificationCodePWDHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendVerificationCodePWDRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendVerificationCodePWDLogic(r.Context(), svcCtx)
		resp, err := l.SendVerificationCodePWD(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
