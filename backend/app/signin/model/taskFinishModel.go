package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskFinishModel = (*customTaskFinishModel)(nil)

type (
	// TaskFinishModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskFinishModel.
	TaskFinishModel interface {
		taskFinishModel
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
