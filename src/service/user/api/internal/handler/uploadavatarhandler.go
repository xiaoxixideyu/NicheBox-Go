package handler

import (
	"net/http"

	"nichebox/service/user/api/internal/logic"
	"nichebox/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadAvatarLogic(r, svcCtx)
		resp, err := l.UploadAvatar()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
