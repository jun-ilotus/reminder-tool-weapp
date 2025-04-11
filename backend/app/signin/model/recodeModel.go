package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RecodeModel = (*customRecodeModel)(nil)

type (
	// RecodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecodeModel.
	RecodeModel interface {
		recodeModel
	}

	customRecodeModel struct {
		*defaultRecodeModel
	}
)

// NewRecodeModel returns a model for the database table.
func NewRecodeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RecodeModel {
	return &customRecodeModel{
		defaultRecodeModel: newRecodeModel(conn, c, opts...),
	}
}
