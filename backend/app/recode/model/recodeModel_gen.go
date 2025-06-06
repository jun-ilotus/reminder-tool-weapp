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
	recodeFieldNames          = builder.RawFieldNames(&Recode{})
	recodeRows                = strings.Join(recodeFieldNames, ",")
	recodeRowsExpectAutoSet   = strings.Join(stringx.Remove(recodeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	recodeRowsWithPlaceHolder = strings.Join(stringx.Remove(recodeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheRecodeRecodeIdPrefix = "cache:recode:recode:id:"
)

type (
	recodeModel interface {
		Insert(ctx context.Context, data *Recode) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Recode) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Recode, error)
		Update(ctx context.Context, data *Recode) error
		List(ctx context.Context, page, limit int64) ([]*Recode, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *Recode) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Recode, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Recode, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Recode, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Recode, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Recode, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultRecodeModel struct {
		sqlc.CachedConn
		table string
	}

	Recode struct {
		Id         int64     `db:"id"`
		UserId     int64     `db:"user_id"`
		ItemsId    int64     `db:"items_id"` // 关联的items_id
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		Content    string    `db:"content"`     // 备注
		RecodeTime time.Time `db:"recode_time"` // 记录的日期
	}
)

func newRecodeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRecodeModel {
	return &defaultRecodeModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`recode`",
	}
}

func (m *defaultRecodeModel) Delete(ctx context.Context, id int64) error {
	recodeRecodeIdKey := fmt.Sprintf("%s%v", cacheRecodeRecodeIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, recodeRecodeIdKey)
	return err
}

func (m *defaultRecodeModel) FindOne(ctx context.Context, id int64) (*Recode, error) {
	recodeRecodeIdKey := fmt.Sprintf("%s%v", cacheRecodeRecodeIdPrefix, id)
	var resp Recode
	err := m.QueryRowCtx(ctx, &resp, recodeRecodeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", recodeRows, m.table)
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

func (m *defaultRecodeModel) Insert(ctx context.Context, data *Recode) (sql.Result, error) {
	recodeRecodeIdKey := fmt.Sprintf("%s%v", cacheRecodeRecodeIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, recodeRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ItemsId, data.Content, data.RecodeTime)
	}, recodeRecodeIdKey)
	return ret, err
}

func (m *defaultRecodeModel) TransInsert(ctx context.Context, session sqlx.Session, data *Recode) (sql.Result, error) {
	recodeRecodeIdKey := fmt.Sprintf("%s%v", cacheRecodeRecodeIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, recodeRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.ItemsId, data.Content, data.RecodeTime)
	}, recodeRecodeIdKey)
	return ret, err
}
func (m *defaultRecodeModel) Update(ctx context.Context, data *Recode) error {
	recodeRecodeIdKey := fmt.Sprintf("%s%v", cacheRecodeRecodeIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, recodeRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ItemsId, data.Content, data.RecodeTime, data.Id)
	}, recodeRecodeIdKey)
	return err
}

func (m *defaultRecodeModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Recode) error {
	recodeRecodeIdKey := fmt.Sprintf("%s%v", cacheRecodeRecodeIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, recodeRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.ItemsId, data.Content, data.RecodeTime, data.Id)
	}, recodeRecodeIdKey)
	return err
}

func (m *defaultRecodeModel) List(ctx context.Context, page, limit int64) ([]*Recode, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", recodeRows, m.table)
	var resp []*Recode
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultRecodeModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultRecodeModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultRecodeModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultRecodeModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Recode, error) {

	builder = builder.Columns(recodeRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Recode
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRecodeModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Recode, error) {

	builder = builder.Columns(recodeRows)

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

	var resp []*Recode
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRecodeModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Recode, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(recodeRows)

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

	var resp []*Recode
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultRecodeModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Recode, error) {

	builder = builder.Columns(recodeRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Recode
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRecodeModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Recode, error) {

	builder = builder.Columns(recodeRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Recode
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultRecodeModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultRecodeModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheRecodeRecodeIdPrefix, primary)
}

func (m *defaultRecodeModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", recodeRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRecodeModel) tableName() string {
	return m.table
}
