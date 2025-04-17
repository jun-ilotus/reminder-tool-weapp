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

var _ PointsRecodeModel = (*customPointsRecodeModel)(nil)

type (
	// PointsRecodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPointsRecodeModel.
	PointsRecodeModel interface {
		pointsRecodeModel

		FindLastOneByUserIdTaskId(ctx context.Context, userId, taskId int64) (*PointsRecode, error)
		FindUserList(ctx context.Context, userId int64) ([]*PointsRecode, error)
	}

	customPointsRecodeModel struct {
		*defaultPointsRecodeModel
	}
)

// NewPointsRecodeModel returns a model for the database table.
func NewPointsRecodeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PointsRecodeModel {
	return &customPointsRecodeModel{
		defaultPointsRecodeModel: newPointsRecodeModel(conn, c, opts...),
	}
}

func (c *customPointsRecodeModel) FindUserList(ctx context.Context, userId int64) ([]*PointsRecode, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ?", pointsRecodeRows, c.table)
	var resp []*PointsRecode
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	return resp, err
}

func (c *customPointsRecodeModel) FindLastOneByUserIdTaskId(ctx context.Context, userId, taskId int64) (*PointsRecode, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where user_id = ? and task_id = ? order by id desc limit 1", pointsRecodeRows, c.table)
	var resp PointsRecode

	err := c.QueryRowNoCacheCtx(ctx, &resp, query, userId, taskId)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindLastOneByUserIdTaskId, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
}
