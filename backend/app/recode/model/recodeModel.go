package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ RecodeModel = (*customRecodeModel)(nil)

type (
	// RecodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecodeModel.
	RecodeModel interface {
		recodeModel

		// 自定义方法
		GetLastId(ctx context.Context) (int64, error)
		RecodeList(ctx context.Context, userId, itemsId int64) ([]*Recode, error)
	}

	customRecodeModel struct {
		*defaultRecodeModel
	}
)

// NewRecodeModel returns a model for the database table.
func NewRecodeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RecodeModel {
	return &customRecodeModel{
		defaultRecodeModel: newRecodeModel(conn, c, opts...),
	}
}

func (c *customRecodeModel) GetLastId(ctx context.Context) (int64, error) {
	var resp int64
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLastId, error: %v", err)
	}
	return resp, nil
}

func (c *customRecodeModel) RecodeList(ctx context.Context, userId, itemsId int64) ([]*Recode, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where user_id = ? AND items_id = ? order by recode_time desc", recodeRows, c.table)
	var resp []*Recode
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId, itemsId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
	return resp, nil
}
