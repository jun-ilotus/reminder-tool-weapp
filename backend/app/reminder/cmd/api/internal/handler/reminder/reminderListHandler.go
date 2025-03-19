package reminder

import (
	"looklook/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/reminder/cmd/api/internal/logic/reminder"
	"looklook/app/reminder/cmd/api/internal/svc"
	"looklook/app/reminder/cmd/api/internal/types"
)

func ReminderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReminderListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := reminder.NewReminderListLogic(r.Context(), svcCtx)
		resp, err := l.ReminderList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			result.HttpResult(r, w, resp, err)
		}
	}
}
