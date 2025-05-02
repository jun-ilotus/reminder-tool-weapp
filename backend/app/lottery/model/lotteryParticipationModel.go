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

var _ LotteryParticipationModel = (*customLotteryParticipationModel)(nil)

type (
	// LotteryParticipationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryParticipationModel.
	LotteryParticipationModel interface {
		lotteryParticipationModel

		GetLotteryParticipationByUserIdLotteryId(ctx context.Context, userId, lotteryId int64) (*LotteryParticipation, error)
		GetLotteryParticipationCount(ctx context.Context, lotteryId, isWin int64) (int64, error)
		GetLotteryParticipationLastId(ctx context.Context, lotteryId int64) (int64, error)
		LotteryParticipationList(ctx context.Context, lotteryId, limit, lastId int64) ([]*LotteryParticipation, error)
		LotteryParticipationWinList(ctx context.Context, lotteryId int64) ([]*LotteryParticipation, error)
		LotteryParticipationUserList(ctx context.Context, userId, isWon, isAnnounced int64) ([]*UserLottery, error)
	}

	customLotteryParticipationModel struct {
		*defaultLotteryParticipationModel
	}
)

// NewLotteryParticipationModel returns a model for the database table.
func NewLotteryParticipationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryParticipationModel {
	return &customLotteryParticipationModel{
		defaultLotteryParticipationModel: newLotteryParticipationModel(conn, c, opts...),
	}
}

func (c *customLotteryParticipationModel) GetLotteryParticipationByUserIdLotteryId(ctx context.Context, userId, lotteryId int64) (*LotteryParticipation, error) {
	var resp LotteryParticipation
	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? AND lottery_id = ? LIMIT 1", lotteryParticipationRows, c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, userId, lotteryId)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationByUserIdLotteryId, error: %v", err)
	}
}

func (c *customLotteryParticipationModel) GetLotteryParticipationCount(ctx context.Context, lotteryId, isWin int64) (int64, error) {
	var resp int64

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE lottery_id = ?", c.table)
	if isWin == 1 {
		query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE lottery_id = ? AND is_won = 1", c.table)
	}
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, lotteryId)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationCount, error: %v", err)
	}
	return resp, nil
}

func (c *customLotteryParticipationModel) GetLotteryParticipationLastId(ctx context.Context, lotteryId int64) (int64, error) {
	var resp int64
	query := fmt.Sprintf("SELECT id FROM %s WHERE lottery_id = ? ORDER BY id DESC LIMIT 1", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, lotteryId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetLotteryParticipationLastId, error: %v", err)
	}
}

func (c *customLotteryParticipationModel) LotteryParticipationList(ctx context.Context, lotteryId, limit, lastId int64) ([]*LotteryParticipation, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where lottery_id = ? AND id < ? order by id desc limit ?", lotteryParticipationRows, c.table)

	var resp []*LotteryParticipation
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lotteryId, lastId, limit)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryParticipationList, &resp:%v, query:%v, lastId:%v, limit:%v, error: %v", &resp, query, lastId, limit, err)
	}
}

func (c *customLotteryParticipationModel) LotteryParticipationWinList(ctx context.Context, lotteryId int64) ([]*LotteryParticipation, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where lottery_id = ? and is_won = 1 order by prize_id", lotteryParticipationRows, c.table)

	var resp []*LotteryParticipation
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lotteryId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryParticipationList, &resp:%v, query:%v, error: %v", &resp, query, err)
	}
}

type UserLottery struct {
	LotteryId     int64     `db:"lottery_id"`     // 参与的抽奖的id
	Name          string    `db:"name"`           // 默认取一等奖名称
	Thumb         string    `db:"thumb"`          // 默认取一等经配图
	PublishTime   time.Time `db:"publish_time"`   // 发布抽奖时间
	JoinNumber    int64     `db:"join_number"`    // 自动开奖人数
	AwardDeadline time.Time `db:"award_deadline"` // 领奖截止时间
	AnnounceType  int64     `db:"announce_type"`  // 开奖设置：1按时间开奖 2按人数开奖 3即抽即中
	AnnounceTime  time.Time `db:"announce_time"`  // 开奖时间
	IsAnnounced   int64     `db:"is_announced"`   // 是否开奖：0未开奖；1已经开奖

	LotteryParticipationId int64     `db:"id"`       // 主键
	UserId                 int64     `db:"user_id"`  // 用户id
	IsWon                  int64     `db:"is_won"`   // 中奖了吗？
	PrizeId                int64     `db:"prize_id"` // 奖品id
	UpdateTime             time.Time `db:"update_time"`
}

func (c *customLotteryParticipationModel) LotteryParticipationUserList(ctx context.Context, userId, isWon, isAnnounced int64) ([]*UserLottery, error) {
	var query string
	if isWon == 0 { // 全部的参与
		query = fmt.Sprintf(`
select p.lottery_id, l.name, l.thumb, l.publish_time, l.join_number, l.award_deadline, l.announce_type, l.announce_time, l.is_announced, p.id, p.user_id, p.is_won, p.prize_id, p.update_time  
from lottery_participation as p
	LEFT JOIN lottery as l
		ON p.lottery_id = l.id
where p.user_id = ? AND l.is_announced = ?
order by p.id desc 
`)
	} else if isWon == 1 { // 中奖的
		query = fmt.Sprintf(`
select p.lottery_id, l.name, l.thumb, l.publish_time, l.join_number, l.award_deadline, l.announce_type, l.announce_time, l.is_announced, p.id, p.user_id, p.is_won, p.prize_id, p.update_time  
from lottery_participation as p
	LEFT JOIN lottery as l
		ON p.lottery_id = l.id
where p.user_id = ? AND p.is_won = 1 AND l.is_announced = ?
order by p.id desc 
`)
	}

	var resp []*UserLottery
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, userId, isAnnounced)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "LotteryParticipationUserList, &resp:%v, query:%v, error: %v", &resp, query, err)
	}
}
