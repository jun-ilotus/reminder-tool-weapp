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

var _ TaskModel = (*customTaskModel)(nil)

type (
	// TaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskModel.
	TaskModel interface {
		taskModel

		FindUserConTaskDone(ctx context.Context, userId int64, days int64) ([]*Task, error)
	}

	customTaskModel struct {
		*defaultTaskModel
	}
)

// NewTaskModel returns a model for the database table.
func NewTaskModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TaskModel {
	return &customTaskModel{
		defaultTaskModel: newTaskModel(conn, c, opts...),
	}
}

func (m *customTaskModel) FindUserConTaskDone(ctx context.Context, userId int64, days int64) ([]*Task, error) {
	var query string
	query = fmt.Sprintf(`select %s from %s as t1 where t1.value >= ? and t1.type = 1 
    AND NOT EXISTS (
      SELECT *
      FROM task_finish AS tf 
      WHERE tf.user_id = ?
        AND tf.task_id = t1.id
	) 
	order by t1.id`, taskRows, m.table)
	var resp []*Task

	err := m.QueryRowNoCacheCtx(ctx, &resp, query, days, userId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindUserConTaskDone, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
}
