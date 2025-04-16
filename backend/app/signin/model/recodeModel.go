package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
	"time"
)

var _ RecodeModel = (*customRecodeModel)(nil)

type (
	// RecodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRecodeModel.
	RecodeModel interface {
		recodeModel

		FindCountUserNowConRecodeDays(ctx context.Context, userId int64) (*int, error)
		FindLastOneByUserIdSignDate(ctx context.Context, userId int64, SignDate time.Time) (*Recode, error)
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

func (c *customRecodeModel) FindLastOneByUserIdSignDate(ctx context.Context, userId int64, SignDate time.Time) (*Recode, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where user_id = ? and sign_date = ? order by id desc limit 1", recodeRows, c.table)
	var resp Recode

	err := c.QueryRowNoCacheCtx(ctx, &resp, query, userId, SignDate)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindLastOneByUserIdSignDate, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
}

func (c *customRecodeModel) FindCountUserNowConRecodeDays(ctx context.Context, userId int64) (*int, error) {
	var query string
	query = `SELECT 
    COALESCE(MAX(consecutive_days), 0) AS max_consecutive_days
FROM (
    SELECT 
        user_id,
        sign_date,
        IF(@prev_user = user_id AND DATEDIFF(sign_date, @prev_date) = 1, 
           @consecutive_days := @consecutive_days + 1,
           IF(@prev_user = user_id AND DATEDIFF(sign_date, @prev_date) > 1, 
              @consecutive_days := 1,
              @consecutive_days := 1)) AS consecutive_days,
        @prev_date := sign_date,
        @prev_user := user_id
    FROM 
        recode,
        (SELECT @prev_date := NULL, @prev_user := NULL, @consecutive_days := 0) AS vars
    WHERE 
        user_id = ? AND
        sign_date <= CURDATE()
    ORDER BY 
        sign_date DESC
) AS subquery;`
	var resp int

	err := c.QueryRowNoCacheCtx(ctx, &resp, query, userId)

	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindLastOneByUserIdTaskId, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
}
