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

var _ RemindModel = (*customRemindModel)(nil)

type (
	// RemindModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRemindModel.
	RemindModel interface {
		remindModel

		FindOneByUserId(ctx context.Context, userId int64) (*Remind, error)
	}

	customRemindModel struct {
		*defaultRemindModel
	}
)

// NewRemindModel returns a model for the database table.
func NewRemindModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RemindModel {
	return &customRemindModel{
		defaultRemindModel: newRemindModel(conn, c, opts...),
	}
}

func (c *customRemindModel) FindOneByUserId(ctx context.Context, userId int64) (*Remind, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where user_id = ?", remindRows, c.table)
	var resp Remind

	err := c.QueryRowNoCacheCtx(ctx, &resp, query, userId)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindOneByUserId remind, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
}
