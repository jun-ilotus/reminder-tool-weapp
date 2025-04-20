package remind

import (
	"looklook/app/signin/cmd/api/internal/handler/translator"
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/signin/cmd/api/internal/logic/remind"
	"looklook/app/signin/cmd/api/internal/svc"
	"looklook/app/signin/cmd/api/internal/types"
)

func ChangeSignRemindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangeSignRemindReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := remind.NewChangeSignRemindLogic(r.Context(), svcCtx)
		resp, err := l.ChangeSignRemind(&req)

		result.HttpResult(r, w, resp, err)
	}
}
