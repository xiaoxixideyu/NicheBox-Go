package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"nichebox/service/relation/api/internal/logic"
	"nichebox/service/relation/api/internal/svc"
	"nichebox/service/relation/api/internal/types"
)

func UnfollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UnfollowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUnfollowLogic(r.Context(), svcCtx)
		resp, err := l.Unfollow(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
