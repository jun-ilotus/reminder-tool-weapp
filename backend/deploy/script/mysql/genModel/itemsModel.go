package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ItemsModel = (*customItemsModel)(nil)

type (
	// ItemsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customItemsModel.
	ItemsModel interface {
		itemsModel
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
