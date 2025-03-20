package items

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/recode/cmd/api/internal/logic/items"
	"looklook/app/recode/cmd/api/internal/svc"
	"looklook/app/recode/cmd/api/internal/types"
)

func ItemsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ItemsListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := items.NewItemsListLogic(r.Context(), svcCtx)
		resp, err := l.ItemsList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			result.HttpResult(r, w, resp, err)
		}
	}
}
