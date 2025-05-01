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
		GetLotteryParticipationCount(ctx context.Context, lotteryId, isWin int64) (int64, error)
		GetLotteryParticipationLastId(ctx context.Context, lotteryId int64) (int64, error)
		LotteryParticipationList(ctx context.Context, lotteryId, limit, lastId int64) ([]*LotteryParticipation, error)
		LotteryParticipationWinList(ctx context.Context, lotteryId int64) ([]*LotteryParticipation, error)
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
	var resp LotteryParticipation
	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? AND lottery_id = ? LIMIT 1", lotteryParticipationRows, c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, userId, lotteryId)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationByUserIdLotteryId, error: %v", err)
	}
}

func (c *customLotteryParticipationModel) GetLotteryParticipationCount(ctx context.Context, lotteryId, isWin int64) (int64, error) {
	var resp int64

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE lottery_id = ?", c.table)
	if isWin == 1 {
		query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE lottery_id = ? AND is_won = 1", c.table)
	}
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, lotteryId)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationCount, error: %v", err)
	}
	return resp, nil
}

func (c *customLotteryParticipationModel) GetLotteryParticipationLastId(ctx context.Context, lotteryId int64) (int64, error) {
	var resp int64
	query := fmt.Sprintf("SELECT id FROM %s WHERE lottery_id = ? ORDER BY id DESC LIMIT 1", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, lotteryId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationLastId, error: %v", err)
	}
}

func (c *customLotteryParticipationModel) LotteryParticipationList(ctx context.Context, lotteryId, limit, lastId int64) ([]*LotteryParticipation, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where lottery_id = ? AND id < ? order by id desc limit ?", lotteryParticipationRows, c.table)

	var resp []*LotteryParticipation
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lotteryId, lastId, limit)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryParticipationList, &resp:%v, query:%v, lastId:%v, limit:%v, error: %v", &resp, query, lastId, limit, err)
	}
}

func (c *customLotteryParticipationModel) LotteryParticipationWinList(ctx context.Context, lotteryId int64) ([]*LotteryParticipation, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where lottery_id = ? and is_won = 1 order by prize_id", lotteryParticipationRows, c.table)

	var resp []*LotteryParticipation
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lotteryId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryParticipationList, &resp:%v, query:%v, error: %v", &resp, query, err)
	}
}
