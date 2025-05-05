package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
	"time"
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
		LotteryUserList(ctx context.Context, userId, isAnnounced int64) ([]*Lottery, error)
		GetLotterysByLessThanCurrentTime(ctx context.Context, currentTime time.Time, announceType int64) ([]int64, error)
		UpdateLotteryStatus(ctx context.Context, lotteryId int64) error
		GetLotterysByLessThanCurrentTime2(ctx context.Context, currentTime time.Time, announceType int64) ([]*Lottery, error)
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}
)

var CacheLotteryJoinNumberIdPrefix = "cache:lottery:join_number:id:"

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c, opts...),
	}
}

func (c *customLotteryModel) GetLotterysByLessThanCurrentTime2(ctx context.Context, currentTime time.Time, announceType int64) ([]*Lottery, error) {
	var resp []*Lottery
	query := fmt.Sprintf("SELECT %s FROM %s WHERE announce_type = ? AND is_announced = 0 AND announce_time <= ?", lotteryRows, c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, announceType, currentTime)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotterysByLessThanCurrentTime, CurrentTime:%v, anounceType:%v, error: %v", currentTime, announceType, err)
	}
}

func (c *customLotteryModel) GetLotterysByLessThanCurrentTime(ctx context.Context, currentTime time.Time, announceType int64) ([]int64, error) {
	var resp []int64
	query := fmt.Sprintf("SELECT id FROM %s WHERE announce_type = ? AND is_announced = 0 AND announce_time <= ?", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, announceType, currentTime)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotterysByLessThanCurrentTime, CurrentTime:%v, anounceType:%v, error: %v", currentTime, announceType, err)
	}
}

// UpdateLotteryStatus 根据lotteryId更新lottery状态为已开奖
func (c *customLotteryModel) UpdateLotteryStatus(ctx context.Context, lotteryId int64) error {
	// 准备更新数据的SQL语句
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, lotteryId)
	query := fmt.Sprintf("UPDATE %s SET is_announced = 1 WHERE id = ?", c.table)

	_, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.Exec(query, lotteryId)
	}, lotteryLotteryIdKey)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "UpdateLotteryStatus, lotteryId:%v error: %v", lotteryId, err)
	}
	return nil
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

func (c *customLotteryModel) LotteryUserList(ctx context.Context, userId, isAnnounced int64) ([]*Lottery, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where user_id = ? AND is_announced = ? order by id desc", lotteryRows, c.table)

	var resp []*Lottery
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId, isAnnounced)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryList, &resp:%v, query:%v, error: %v", &resp, query, err)
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
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryById, &resp:%v, query:%v, error: %v, lotteryId: %v", &resp, query, err, lotteryId)
	}
	return &resp, nil
}
