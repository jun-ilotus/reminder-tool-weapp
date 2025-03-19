// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: reminder.proto

package server

import (
	"context"

	"looklook/app/reminder/cmd/rpc/internal/logic"
	"looklook/app/reminder/cmd/rpc/internal/svc"
	"looklook/app/reminder/cmd/rpc/pb"
)

type ReminderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedReminderServer
}

func NewReminderServer(svcCtx *svc.ServiceContext) *ReminderServer {
	return &ReminderServer{
		svcCtx: svcCtx,
	}
}

// -----------------------reminder-----------------------
func (s *ReminderServer) AddReminder(ctx context.Context, in *pb.AddReminderReq) (*pb.AddReminderResp, error) {
	l := logic.NewAddReminderLogic(ctx, s.svcCtx)
	return l.AddReminder(in)
}

func (s *ReminderServer) UpdateReminder(ctx context.Context, in *pb.UpdateReminderReq) (*pb.UpdateReminderResp, error) {
	l := logic.NewUpdateReminderLogic(ctx, s.svcCtx)
	return l.UpdateReminder(in)
}

func (s *ReminderServer) DelReminder(ctx context.Context, in *pb.DelReminderReq) (*pb.DelReminderResp, error) {
	l := logic.NewDelReminderLogic(ctx, s.svcCtx)
	return l.DelReminder(in)
}

func (s *ReminderServer) GetReminderById(ctx context.Context, in *pb.GetReminderByIdReq) (*pb.GetReminderByIdResp, error) {
	l := logic.NewGetReminderByIdLogic(ctx, s.svcCtx)
	return l.GetReminderById(in)
}

func (s *ReminderServer) SearchReminder(ctx context.Context, in *pb.SearchReminderReq) (*pb.SearchReminderResp, error) {
	l := logic.NewSearchReminderLogic(ctx, s.svcCtx)
	return l.SearchReminder(in)
}
