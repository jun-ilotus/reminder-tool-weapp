package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ LotteryModel = (*customLotteryModel)(nil)

type (
	// LotteryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryModel.
	LotteryModel interface {
		lotteryModel
		LotteryList(ctx context.Context, limit, selected, lastId int64) ([]*Lottery, error)
		GetLastId(ctx context.Context) (int64, error)
		GetLotteryById(ctx context.Context, lotteryId int64) (*Lottery, error)
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}
)

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c, opts...),
	}
}

func (c *customLotteryModel) LotteryList(ctx context.Context, limit, selected, lastId int64) ([]*Lottery, error) {
	var query string
	if selected != 0 {
		query = fmt.Sprintf("select %s from %s where is_selected = 1 and is_announced = 0 and id < ? order by id desc limit ?", lotteryRows, c.table)
	} else {
		query = fmt.Sprintf("select %s from %s where is_announced = 0 and id < ? order by id desc limit ?", lotteryRows, c.table)
	}
	var resp []*Lottery
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lastId, limit)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryList, &resp:%v, query:%v, lastId:%v, limit:%v, error: %v", &resp, query, lastId, limit, err)
	}
	return resp, nil
}

func (c *customLotteryModel) GetLastId(ctx context.Context) (int64, error) {
	var resp int64
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLastId, error: %v", err)
	}
	return resp, nil
}

func (c *customLotteryModel) GetLotteryById(ctx context.Context, lotteryId int64) (*Lottery, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where id = ?", lotteryRows, c.table)

	var resp Lottery
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, lotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryById, &resp:%v, query:%v, lastId:%v, limit:%v, error: %v", &resp, query, err)
	}
	return &resp, nil
}
