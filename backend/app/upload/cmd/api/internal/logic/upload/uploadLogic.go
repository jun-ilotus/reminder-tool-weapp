package upload

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/upload/cmd/api/internal/svc"
	"looklook/app/upload/cmd/api/internal/types"
	"looklook/app/upload/cmd/rpc/upload"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UserUploadReq) (resp *types.UserUploadRes, err error) {
	id := ctxdata.GetUidFromCtx(l.ctx)
	uploadResp, err := l.svcCtx.FileUploadRpc.Upload(l.ctx, &upload.FileUploadReq{
		UserId:   id,
		FileName: req.FileName,
		Size:     req.Size,
		Ext:      req.Ext,
		FileData: req.FileData,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Upload file fail"), "upload file fail with rpc err: %+v, err: %v", req, err)
	}
	resp = new(types.UserUploadRes)
	resp.Url = uploadResp.Url
	return
}
