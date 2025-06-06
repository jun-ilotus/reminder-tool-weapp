package recode

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/recode/cmd/api/internal/logic/recode"
	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"
)

func DeleteRecodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteRecodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := recode.NewDeleteRecodeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteRecode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			result.HttpResult(r, w, resp, err)
		}
	}
}
