package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskModel = (*customTaskModel)(nil)

type (
	// TaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskModel.
	TaskModel interface {
		taskModel
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
