package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NoticeLogModel = (*customNoticeLogModel)(nil)

type (
	// NoticeLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNoticeLogModel.
	NoticeLogModel interface {
		noticeLogModel
	}

	customNoticeLogModel struct {
		*defaultNoticeLogModel
	}
)

// NewNoticeLogModel returns a model for the database table.
func NewNoticeLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NoticeLogModel {
	return &customNoticeLogModel{
		defaultNoticeLogModel: newNoticeLogModel(conn, c, opts...),
	}
}
