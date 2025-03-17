package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PrizeModel = (*customPrizeModel)(nil)

type (
	// PrizeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeModel.
	PrizeModel interface {
		prizeModel
		//自定义的方法写道这里，避免覆盖
		TransInsert(ctx context.Context, session sqlx.Session, data *Prize) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
	}

	customPrizeModel struct {
		*defaultPrizeModel
	}
)

// NewPrizeModel returns a model for the database table.
func NewPrizeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PrizeModel {
	return &customPrizeModel{
		defaultPrizeModel: newPrizeModel(conn, c, opts...),
	}
}

func (m *defaultPrizeModel) TransInsert(ctx context.Context, session sqlx.Session, data *Prize) (sql.Result, error) {
	lotteryPrizeIdKey := fmt.Sprintf("%s%v", cacheLotteryPrizeIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, prizeRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.LotteryId, data.Type, data.Name, data.Level, data.Thumb, data.Count, data.GrantType)
	}, lotteryPrizeIdKey)
	return ret, err
}

func (m *defaultPrizeModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}
