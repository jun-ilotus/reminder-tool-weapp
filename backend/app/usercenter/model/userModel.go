package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel

		ModifyIntimateId(ctx context.Context, userId, intimateId int64) error
		FindUserById(ctx context.Context, userId int64) (*User, error)
		TransUpdatePoints(ctx context.Context, session sqlx.Session, userId, points int64) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (c *customUserModel) ModifyIntimateId(ctx context.Context, userId, intimateId int64) error {
	query := fmt.Sprintf("update %s set `intimate_id` = ? where `id` = ?", c.table)

	_, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.Exec(query, intimateId, userId)
	})
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "DB ModifyIntimateId, intimateId:%v error: %v", intimateId, err)
	}
	return nil
}

func (c *customUserModel) FindUserById(ctx context.Context, userId int64) (*User, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where id = ?", userRows, c.table)
	var resp User

	err := c.QueryRowNoCacheCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindUserById, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
	return &resp, nil
}

func (c *customUserModel) TransUpdatePoints(ctx context.Context, session sqlx.Session, userId, points int64) error {
	looklookUsercenterPointsRecodeIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterPointsRecodeIdPrefix, userId)
	_, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set points = points + ? where `id` = ?", c.table)
		return session.ExecCtx(ctx, query, points, userId)
	}, looklookUsercenterPointsRecodeIdKey)
	return err
}
