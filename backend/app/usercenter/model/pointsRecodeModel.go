package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PointsRecodeModel = (*customPointsRecodeModel)(nil)

type (
	// PointsRecodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPointsRecodeModel.
	PointsRecodeModel interface {
		pointsRecodeModel
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
