package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/lottery/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/constants"
	"looklook/common/kqueue"
	"looklook/common/wxnotice"
	"looklook/common/xerr"
	"math/rand"
	"time"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnnounceLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnnounceLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnnounceLotteryLogic {
	return &AnnounceLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LotteryStrategy 定义抽奖策略接口
type LotteryStrategy interface {
	Run() error
}

// TimeLotteryStrategy 实现基于时间的抽奖策略
type TimeLotteryStrategy struct {
	*AnnounceLotteryLogic
	CurrentTime time.Time
}

// PeopleLotteryStrategy 实现基于人数的抽奖策略
type PeopleLotteryStrategy struct {
	*AnnounceLotteryLogic
	CurrentTime time.Time
}

type Winner struct {
	LotteryId int64
	UserId    int64
	PrizeId   int64
}

func (l *AnnounceLotteryLogic) AnnounceLottery(in *pb.AnnounceLotteryReq) (*pb.AnnounceLotteryResp, error) {
	// 创建相应的抽奖策略
	var strategy LotteryStrategy
	switch in.AnnounceType {
	case constants.AnnounceTypeTimeLottery:
		// 开奖时间类型
		strategy = &TimeLotteryStrategy{
			AnnounceLotteryLogic: l,
			CurrentTime:          time.Now(),
		}
	case constants.AnnounceTypePeopleLottery:
		// 开奖人数类型
		strategy = &PeopleLotteryStrategy{
			AnnounceLotteryLogic: l,
			CurrentTime:          time.Now(),
		}
	}
	err := strategy.Run()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("抽奖运行错误"), "AnnounceStrategy run error: %v", err)
	}
	//fmt.Println("AnnounceFinish") // t
	return &pb.AnnounceLotteryResp{}, nil
}

// DrawLottery 实现抽奖策略
func (l *AnnounceLotteryLogic) DrawLottery(ctx context.Context, lotteryId int64, prizes []*model.Prize, participantor []int64) ([]Winner, error) {

	if participantor == nil || len(participantor) == 0 {
		return nil, nil
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	// todo 实现可以动态调整的抽奖策略
	// 使用洗牌算法打乱参与人列表
	for i := len(participantor) - 1; i > 0; i-- {
		j := rand.Intn(i)
		participantor[i], participantor[j] = participantor[j], participantor[i]
	}

	// 计算出奖品总数
	prizesSum := 0
	for _, prize := range prizes {
		prizesSum += int(prize.Count)
	}

	winners := make([]Winner, prizesSum) // 保存中奖者
	prizesIndex := 0                     // 遍历奖品的下标

	for i := 0; i < prizesSum; i++ {

		// 抽奖的参与人数小于奖品数
		if i >= len(participantor) {
			break
		}

		winners[i] = Winner{
			LotteryId: lotteryId,
			UserId:    participantor[i],
			PrizeId:   prizes[prizesIndex].Id,
		}

		// 该下标奖品数减一  该奖品数为0时  奖品下标++
		prizes[prizesIndex].Count--
		if prizes[prizesIndex].Count == 0 {
			prizesIndex++
		}
	}

	return winners, nil
}

// NotifyParticipators 通知MQ队列
func (l *AnnounceLotteryLogic) NotifyParticipators(participators []int64, lotteryId int64) error {
	fmt.Println("NotifyParticipators", participators, lotteryId)

	now := time.Now()
	msg := wxnotice.MessageReminder{
		ReminderContent: wxnotice.Item{Value: "开奖结束，快来看看你有没有中奖"},
		ReminderTime:    wxnotice.Item{Value: now.Format("2006-01-02 15:04")},
	}
	data, _ := power.StructToHashMap(&msg)

	// 拼接小程序页面地址
	pageAddr := fmt.Sprintf("pages/lottery/my/index")

	for _, userId := range participators {
		userOpenId, err := l.svcCtx.UsercenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
			UserId:   userId,
			AuthType: constants.UserAuthTypeSmallWX,
		})
		if err != nil {
			continue
		}
		m := kqueue.WXMiniNoticeSendMessage{
			ToUser: userOpenId.UserAuth.AuthKey,
			Page:   pageAddr,
			UserId: userId,
			Data:   data,
		}
		body, err := json.Marshal(m)
		_ = l.svcCtx.KqueueNoticeSendClient.Push(l.ctx, string(body))
	}

	return nil
}

// WriteWinnersToLotteryParticipation 更新参与抽奖表
func (l *AnnounceLotteryLogic) WriteWinnersToLotteryParticipation(winners []Winner) error {
	for _, w := range winners {
		if w.UserId == 0 {
			continue
		}
		err := l.svcCtx.LotteryParticipationModel.UpdateWinners(l.ctx, w.LotteryId, w.UserId, w.PrizeId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *TimeLotteryStrategy) Run() error {
	// 查询满足条件的抽奖
	lotteries, err := s.svcCtx.LotteryModel.GetLotterysByLessThanCurrentTime(s.ctx, s.CurrentTime, constants.AnnounceTypeTimeLottery)
	if err != nil && err != model.ErrNotFound {
		return err
	}
	if err == model.ErrNotFound || len(lotteries) == 0 {
		return nil
	}

	// 遍历每一个抽奖
	for _, lotteryId := range lotteries {
		var participators []int64
		// 事务处理
		err = s.svcCtx.LotteryModel.Trans(s.ctx, func(context context.Context, session sqlx.Session) error {
			//根据抽奖id得到对应的所有奖品
			prizes, err := s.svcCtx.PrizeModel.PrizeListByLotteryId(s.ctx, lotteryId)
			if err != nil {
				return err
			}

			//fmt.Println("开始开奖的lottery:", lotteryId)

			// 获取该lotteryId对应的所有参与者
			participators, err = s.svcCtx.LotteryParticipationModel.GetParticipationUserIdsByLotteryId(s.ctx, lotteryId)
			if err != nil && err != model.ErrNotFound {
				return err
			}

			winners, err := s.DrawLottery(s.ctx, lotteryId, prizes, participators)
			if err != nil {
				return errors.Wrapf(xerr.NewErrMsg("开奖错误"), "DrawLottery,lotteryId:%v,prizes:%v,participators:%v error: %v", lotteryId, prizes, participators, err)
			}

			// 测试查看所有winners
			for _, w := range winners {
				fmt.Printf("testwinners:%+v\n", w)
			}

			//更新抽奖状态为"已开奖"
			err = s.svcCtx.LotteryModel.UpdateLotteryStatus(s.ctx, lotteryId)
			if err != nil {
				return err
			}

			// 将得到的中奖信息，写入数据库participants
			err = s.WriteWinnersToLotteryParticipation(winners)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("开奖事务错误"), "AnnounceLotteryTrans error: %v", err)
		}

		// 执行开奖结果通知任务
		err := s.NotifyParticipators(participators, lotteryId)
		if err != nil {
			return err
		}
	}
	return nil
}

// Run 按人数开奖策略
func (s *PeopleLotteryStrategy) Run() error {

	// 查询开奖类型为2并 且没有开奖 且已开始的所有抽奖
	lotteries, err := s.svcCtx.LotteryModel.GetLotterysByLessThanCurrentTime2(s.ctx, s.CurrentTime, constants.AnnounceTypePeopleLottery)
	if err != nil {
		return err
	}

	CheckedLottery, err := s.CheckLottery(lotteries)
	if err != nil {
		return err
	}
	// 遍历每一个抽奖
	for _, lottery := range CheckedLottery {
		var participators []int64
		// 事务处理
		err = s.svcCtx.LotteryModel.Trans(s.ctx, func(context context.Context, session sqlx.Session) error {
			//根据抽奖id得到对应的所有奖品
			prizes, err := s.svcCtx.PrizeModel.PrizeListByLotteryId(s.ctx, lottery.Id)
			if err != nil {
				return err
			}

			//fmt.Println("开始开奖的lottery:", lottery.Id)

			// 获取该lotteryId对应的所有参与者
			participators, err = s.svcCtx.LotteryParticipationModel.GetParticipationUserIdsByLotteryId(s.ctx, lottery.Id)
			if err != nil && err != model.ErrNotFound {
				return err
			}

			//testParticipators := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			//participators = testParticipators
			//fmt.Println("participators:", participators)

			winners, err := s.DrawLottery(s.ctx, lottery.Id, prizes, participators)
			if err != nil {
				return errors.Wrapf(xerr.NewErrMsg("抽奖策略错误"), "DrawLottery,lotteryId:%v,prizes:%v,participators:%v, error: %v", lottery.Id, prizes, participators, err)
			}

			//测试查看所有winners
			for _, w := range winners {
				fmt.Printf("testwinners:%+v\n", w)
			}

			//更新抽奖状态为"已开奖"
			err = s.svcCtx.LotteryModel.UpdateLotteryStatus(s.ctx, lottery.Id)
			if err != nil {
				return err
			}

			// 将得到的中奖信息，写入数据库participants
			err = s.WriteWinnersToLotteryParticipation(winners)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrMsg("人数开奖错误"), "AnnounceLotteryTrans error: %v", err)
		}

		// 执行开奖结果通知任务
		err := s.NotifyParticipators(participators, lottery.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *PeopleLotteryStrategy) CheckLottery(lotteries []*model.Lottery) (CheckedLotterys []*model.Lottery, err error) {
	// 筛选满足条件的抽奖
	// 1. 超过当前时间的，即使没有满足人数条件也需要进行开奖
	// 2. 当参与人数 >= 开奖人数，进行开奖

	for _, l := range lotteries {
		// l.AnnounceTime 小于等于 s.CurrentTime,即超过当前时间，需要开奖
		if l.AwardDeadline.Before(s.CurrentTime) || l.AwardDeadline.Equal(s.CurrentTime) {
			CheckedLotterys = append(CheckedLotterys, l)
		} else {
			//fmt.Println("lotteryId:", l.Id)
			ParticipatorCount, err := s.svcCtx.LotteryParticipationModel.GetParticipatorsCountByLotteryId(s.ctx, l.Id)
			//fmt.Println("ParticipatorCount:", ParticipatorCount)
			if err != nil {
				return nil, err
			}
			// 检查参与人数是否达到指定人数
			if ParticipatorCount >= l.JoinNumber {
				CheckedLotterys = append(CheckedLotterys, l)
			}
		}
	}
	return
}
