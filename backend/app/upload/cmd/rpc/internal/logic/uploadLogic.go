package logic

import (
	"bytes"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/upload/cmd/rpc/internal/svc"
	"looklook/app/upload/cmd/rpc/pb"
	"looklook/app/upload/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *pb.FileUploadReq) (*pb.FileUploadResp, error) {
	uploadFile := new(model.UploadFile)
	_ = copier.Copy(uploadFile, in)
	bucket, err := l.svcCtx.OssClient.Bucket(l.svcCtx.Config.Oss.BucketName)
	if err != nil {
		//todo 优化这里不应报数据库错误 应该提示oss错误
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "upload file with oss client err: %+v , err: %v", in, err)
	}
	// 创建一个 io.Reader
	reader := bytes.NewReader(in.FileData)
	fileName := in.FileName + in.Ext
	err = bucket.PutObject(fileName, reader)
	if err != nil {
		//todo 优化报错
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "put object with oss client err: %+v , err: %v", in, err)
	}
	url := l.svcCtx.Config.Oss.FilePathPrefix + fileName
	uploadFile.Url = url
	_, err = l.svcCtx.FileUploadModel.Insert(l.ctx, uploadFile)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "save file info with err: %+v , err: %v", in, err)
	}
	return &pb.FileUploadResp{
		Url: url,
	}, nil
}
