// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6
// Source: recode.proto

package recode

import (
	"context"

	"looklook/app/recode/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddItemsReq       = pb.AddItemsReq
	AddItemsResp      = pb.AddItemsResp
	AddRecodeReq      = pb.AddRecodeReq
	AddRecodeResp     = pb.AddRecodeResp
	DelItemsReq       = pb.DelItemsReq
	DelItemsResp      = pb.DelItemsResp
	DelRecodeReq      = pb.DelRecodeReq
	DelRecodeResp     = pb.DelRecodeResp
	GetItemsByIdReq   = pb.GetItemsByIdReq
	GetItemsByIdResp  = pb.GetItemsByIdResp
	GetRecodeByIdReq  = pb.GetRecodeByIdReq
	GetRecodeByIdResp = pb.GetRecodeByIdResp
	Items             = pb.Items
	Recode            = pb.Recode
	SearchItemsReq    = pb.SearchItemsReq
	SearchItemsResp   = pb.SearchItemsResp
	SearchRecodeReq   = pb.SearchRecodeReq
	SearchRecodeResp  = pb.SearchRecodeResp
	UpdateItemsReq    = pb.UpdateItemsReq
	UpdateItemsResp   = pb.UpdateItemsResp
	UpdateRecodeReq   = pb.UpdateRecodeReq
	UpdateRecodeResp  = pb.UpdateRecodeResp

	RecodeZrpcClient interface {
		// -----------------------items-----------------------
		AddItems(ctx context.Context, in *AddItemsReq, opts ...grpc.CallOption) (*AddItemsResp, error)
		UpdateItems(ctx context.Context, in *UpdateItemsReq, opts ...grpc.CallOption) (*UpdateItemsResp, error)
		DelItems(ctx context.Context, in *DelItemsReq, opts ...grpc.CallOption) (*DelItemsResp, error)
		GetItemsById(ctx context.Context, in *GetItemsByIdReq, opts ...grpc.CallOption) (*GetItemsByIdResp, error)
		SearchItems(ctx context.Context, in *SearchItemsReq, opts ...grpc.CallOption) (*SearchItemsResp, error)
		// -----------------------recode-----------------------
		AddRecode(ctx context.Context, in *AddRecodeReq, opts ...grpc.CallOption) (*AddRecodeResp, error)
		UpdateRecode(ctx context.Context, in *UpdateRecodeReq, opts ...grpc.CallOption) (*UpdateRecodeResp, error)
		DelRecode(ctx context.Context, in *DelRecodeReq, opts ...grpc.CallOption) (*DelRecodeResp, error)
		GetRecodeById(ctx context.Context, in *GetRecodeByIdReq, opts ...grpc.CallOption) (*GetRecodeByIdResp, error)
		SearchRecode(ctx context.Context, in *SearchRecodeReq, opts ...grpc.CallOption) (*SearchRecodeResp, error)
	}

	defaultRecodeZrpcClient struct {
		cli zrpc.Client
	}
)

func NewRecodeZrpcClient(cli zrpc.Client) RecodeZrpcClient {
	return &defaultRecodeZrpcClient{
		cli: cli,
	}
}

// -----------------------items-----------------------
func (m *defaultRecodeZrpcClient) AddItems(ctx context.Context, in *AddItemsReq, opts ...grpc.CallOption) (*AddItemsResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.AddItems(ctx, in, opts...)
}

func (m *defaultRecodeZrpcClient) UpdateItems(ctx context.Context, in *UpdateItemsReq, opts ...grpc.CallOption) (*UpdateItemsResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.UpdateItems(ctx, in, opts...)
}

func (m *defaultRecodeZrpcClient) DelItems(ctx context.Context, in *DelItemsReq, opts ...grpc.CallOption) (*DelItemsResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.DelItems(ctx, in, opts...)
}

func (m *defaultRecodeZrpcClient) GetItemsById(ctx context.Context, in *GetItemsByIdReq, opts ...grpc.CallOption) (*GetItemsByIdResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.GetItemsById(ctx, in, opts...)
}

func (m *defaultRecodeZrpcClient) SearchItems(ctx context.Context, in *SearchItemsReq, opts ...grpc.CallOption) (*SearchItemsResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.SearchItems(ctx, in, opts...)
}

// -----------------------recode-----------------------
func (m *defaultRecodeZrpcClient) AddRecode(ctx context.Context, in *AddRecodeReq, opts ...grpc.CallOption) (*AddRecodeResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.AddRecode(ctx, in, opts...)
}

func (m *defaultRecodeZrpcClient) UpdateRecode(ctx context.Context, in *UpdateRecodeReq, opts ...grpc.CallOption) (*UpdateRecodeResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.UpdateRecode(ctx, in, opts...)
}

func (m *defaultRecodeZrpcClient) DelRecode(ctx context.Context, in *DelRecodeReq, opts ...grpc.CallOption) (*DelRecodeResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.DelRecode(ctx, in, opts...)
}

func (m *defaultRecodeZrpcClient) GetRecodeById(ctx context.Context, in *GetRecodeByIdReq, opts ...grpc.CallOption) (*GetRecodeByIdResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.GetRecodeById(ctx, in, opts...)
}

func (m *defaultRecodeZrpcClient) SearchRecode(ctx context.Context, in *SearchRecodeReq, opts ...grpc.CallOption) (*SearchRecodeResp, error) {
	client := pb.NewRecodeClient(m.cli.Conn())
	return client.SearchRecode(ctx, in, opts...)
}
