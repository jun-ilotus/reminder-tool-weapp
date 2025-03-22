package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UploadFileModel = (*customUploadFileModel)(nil)

type (
	// UploadFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUploadFileModel.
	UploadFileModel interface {
		uploadFileModel
	}

	customUploadFileModel struct {
		*defaultUploadFileModel
	}
)

// NewUploadFileModel returns a model for the database table.
func NewUploadFileModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UploadFileModel {
	return &customUploadFileModel{
		defaultUploadFileModel: newUploadFileModel(conn, c, opts...),
	}
}
