package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/upload/cmd/rpc/internal/config"
	"looklook/app/upload/model"
)

type ServiceContext struct {
	Config          config.Config
	RedisClient     *redis.Redis
	FileUploadModel model.UploadFileModel
	OssClient       *oss.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	client, _ := oss.New(c.Oss.Endpoint, c.Oss.AccessKeyId, c.Oss.AccessKeySecret)
	_ = client.SetBucketACL(c.Oss.BucketName, oss.ACLPublicRead)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		FileUploadModel: model.NewUploadFileModel(sqlConn, c.Cache),
		OssClient:       client,
	}
}
