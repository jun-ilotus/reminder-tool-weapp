// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"looklook/common/globalkey"
)

var (
	reminderFieldNames          = builder.RawFieldNames(&Reminder{})
	reminderRows                = strings.Join(reminderFieldNames, ",")
	reminderRowsExpectAutoSet   = strings.Join(stringx.Remove(reminderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	reminderRowsWithPlaceHolder = strings.Join(stringx.Remove(reminderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheReminderReminderIdPrefix = "cache:reminder:reminder:id:"
)

type (
	reminderModel interface {
		Insert(ctx context.Context, data *Reminder) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Reminder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Reminder, error)
		Update(ctx context.Context, data *Reminder) error
		List(ctx context.Context, page, limit int64) ([]*Reminder, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *Reminder) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Reminder, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Reminder, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Reminder, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Reminder, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Reminder, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultReminderModel struct {
		sqlc.CachedConn
		table string
	}

	Reminder struct {
		Id           int64          `db:"id"`
		UserId       int64          `db:"user_id"` // 创建提醒用户id
		CreateTime   time.Time      `db:"create_time"`
		UpdateTime   time.Time      `db:"update_time"`
		ReminderTime time.Time      `db:"reminder_time"` // 提醒时间
		Content      string         `db:"content"`       // 提醒内容
		Status       int64          `db:"status"`        // 0待提醒，1已提醒，2已完成
		Member       int64          `db:"member"`        // 0提醒自己，1提醒亲友，2共同提醒
		Rereminder   sql.NullString `db:"rereminder"`    // 重复
	}
)

func newReminderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultReminderModel {
	return &defaultReminderModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`reminder`",
	}
}

func (m *defaultReminderModel) Delete(ctx context.Context, id int64) error {
	reminderReminderIdKey := fmt.Sprintf("%s%v", cacheReminderReminderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, reminderReminderIdKey)
	return err
}

func (m *defaultReminderModel) FindOne(ctx context.Context, id int64) (*Reminder, error) {
	reminderReminderIdKey := fmt.Sprintf("%s%v", cacheReminderReminderIdPrefix, id)
	var resp Reminder
	err := m.QueryRowCtx(ctx, &resp, reminderReminderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", reminderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultReminderModel) Insert(ctx context.Context, data *Reminder) (sql.Result, error) {
	reminderReminderIdKey := fmt.Sprintf("%s%v", cacheReminderReminderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, reminderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ReminderTime, data.Content, data.Status, data.Member, data.Rereminder)
	}, reminderReminderIdKey)
	return ret, err
}

func (m *defaultReminderModel) TransInsert(ctx context.Context, session sqlx.Session, data *Reminder) (sql.Result, error) {
	reminderReminderIdKey := fmt.Sprintf("%s%v", cacheReminderReminderIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, reminderRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.ReminderTime, data.Content, data.Status, data.Member, data.Rereminder)
	}, reminderReminderIdKey)
	return ret, err
}
func (m *defaultReminderModel) Update(ctx context.Context, data *Reminder) error {
	reminderReminderIdKey := fmt.Sprintf("%s%v", cacheReminderReminderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, reminderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ReminderTime, data.Content, data.Status, data.Member, data.Rereminder, data.Id)
	}, reminderReminderIdKey)
	return err
}

func (m *defaultReminderModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Reminder) error {
	reminderReminderIdKey := fmt.Sprintf("%s%v", cacheReminderReminderIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, reminderRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.ReminderTime, data.Content, data.Status, data.Member, data.Rereminder, data.Id)
	}, reminderReminderIdKey)
	return err
}

func (m *defaultReminderModel) List(ctx context.Context, page, limit int64) ([]*Reminder, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", reminderRows, m.table)
	var resp []*Reminder
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultReminderModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultReminderModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultReminderModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultReminderModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Reminder, error) {

	builder = builder.Columns(reminderRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Reminder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultReminderModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Reminder, error) {

	builder = builder.Columns(reminderRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Reminder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultReminderModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Reminder, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(reminderRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*Reminder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultReminderModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Reminder, error) {

	builder = builder.Columns(reminderRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Reminder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultReminderModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Reminder, error) {

	builder = builder.Columns(reminderRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Reminder
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultReminderModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultReminderModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheReminderReminderIdPrefix, primary)
}

func (m *defaultReminderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", reminderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultReminderModel) tableName() string {
	return m.table
}
