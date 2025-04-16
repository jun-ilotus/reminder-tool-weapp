package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskFinishModel = (*customTaskFinishModel)(nil)

type (
	// TaskFinishModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskFinishModel.
	TaskFinishModel interface {
		taskFinishModel

		DeleteByUserIdTaskId(ctx context.Context, userId, taskId int64) error
	}

	customTaskFinishModel struct {
		*defaultTaskFinishModel
	}
)

// NewTaskFinishModel returns a model for the database table.
func NewTaskFinishModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TaskFinishModel {
	return &customTaskFinishModel{
		defaultTaskFinishModel: newTaskFinishModel(conn, c, opts...),
	}
}

func (m *customTaskFinishModel) DeleteByUserIdTaskId(ctx context.Context, userId, taskId int64) error {
	signinTaskFinishIdKey := fmt.Sprintf("%s%v", cacheSigninTaskFinishIdPrefix, userId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ? and task_id = ?", m.table)
		return conn.ExecCtx(ctx, query, userId, taskId)
	}, signinTaskFinishIdKey)
	return err
}
