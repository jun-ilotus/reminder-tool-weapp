package recode

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/signin/cmd/rpc/signin"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"time"

	"looklook/app/signin/cmd/api/internal/svc"
	"looklook/app/signin/cmd/api/internal/types"

	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero" // 这行导入gozero的dtm驱动
)

var dtmServer = "etcd://etcd:2379/dtmservice"

type AddTodayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTodayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTodayLogic {
	return &AddTodayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTodayLogic) AddToday(req *types.AddRecodeTodayReq) (resp *types.AddRecodeTodayResp, err error) {
	//userId := ctxdata.GetUidFromCtx(l.ctx)
	userId := int64(2)

	signinRpcBusiServer, err := l.svcCtx.Config.SigninRpcConf.BuildTarget()
	if err != nil {
		return nil, err
	}
	usercenterRpcBusiServer, err := l.svcCtx.Config.UsercenterRpcConf.BuildTarget()
	if err != nil {
		return nil, err
	}
	addRecodeReq := &signin.AddRecodeReq{UserId: userId, SignDate: time.Now().Unix()}

	// 开始分布式事务 tcc
	gid := dtmgrpc.MustGenGid(dtmServer)
	err = dtmgrpc.TccGlobalTransaction(dtmServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
		addRecodeResp := &signin.AddRecodeResp{}
		addPointsRecodeResp := &usercenter.AddPointsRecodeResp{}
		addFinishResp := &signin.AddTaskFinishResp{}
		err := tcc.CallBranch(addRecodeReq, signinRpcBusiServer+"/pb.signin/AddRecode",
			"", signinRpcBusiServer+"/pb.signin/AddRecodeRollback", addRecodeResp)
		if err != nil {
			return err
		}
		addTaskFinishReq := &signin.AddTaskFinishReq{Task: addRecodeResp.Task, UserId: userId}
		err = tcc.CallBranch(addTaskFinishReq, signinRpcBusiServer+"/pb.signin/AddTaskFinish",
			"", signinRpcBusiServer+"/pb.signin/AddTaskFinishRollback", addFinishResp)
		if err != nil {
			return err
		}
		content := ""
		if len(addRecodeResp.Task) > 0 {
			content = addRecodeResp.Task[0].Content
		}
		addPointsRecodeReq := &usercenter.AddPointsRecodeReq{
			UserId:  userId,
			Points:  addFinishResp.Points,
			Content: content,
		}
		if addFinishResp.Points != 0 {
			err = tcc.CallBranch(addPointsRecodeReq, usercenterRpcBusiServer+"/pb.usercenter/AddPointsRecode",
				"", usercenterRpcBusiServer+"/pb.usercenter/AddPointsRecodeRollback", addPointsRecodeResp)
			if err != nil {
				return err
			}
		}
		return nil
	})

	//saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
	//	Add(signinRpcBusiServer+"/pb.signin/AddRecode", signinRpcBusiServer+"/pb.signin/AddRecodeRollback", addRecodeReq).
	//	Add(usercenterRpcBusiServer+"/pb.usercenter/AddPointsRecode", usercenterRpcBusiServer+"/pb.usercenter/AddPointsRecodeRollback", addPointsRecodeReq)
	//saga.WaitResult = true
	//err = saga.Submit()
	if err != nil {
		return nil, fmt.Errorf("submit data to  dtm-server err  : %+v \n", err)
	}
	return &types.AddRecodeTodayResp{}, nil
}
