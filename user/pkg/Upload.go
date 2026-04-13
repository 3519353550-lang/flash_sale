package pkg

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"mime/multipart"
	"zgw/ks/flash_sale/user/configs"
)

func Upload(file multipart.File, m *multipart.FileHeader) (string, error) {
	yun := configs.Conf.QiNiuYun
	accessKey := yun.AccessKey
	secretKey := yun.SecretKey
	mac := credentials.NewCredentials(accessKey, secretKey)
	bucket := yun.Bucket
	key := m.Filename
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err := uploadManager.UploadReader(context.Background(), file, &uploader.ObjectOptions{
		BucketName: bucket,
		ObjectName: &key,
		CustomVars: map[string]string{
			"name": "github logo",
		},
		FileName: key,
	}, nil)
	fmt.Println(err)
	return yun.Url + key, err
}
