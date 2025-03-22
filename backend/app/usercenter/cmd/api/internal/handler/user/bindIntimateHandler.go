package user

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/usercenter/cmd/api/internal/logic/user"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
)

func BindIntimateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BindIntimateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewBindIntimateLogic(r.Context(), ctx)
		resp, err := l.BindIntimate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			result.HttpResult(r, w, resp, err)
		}
	}
}
