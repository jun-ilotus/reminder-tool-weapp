package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ ItemsModel = (*customItemsModel)(nil)

type (
	// ItemsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customItemsModel.
	ItemsModel interface {
		itemsModel

		// 自定义方法
		GetLastId(ctx context.Context) (int64, error)
		ItemsList(ctx context.Context, userId int64) ([]*Items, error)
		ItemsIntimateList(ctx context.Context, userId int64) ([]*Items, error)
	}

	customItemsModel struct {
		*defaultItemsModel
	}
)

// NewItemsModel returns a model for the database table.
func NewItemsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ItemsModel {
	return &customItemsModel{
		defaultItemsModel: newItemsModel(conn, c, opts...),
	}
}

func (c *customItemsModel) GetLastId(ctx context.Context) (int64, error) {
	var resp int64
	query := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLastId, error: %v", err)
	}
	return resp, nil
}

func (c *customItemsModel) ItemsList(ctx context.Context, userId int64) ([]*Items, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where user_id = ? order by create_time desc", itemsRows, c.table)
	var resp []*Items
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
	return resp, nil
}

func (c *customItemsModel) ItemsIntimateList(ctx context.Context, userId int64) ([]*Items, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where user_id = ? and member = 1 order by create_time desc", itemsRows, c.table)
	var resp []*Items
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
	return resp, nil
}
