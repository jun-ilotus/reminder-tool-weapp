package logic

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

func TestUpload(T *testing.T) {
	/// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("yourEndpoint", "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// yourBucketName填写存储空间名称。
	// 设置存储空间的读写权限为公共读。
	err = client.SetBucketACL("yourBucketName", oss.ACLPublicRead)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
