package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"nichebox/service/box_info/api/internal/logic"
	"nichebox/service/box_info/api/internal/svc"
	"nichebox/service/box_info/api/internal/types"
)

func CreateBoxHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBoxRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateBoxLogic(r.Context(), svcCtx)
		resp, err := l.CreateBox(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
