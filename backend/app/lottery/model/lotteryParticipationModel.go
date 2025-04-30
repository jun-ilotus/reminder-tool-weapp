package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ LotteryParticipationModel = (*customLotteryParticipationModel)(nil)

type (
	// LotteryParticipationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryParticipationModel.
	LotteryParticipationModel interface {
		lotteryParticipationModel

		GetLotteryParticipationByUserIdLotteryId(ctx context.Context, userId, lotteryId int64) (*LotteryParticipation, error)
	}

	customLotteryParticipationModel struct {
		*defaultLotteryParticipationModel
	}
)

// NewLotteryParticipationModel returns a model for the database table.
func NewLotteryParticipationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryParticipationModel {
	return &customLotteryParticipationModel{
		defaultLotteryParticipationModel: newLotteryParticipationModel(conn, c, opts...),
	}
}

func (c *customLotteryParticipationModel) GetLotteryParticipationByUserIdLotteryId(ctx context.Context, userId, lotteryId int64) (*LotteryParticipation, error) {
	var resp *LotteryParticipation
	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? AND lottery_id = ? LIMIT 1", lotteryParticipationRows, c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, userId, lotteryId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationByUserIdLotteryId, error: %v", err)
	}
}
