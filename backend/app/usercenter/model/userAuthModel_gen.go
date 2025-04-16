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
	userAuthFieldNames          = builder.RawFieldNames(&UserAuth{})
	userAuthRows                = strings.Join(userAuthFieldNames, ",")
	userAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userAuthRowsWithPlaceHolder = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLooklookUsercenterUserAuthIdPrefix              = "cache:looklookUsercenter:userAuth:id:"
	cacheLooklookUsercenterUserAuthAuthTypeAuthKeyPrefix = "cache:looklookUsercenter:userAuth:authType:authKey:"
	cacheLooklookUsercenterUserAuthUserIdAuthTypePrefix  = "cache:looklookUsercenter:userAuth:userId:authType:"
)

type (
	userAuthModel interface {
		Insert(ctx context.Context, data *UserAuth) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserAuth, error)
		FindOneByAuthTypeAuthKey(ctx context.Context, authType string, authKey string) (*UserAuth, error)
		FindOneByUserIdAuthType(ctx context.Context, userId int64, authType string) (*UserAuth, error)
		Update(ctx context.Context, data *UserAuth) error
		List(ctx context.Context, page, limit int64) ([]*UserAuth, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *UserAuth) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*UserAuth, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAuth, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAuth, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserAuth, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserAuth, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultUserAuthModel struct {
		sqlc.CachedConn
		table string
	}

	UserAuth struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		Version    int64     `db:"version"` // 版本号
		UserId     int64     `db:"user_id"`
		AuthKey    string    `db:"auth_key"`  // 平台唯一id
		AuthType   string    `db:"auth_type"` // 平台类型
	}
)

func newUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserAuthModel {
	return &defaultUserAuthModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_auth`",
	}
}

func (m *defaultUserAuthModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	looklookUsercenterUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	looklookUsercenterUserAuthIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserAuthIdPrefix, id)
	looklookUsercenterUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, looklookUsercenterUserAuthAuthTypeAuthKeyKey, looklookUsercenterUserAuthIdKey, looklookUsercenterUserAuthUserIdAuthTypeKey)
	return err
}

func (m *defaultUserAuthModel) FindOne(ctx context.Context, id int64) (*UserAuth, error) {
	looklookUsercenterUserAuthIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserAuthIdPrefix, id)
	var resp UserAuth
	err := m.QueryRowCtx(ctx, &resp, looklookUsercenterUserAuthIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAuthRows, m.table)
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

func (m *defaultUserAuthModel) FindOneByAuthTypeAuthKey(ctx context.Context, authType string, authKey string) (*UserAuth, error) {
	looklookUsercenterUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthAuthTypeAuthKeyPrefix, authType, authKey)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, looklookUsercenterUserAuthAuthTypeAuthKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `auth_type` = ? and `auth_key` = ? limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, authType, authKey); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindOneByUserIdAuthType(ctx context.Context, userId int64, authType string) (*UserAuth, error) {
	looklookUsercenterUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthUserIdAuthTypePrefix, userId, authType)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, looklookUsercenterUserAuthUserIdAuthTypeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `auth_type` = ? limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, authType); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) Insert(ctx context.Context, data *UserAuth) (sql.Result, error) {
	looklookUsercenterUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	looklookUsercenterUserAuthIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserAuthIdPrefix, data.Id)
	looklookUsercenterUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userAuthRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.UserId, data.AuthKey, data.AuthType)
	}, looklookUsercenterUserAuthAuthTypeAuthKeyKey, looklookUsercenterUserAuthIdKey, looklookUsercenterUserAuthUserIdAuthTypeKey)
	return ret, err
}

func (m *defaultUserAuthModel) TransInsert(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error) {
	looklookUsercenterUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	looklookUsercenterUserAuthIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserAuthIdPrefix, data.Id)
	looklookUsercenterUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userAuthRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.Version, data.UserId, data.AuthKey, data.AuthType)
	}, looklookUsercenterUserAuthAuthTypeAuthKeyKey, looklookUsercenterUserAuthIdKey, looklookUsercenterUserAuthUserIdAuthTypeKey)
	return ret, err
}
func (m *defaultUserAuthModel) Update(ctx context.Context, newData *UserAuth) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	looklookUsercenterUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	looklookUsercenterUserAuthIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserAuthIdPrefix, data.Id)
	looklookUsercenterUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAuthRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.UserId, newData.AuthKey, newData.AuthType, newData.Id)
	}, looklookUsercenterUserAuthAuthTypeAuthKeyKey, looklookUsercenterUserAuthIdKey, looklookUsercenterUserAuthUserIdAuthTypeKey)
	return err
}

func (m *defaultUserAuthModel) TransUpdate(ctx context.Context, session sqlx.Session, newData *UserAuth) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	looklookUsercenterUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	looklookUsercenterUserAuthIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserAuthIdPrefix, data.Id)
	looklookUsercenterUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheLooklookUsercenterUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAuthRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.UserId, newData.AuthKey, newData.AuthType, newData.Id)
	}, looklookUsercenterUserAuthAuthTypeAuthKeyKey, looklookUsercenterUserAuthIdKey, looklookUsercenterUserAuthUserIdAuthTypeKey)
	return err
}

func (m *defaultUserAuthModel) List(ctx context.Context, page, limit int64) ([]*UserAuth, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", userAuthRows, m.table)
	var resp []*UserAuth
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultUserAuthModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUserAuthModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultUserAuthModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultUserAuthModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*UserAuth, error) {

	builder = builder.Columns(userAuthRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAuth
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAuth, error) {

	builder = builder.Columns(userAuthRows)

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

	var resp []*UserAuth
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserAuth, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(userAuthRows)

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

	var resp []*UserAuth
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultUserAuthModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserAuth, error) {

	builder = builder.Columns(userAuthRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAuth
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserAuth, error) {

	builder = builder.Columns(userAuthRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserAuth
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultUserAuthModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLooklookUsercenterUserAuthIdPrefix, primary)
}

func (m *defaultUserAuthModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAuthRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserAuthModel) tableName() string {
	return m.table
}
