// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"net/http"
	"path/filepath"
	"zgw/ks/flash_sale/user/pkg"

	"zgw/ks/flash_sale/api/internal/svc"
	"zgw/ks/flash_sale/api/internal/types"

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

func (l *UploadLogic) Upload(req *http.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	file, m, err := req.FormFile("file")
	if err != nil {
		return &types.Response{
			Data: nil,
			Msg:  err.Error(),
			Code: 400,
		}, nil
	}
	if m.Size > 1024*1024*1024*8 {
		return &types.Response{
			Data: nil,
			Msg:  "视频文件大小不得超过8GB",
			Code: 400,
		}, nil
	}
	ext := filepath.Ext(m.Filename)
	if ext != ".mp4" {
		return &types.Response{
			Data: nil,
			Msg:  "允许上传mp4格式的视频文件。",
			Code: 400,
		}, nil
	}
	upload, err := pkg.Upload(file, m)
	if err != nil {
		return &types.Response{
			Data: nil,
			Msg:  err.Error(),
			Code: 400,
		}, nil
	}
	return &types.Response{
		Data: upload,
		Msg:  "success",
		Code: 200,
	}, nil
}
