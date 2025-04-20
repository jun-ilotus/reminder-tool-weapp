package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RemindModel = (*customRemindModel)(nil)

type (
	// RemindModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRemindModel.
	RemindModel interface {
		remindModel
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
