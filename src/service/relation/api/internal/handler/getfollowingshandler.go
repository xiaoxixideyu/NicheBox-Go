package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"nichebox/service/relation/api/internal/logic"
	"nichebox/service/relation/api/internal/svc"
	"nichebox/service/relation/api/internal/types"
)

func GetFollowingsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFollowingsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetFollowingsLogic(r.Context(), svcCtx)
		resp, err := l.GetFollowings(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
