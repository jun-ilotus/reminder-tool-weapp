package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
	"time"
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
		DoneRemindered(ctx context.Context, reminderId, status int64) error
		ReminderListByLessThanCurrentTime(ctx context.Context, currentTime time.Time) ([]*Reminder, error)
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
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ReminderList, &resp:%v, query:%v, status:%v, userId:%v, error: %v", &resp, query, status, userId, err)
	}
	return resp, nil
}

func (c *customReminderModel) ReminderListByLessThanCurrentTime(ctx context.Context, currentTime time.Time) ([]*Reminder, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where reminder_time <= ? and status = 0", reminderRows, c.table)
	var resp []*Reminder
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, currentTime)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ReminderListByLessThanCurrentTime, &resp:%v, query:%v,  error: %v", &resp, query, err)
	}
}

func (c *customReminderModel) DoneRemindered(ctx context.Context, reminderId, status int64) error {
	query := fmt.Sprintf("update %s set `status` = ? where `id` = ?", c.table)

	_, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.Exec(query, status, reminderId)
	})
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DoneRemindered, reminderId:%v error: %v", reminderId, err)
	}
	return nil
}
