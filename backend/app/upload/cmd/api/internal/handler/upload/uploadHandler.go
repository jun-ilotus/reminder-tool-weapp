package upload

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"io/ioutil"
	"looklook/app/upload/cmd/api/internal/logic/upload"
	"looklook/app/upload/cmd/api/internal/svc"
	"looklook/app/upload/cmd/api/internal/types"
	"looklook/common/result"
	"net/http"
	"path/filepath"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		// 读取文件内容到字节切片
		buffer := make([]byte, 512)
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			result.ParamErrorResult(r, w, errors.New("Error reading file"))
			return
		}
		fileType := http.DetectContentType(buffer[:n])
		if !isImage(fileType) {
			result.ParamErrorResult(r, w, errors.New("File is not an image"))
			return
		}
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			result.ParamErrorResult(r, w, errors.New("Error reading file"))
			return
		} // 将文件指针重置到文件开头
		//todo 添加文件大小限制
		//当文件大于10M
		//if header.Size >= (10 << 20) {
		//	result.ParamErrorResult(r, w, errors.New("上传文件大小超出限制，请上传小于10M的文件"))
		//	return
		//}
		//todo 压缩文件

		//获取文件类型
		ext := filepath.Ext(header.Filename)
		//获取文件内容
		all, err := ioutil.ReadAll(file)
		if err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		//填充参数
		//req.FileName = header.Filename
		//使用uuid生成文件名
		req.FileName = uuid.NewV4().String()
		req.Size = header.Size
		req.Ext = ext
		req.FileData = all

		l := upload.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		if err != nil {
			result.ParamErrorResult(r, w, err)
		} else {
			result.HttpResult(r, w, resp, err)
		}
	}
}

// isImage 检查 MIME 类型是否是图片类型
func isImage(fileType string) bool {
	imageTypes := []string{
		"image/jpeg",
		"image/png",
		"image/gif",
		"image/webp",
		"image/bmp",
		"image/tiff",
	}
	for _, t := range imageTypes {
		if fileType == t {
			return true
		}
	}
	return false
}
