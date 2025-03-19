package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ReminderModel = (*customReminderModel)(nil)

type (
	// ReminderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReminderModel.
	ReminderModel interface {
		reminderModel
	}

	customReminderModel struct {
		*defaultReminderModel
	}
)

// NewReminderModel returns a model for the database table.
func NewReminderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ReminderModel {
	return &customReminderModel{
		defaultReminderModel: newReminderModel(conn, c, opts...),
	}
}
