package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ ReminderModel = (*customReminderModel)(nil)

type (
	// ReminderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customReminderModel.
	ReminderModel interface {
		reminderModel

		// 自定义方法
		GetLastId(ctx context.Context) (int64, error)
		ReminderList(ctx context.Context, status, userId int64) ([]*Reminder, error)
		DoneRemindered(ctx context.Context, reminderId int64) error
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

func (c *customReminderModel) GetLastId(ctx context.Context) (int64, error) {
	var resp int64
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLastId, error: %v", err)
	}
	return resp, nil
}

func (c *customReminderModel) ReminderList(ctx context.Context, status, userId int64) ([]*Reminder, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where status = ? and user_id = ? order by reminder_time desc", reminderRows, c.table)
	var resp []*Reminder
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, status, userId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, status:%v, userId:%v, error: %v", &resp, query, status, userId, err)
	}
	return resp, nil
}

func (c *customReminderModel) DoneRemindered(ctx context.Context, reminderId int64) error {
	var resp int64
	query := fmt.Sprintf("update %s set `status` = 1 where `id` = ?", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, reminderId)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLastId, error: %v", err)
	}
	return nil
}
