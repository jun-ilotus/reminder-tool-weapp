package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ PrizeModel = (*customPrizeModel)(nil)

type (
	// PrizeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeModel.
	PrizeModel interface {
		prizeModel

		PrizeListByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error)
	}

	customPrizeModel struct {
		*defaultPrizeModel
	}
)

// NewPrizeModel returns a model for the database table.
func NewPrizeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PrizeModel {
	return &customPrizeModel{
		defaultPrizeModel: newPrizeModel(conn, c, opts...),
	}
}

func (c *customPrizeModel) PrizeListByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where lottery_id = ? order by `level`", prizeRows, c.table)

	var resp []*Prize
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "PrizeListByLotteryId, &resp:%v, query:%v, lastId:%v, limit:%v, error: %v", &resp, query, err)
	}
	return resp, nil
}
