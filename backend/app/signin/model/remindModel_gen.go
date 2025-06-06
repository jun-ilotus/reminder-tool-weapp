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
	remindFieldNames          = builder.RawFieldNames(&Remind{})
	remindRows                = strings.Join(remindFieldNames, ",")
	remindRowsExpectAutoSet   = strings.Join(stringx.Remove(remindFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	remindRowsWithPlaceHolder = strings.Join(stringx.Remove(remindFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSigninRemindIdPrefix = "cache:signin:remind:id:"
)

type (
	remindModel interface {
		Insert(ctx context.Context, data *Remind) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Remind) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Remind, error)
		Update(ctx context.Context, data *Remind) error
		List(ctx context.Context, page, limit int64) ([]*Remind, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *Remind) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Remind, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Remind, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Remind, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Remind, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Remind, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultRemindModel struct {
		sqlc.CachedConn
		table string
	}

	Remind struct {
		Id         int64     `db:"id"`
		UserId     int64     `db:"user_id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		Status     int64     `db:"status"` // 0不开启，1开启
	}
)

func newRemindModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRemindModel {
	return &defaultRemindModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`remind`",
	}
}

func (m *defaultRemindModel) Delete(ctx context.Context, id int64) error {
	signinRemindIdKey := fmt.Sprintf("%s%v", cacheSigninRemindIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, signinRemindIdKey)
	return err
}

func (m *defaultRemindModel) FindOne(ctx context.Context, id int64) (*Remind, error) {
	signinRemindIdKey := fmt.Sprintf("%s%v", cacheSigninRemindIdPrefix, id)
	var resp Remind
	err := m.QueryRowCtx(ctx, &resp, signinRemindIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", remindRows, m.table)
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

func (m *defaultRemindModel) Insert(ctx context.Context, data *Remind) (sql.Result, error) {
	signinRemindIdKey := fmt.Sprintf("%s%v", cacheSigninRemindIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, remindRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Status)
	}, signinRemindIdKey)
	return ret, err
}

func (m *defaultRemindModel) TransInsert(ctx context.Context, session sqlx.Session, data *Remind) (sql.Result, error) {
	signinRemindIdKey := fmt.Sprintf("%s%v", cacheSigninRemindIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, remindRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.Status)
	}, signinRemindIdKey)
	return ret, err
}
func (m *defaultRemindModel) Update(ctx context.Context, data *Remind) error {
	signinRemindIdKey := fmt.Sprintf("%s%v", cacheSigninRemindIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, remindRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Status, data.Id)
	}, signinRemindIdKey)
	return err
}

func (m *defaultRemindModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Remind) error {
	signinRemindIdKey := fmt.Sprintf("%s%v", cacheSigninRemindIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, remindRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.Status, data.Id)
	}, signinRemindIdKey)
	return err
}

func (m *defaultRemindModel) List(ctx context.Context, page, limit int64) ([]*Remind, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", remindRows, m.table)
	var resp []*Remind
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultRemindModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultRemindModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultRemindModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultRemindModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Remind, error) {

	builder = builder.Columns(remindRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Remind
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRemindModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Remind, error) {

	builder = builder.Columns(remindRows)

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

	var resp []*Remind
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRemindModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Remind, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(remindRows)

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

	var resp []*Remind
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultRemindModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Remind, error) {

	builder = builder.Columns(remindRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Remind
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRemindModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Remind, error) {

	builder = builder.Columns(remindRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Remind
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRemindModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultRemindModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSigninRemindIdPrefix, primary)
}

func (m *defaultRemindModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", remindRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRemindModel) tableName() string {
	return m.table
}
