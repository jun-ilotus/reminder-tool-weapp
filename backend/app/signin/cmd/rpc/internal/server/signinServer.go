// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: signin.proto

package server

import (
	"context"

	"looklook/app/signin/cmd/rpc/internal/logic"
	"looklook/app/signin/cmd/rpc/internal/svc"
	"looklook/app/signin/cmd/rpc/pb"
)

type SigninServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSigninServer
}

func NewSigninServer(svcCtx *svc.ServiceContext) *SigninServer {
	return &SigninServer{
		svcCtx: svcCtx,
	}
}

// -----------------------recode-----------------------
func (s *SigninServer) AddRecode(ctx context.Context, in *pb.AddRecodeReq) (*pb.AddRecodeResp, error) {
	l := logic.NewAddRecodeLogic(ctx, s.svcCtx)
	return l.AddRecode(in)
}

func (s *SigninServer) AddRecodeRollback(ctx context.Context, in *pb.AddRecodeReq) (*pb.AddRecodeResp, error) {
	l := logic.NewAddRecodeRollbackLogic(ctx, s.svcCtx)
	return l.AddRecodeRollback(in)
}

func (s *SigninServer) UpdateRecode(ctx context.Context, in *pb.UpdateRecodeReq) (*pb.UpdateRecodeResp, error) {
	l := logic.NewUpdateRecodeLogic(ctx, s.svcCtx)
	return l.UpdateRecode(in)
}

func (s *SigninServer) DelRecode(ctx context.Context, in *pb.DelRecodeReq) (*pb.DelRecodeResp, error) {
	l := logic.NewDelRecodeLogic(ctx, s.svcCtx)
	return l.DelRecode(in)
}

func (s *SigninServer) GetRecodeById(ctx context.Context, in *pb.GetRecodeByIdReq) (*pb.GetRecodeByIdResp, error) {
	l := logic.NewGetRecodeByIdLogic(ctx, s.svcCtx)
	return l.GetRecodeById(in)
}

func (s *SigninServer) SearchRecode(ctx context.Context, in *pb.SearchRecodeReq) (*pb.SearchRecodeResp, error) {
	l := logic.NewSearchRecodeLogic(ctx, s.svcCtx)
	return l.SearchRecode(in)
}

// -----------------------task-----------------------
func (s *SigninServer) AddTask(ctx context.Context, in *pb.AddTaskReq) (*pb.AddTaskResp, error) {
	l := logic.NewAddTaskLogic(ctx, s.svcCtx)
	return l.AddTask(in)
}

func (s *SigninServer) UpdateTask(ctx context.Context, in *pb.UpdateTaskReq) (*pb.UpdateTaskResp, error) {
	l := logic.NewUpdateTaskLogic(ctx, s.svcCtx)
	return l.UpdateTask(in)
}

func (s *SigninServer) DelTask(ctx context.Context, in *pb.DelTaskReq) (*pb.DelTaskResp, error) {
	l := logic.NewDelTaskLogic(ctx, s.svcCtx)
	return l.DelTask(in)
}

func (s *SigninServer) GetTaskById(ctx context.Context, in *pb.GetTaskByIdReq) (*pb.GetTaskByIdResp, error) {
	l := logic.NewGetTaskByIdLogic(ctx, s.svcCtx)
	return l.GetTaskById(in)
}

func (s *SigninServer) SearchTask(ctx context.Context, in *pb.SearchTaskReq) (*pb.SearchTaskResp, error) {
	l := logic.NewSearchTaskLogic(ctx, s.svcCtx)
	return l.SearchTask(in)
}

// -----------------------taskFinish-----------------------
func (s *SigninServer) AddTaskFinish(ctx context.Context, in *pb.AddTaskFinishReq) (*pb.AddTaskFinishResp, error) {
	l := logic.NewAddTaskFinishLogic(ctx, s.svcCtx)
	return l.AddTaskFinish(in)
}

func (s *SigninServer) AddTaskFinishRollback(ctx context.Context, in *pb.AddTaskFinishReq) (*pb.AddTaskFinishResp, error) {
	l := logic.NewAddTaskFinishRollbackLogic(ctx, s.svcCtx)
	return l.AddTaskFinishRollback(in)
}

func (s *SigninServer) UpdateTaskFinish(ctx context.Context, in *pb.UpdateTaskFinishReq) (*pb.UpdateTaskFinishResp, error) {
	l := logic.NewUpdateTaskFinishLogic(ctx, s.svcCtx)
	return l.UpdateTaskFinish(in)
}

func (s *SigninServer) DelTaskFinish(ctx context.Context, in *pb.DelTaskFinishReq) (*pb.DelTaskFinishResp, error) {
	l := logic.NewDelTaskFinishLogic(ctx, s.svcCtx)
	return l.DelTaskFinish(in)
}

func (s *SigninServer) GetTaskFinishById(ctx context.Context, in *pb.GetTaskFinishByIdReq) (*pb.GetTaskFinishByIdResp, error) {
	l := logic.NewGetTaskFinishByIdLogic(ctx, s.svcCtx)
	return l.GetTaskFinishById(in)
}

func (s *SigninServer) SearchTaskFinish(ctx context.Context, in *pb.SearchTaskFinishReq) (*pb.SearchTaskFinishResp, error) {
	l := logic.NewSearchTaskFinishLogic(ctx, s.svcCtx)
	return l.SearchTaskFinish(in)
}

// -----------------------remind-----------------------
func (s *SigninServer) ChangeSignRemind(ctx context.Context, in *pb.ChangeRemindReq) (*pb.ChangeRemindResp, error) {
	l := logic.NewChangeSignRemindLogic(ctx, s.svcCtx)
	return l.ChangeSignRemind(in)
}

func (s *SigninServer) GetRemindStatus(ctx context.Context, in *pb.GetRemindStatusReq) (*pb.GetRemindStatusResp, error) {
	l := logic.NewGetRemindStatusLogic(ctx, s.svcCtx)
	return l.GetRemindStatus(in)
}

func (s *SigninServer) SearchRemind(ctx context.Context, in *pb.SearchRemindReq) (*pb.SearchRemindResp, error) {
	l := logic.NewSearchRemindLogic(ctx, s.svcCtx)
	return l.SearchRemind(in)
}

func (s *SigninServer) SendRemind(ctx context.Context, in *pb.SendRemindReq) (*pb.SendRemindResp, error) {
	l := logic.NewSendRemindLogic(ctx, s.svcCtx)
	return l.SendRemind(in)
}
