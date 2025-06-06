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
	taskFinishFieldNames          = builder.RawFieldNames(&TaskFinish{})
	taskFinishRows                = strings.Join(taskFinishFieldNames, ",")
	taskFinishRowsExpectAutoSet   = strings.Join(stringx.Remove(taskFinishFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	taskFinishRowsWithPlaceHolder = strings.Join(stringx.Remove(taskFinishFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSigninTaskFinishIdPrefix = "cache:signin:taskFinish:id:"
)

type (
	taskFinishModel interface {
		Insert(ctx context.Context, data *TaskFinish) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *TaskFinish) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TaskFinish, error)
		Update(ctx context.Context, data *TaskFinish) error
		List(ctx context.Context, page, limit int64) ([]*TaskFinish, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *TaskFinish) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*TaskFinish, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TaskFinish, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TaskFinish, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*TaskFinish, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*TaskFinish, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultTaskFinishModel struct {
		sqlc.CachedConn
		table string
	}

	TaskFinish struct {
		Id         int64     `db:"id"`
		TaskId     int64     `db:"task_id"` // 任务id
		UserId     int64     `db:"user_id"`
		CreateTime time.Time `db:"create_time"`
		FinishDate time.Time `db:"finish_date"` // 完成日期
		Points     int64     `db:"points"`      // 获得的积分
	}
)

func newTaskFinishModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTaskFinishModel {
	return &defaultTaskFinishModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`task_finish`",
	}
}

func (m *defaultTaskFinishModel) Delete(ctx context.Context, id int64) error {
	signinTaskFinishIdKey := fmt.Sprintf("%s%v", cacheSigninTaskFinishIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, signinTaskFinishIdKey)
	return err
}

func (m *defaultTaskFinishModel) FindOne(ctx context.Context, id int64) (*TaskFinish, error) {
	signinTaskFinishIdKey := fmt.Sprintf("%s%v", cacheSigninTaskFinishIdPrefix, id)
	var resp TaskFinish
	err := m.QueryRowCtx(ctx, &resp, signinTaskFinishIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", taskFinishRows, m.table)
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

func (m *defaultTaskFinishModel) Insert(ctx context.Context, data *TaskFinish) (sql.Result, error) {
	signinTaskFinishIdKey := fmt.Sprintf("%s%v", cacheSigninTaskFinishIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, taskFinishRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.TaskId, data.UserId, data.FinishDate, data.Points)
	}, signinTaskFinishIdKey)
	return ret, err
}

func (m *defaultTaskFinishModel) TransInsert(ctx context.Context, session sqlx.Session, data *TaskFinish) (sql.Result, error) {
	signinTaskFinishIdKey := fmt.Sprintf("%s%v", cacheSigninTaskFinishIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, taskFinishRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.TaskId, data.UserId, data.FinishDate, data.Points)
	}, signinTaskFinishIdKey)
	return ret, err
}
func (m *defaultTaskFinishModel) Update(ctx context.Context, data *TaskFinish) error {
	signinTaskFinishIdKey := fmt.Sprintf("%s%v", cacheSigninTaskFinishIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskFinishRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.TaskId, data.UserId, data.FinishDate, data.Points, data.Id)
	}, signinTaskFinishIdKey)
	return err
}

func (m *defaultTaskFinishModel) TransUpdate(ctx context.Context, session sqlx.Session, data *TaskFinish) error {
	signinTaskFinishIdKey := fmt.Sprintf("%s%v", cacheSigninTaskFinishIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskFinishRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.TaskId, data.UserId, data.FinishDate, data.Points, data.Id)
	}, signinTaskFinishIdKey)
	return err
}

func (m *defaultTaskFinishModel) List(ctx context.Context, page, limit int64) ([]*TaskFinish, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", taskFinishRows, m.table)
	var resp []*TaskFinish
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultTaskFinishModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultTaskFinishModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultTaskFinishModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultTaskFinishModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*TaskFinish, error) {

	builder = builder.Columns(taskFinishRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TaskFinish
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTaskFinishModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TaskFinish, error) {

	builder = builder.Columns(taskFinishRows)

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

	var resp []*TaskFinish
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTaskFinishModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TaskFinish, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(taskFinishRows)

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

	var resp []*TaskFinish
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultTaskFinishModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*TaskFinish, error) {

	builder = builder.Columns(taskFinishRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TaskFinish
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTaskFinishModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*TaskFinish, error) {

	builder = builder.Columns(taskFinishRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TaskFinish
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTaskFinishModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultTaskFinishModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSigninTaskFinishIdPrefix, primary)
}

func (m *defaultTaskFinishModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", taskFinishRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTaskFinishModel) tableName() string {
	return m.table
}
