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
	uploadFileFieldNames          = builder.RawFieldNames(&UploadFile{})
	uploadFileRows                = strings.Join(uploadFileFieldNames, ",")
	uploadFileRowsExpectAutoSet   = strings.Join(stringx.Remove(uploadFileFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	uploadFileRowsWithPlaceHolder = strings.Join(stringx.Remove(uploadFileFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUploadUploadFileIdPrefix = "cache:upload:uploadFile:id:"
)

type (
	uploadFileModel interface {
		Insert(ctx context.Context, data *UploadFile) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *UploadFile) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UploadFile, error)
		Update(ctx context.Context, data *UploadFile) error
		List(ctx context.Context, page, limit int64) ([]*UploadFile, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *UploadFile) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*UploadFile, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UploadFile, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UploadFile, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UploadFile, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UploadFile, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultUploadFileModel struct {
		sqlc.CachedConn
		table string
	}

	UploadFile struct {
		Id         int64          `db:"id"`
		UserId     int64          `db:"user_id"`   // 上传用户id
		FileName   sql.NullString `db:"file_name"` // 文件名
		Ext        sql.NullString `db:"ext"`       // 扩展名
		Size       sql.NullInt64  `db:"size"`      // 文件大小
		Url        string         `db:"url"`       // 下载链接
		CreateTime time.Time      `db:"create_time"`
		UpdateTime time.Time      `db:"update_time"`
		DeleteTime sql.NullTime   `db:"delete_time"`
	}
)

func newUploadFileModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUploadFileModel {
	return &defaultUploadFileModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`upload_file`",
	}
}

func (m *defaultUploadFileModel) Delete(ctx context.Context, id int64) error {
	uploadUploadFileIdKey := fmt.Sprintf("%s%v", cacheUploadUploadFileIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, uploadUploadFileIdKey)
	return err
}

func (m *defaultUploadFileModel) FindOne(ctx context.Context, id int64) (*UploadFile, error) {
	uploadUploadFileIdKey := fmt.Sprintf("%s%v", cacheUploadUploadFileIdPrefix, id)
	var resp UploadFile
	err := m.QueryRowCtx(ctx, &resp, uploadUploadFileIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", uploadFileRows, m.table)
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

func (m *defaultUploadFileModel) Insert(ctx context.Context, data *UploadFile) (sql.Result, error) {
	uploadUploadFileIdKey := fmt.Sprintf("%s%v", cacheUploadUploadFileIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, uploadFileRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.FileName, data.Ext, data.Size, data.Url, data.DeleteTime)
	}, uploadUploadFileIdKey)
	return ret, err
}

func (m *defaultUploadFileModel) TransInsert(ctx context.Context, session sqlx.Session, data *UploadFile) (sql.Result, error) {
	uploadUploadFileIdKey := fmt.Sprintf("%s%v", cacheUploadUploadFileIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, uploadFileRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.FileName, data.Ext, data.Size, data.Url, data.DeleteTime)
	}, uploadUploadFileIdKey)
	return ret, err
}
func (m *defaultUploadFileModel) Update(ctx context.Context, data *UploadFile) error {
	uploadUploadFileIdKey := fmt.Sprintf("%s%v", cacheUploadUploadFileIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, uploadFileRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.FileName, data.Ext, data.Size, data.Url, data.DeleteTime, data.Id)
	}, uploadUploadFileIdKey)
	return err
}

func (m *defaultUploadFileModel) TransUpdate(ctx context.Context, session sqlx.Session, data *UploadFile) error {
	uploadUploadFileIdKey := fmt.Sprintf("%s%v", cacheUploadUploadFileIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, uploadFileRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.FileName, data.Ext, data.Size, data.Url, data.DeleteTime, data.Id)
	}, uploadUploadFileIdKey)
	return err
}

func (m *defaultUploadFileModel) List(ctx context.Context, page, limit int64) ([]*UploadFile, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", uploadFileRows, m.table)
	var resp []*UploadFile
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultUploadFileModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUploadFileModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultUploadFileModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultUploadFileModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*UploadFile, error) {

	builder = builder.Columns(uploadFileRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UploadFile
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUploadFileModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UploadFile, error) {

	builder = builder.Columns(uploadFileRows)

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

	var resp []*UploadFile
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUploadFileModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UploadFile, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(uploadFileRows)

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

	var resp []*UploadFile
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultUploadFileModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UploadFile, error) {

	builder = builder.Columns(uploadFileRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UploadFile
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUploadFileModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UploadFile, error) {

	builder = builder.Columns(uploadFileRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UploadFile
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUploadFileModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultUploadFileModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUploadUploadFileIdPrefix, primary)
}

func (m *defaultUploadFileModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", uploadFileRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUploadFileModel) tableName() string {
	return m.table
}
