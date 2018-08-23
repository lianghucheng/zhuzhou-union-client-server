package utils

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pborman/uuid"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"golang.org/x/net/context"
	"path/filepath"
)

func UploadFile(fname string, data []byte) (url string, err error) {
	fileName := fname
	fileNameUuid := uuid.New()
	fileNameExt := filepath.Ext(fileName)
	filename := fileNameUuid + fileNameExt

	accessKey := beego.AppConfig.String("accessKey")
	secretKey := beego.AppConfig.String("secretKey")

	//pcvijeg6q.bkt.clouddn.com
	//localFile := file
	bucket := beego.AppConfig.String("qiniuBucket")
	key := filename
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"test": "test",
		},
	}

	dataLen := int64(len(data))

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	url = beego.AppConfig.String("qiniuUrl") + filename
	return url, nil
}
