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

var _ TaskModel = (*customTaskModel)(nil)

type (
	// TaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskModel.
	TaskModel interface {
		taskModel

		FindUserConTaskDone(ctx context.Context, userId int64, days int64) ([]*Task, error)
		FindUserTaskList(ctx context.Context, userId int64) ([]*TaskDoneList, error)
	}

	customTaskModel struct {
		*defaultTaskModel
	}
)

// NewTaskModel returns a model for the database table.
func NewTaskModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TaskModel {
	return &customTaskModel{
		defaultTaskModel: newTaskModel(conn, c, opts...),
	}
}

type TaskDoneList struct {
	Id         int64     `db:"id"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
	Title      string    `db:"title"`       // 任务标题
	Type       int64     `db:"type"`        // 任务类型：1连续签到
	Value      int64     `db:"value"`       // 任务值，依据type来确定：为1时是连续天数
	Content    string    `db:"content"`     // 备注
	Points     int64     `db:"points"`      // 获得的积分值
	IsFinished int64     `db:"is_finished"` // 1为完成
}

func (m *customTaskModel) FindUserTaskList(ctx context.Context, userId int64) ([]*TaskDoneList, error) {
	var query string
	query = fmt.Sprintf(`SELECT t.*,
		CASE 
			WHEN tf.user_id = ? THEN 1
			ELSE 0
		END AS is_finished
	FROM 
		task t
	LEFT JOIN 
		task_finish tf
	ON 
    	t.id = tf.task_id AND tf.user_id = ?`)
	var resp []*TaskDoneList

	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId, userId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindUserTaskList, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
}

func (m *customTaskModel) FindUserConTaskDone(ctx context.Context, userId int64, days int64) ([]*Task, error) {
	var query string
	query = fmt.Sprintf(`select %s from %s as t1 where t1.value <= ? and t1.type = 1 
    AND NOT EXISTS (
      SELECT *
      FROM task_finish AS tf 
      WHERE tf.user_id = ?
        AND tf.task_id = t1.id
	) 
	order by t1.id`, taskRows, m.table)
	var resp []*Task

	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, days, userId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "FindUserConTaskDone, &resp:%v, query:%v, userId:%v, error: %v", &resp, query, userId, err)
	}
}
