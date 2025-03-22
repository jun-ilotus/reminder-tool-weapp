package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"looklook/app/upload/cmd/api/internal/config"
	"looklook/app/upload/cmd/rpc/upload"
)

type ServiceContext struct {
	Config        config.Config
	FileUploadRpc upload.Upload
}

func NewServiceContext(c config.Config) *ServiceContext {
	dialOption := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024 * 1024 * 1024))
	opt := zrpc.WithDialOption(dialOption)
	return &ServiceContext{
		Config:        c,
		FileUploadRpc: upload.NewUpload(zrpc.MustNewClient(c.FileUploadRpcConf, opt)),
	}
}
