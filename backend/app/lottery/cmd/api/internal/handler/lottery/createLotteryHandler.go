package lottery

import (
	"looklook/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/lottery/cmd/api/internal/logic/lottery"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
)

func CreateLotteryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateLotteryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := lottery.NewCreateLotteryLogic(r.Context(), svcCtx)
		resp, err := l.CreateLottery(&req)
		//注意 handler这里需要用result.HttpResult() 才会返回    "code": 200, "msg": "OK",
		result.HttpResult(r, w, resp, err)
	}
}
