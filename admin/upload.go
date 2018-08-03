package admin

import (
	"github.com/astaxie/beego"
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"fmt"
	"golang.org/x/net/context"
	"bytes"
	"io/ioutil"
	"github.com/google/uuid"
	"path/filepath"
)

type FileUploadController struct {
	beego.Controller
}

//@router /image/kindeditor/upload [*]
func (c *FileUploadController) Upload() {
	_, fileHeader, err := c.GetFile("imgFile")
	if err != nil {
		beego.Debug("get file error :", err)
		return
	}
	result := make(map[string]interface{})
	file, err := fileHeader.Open()
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		result["error"] = 1
		result["message"] = "上传失败"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	fileName := fileHeader.Filename
	fileNameUuid, _ := uuid.NewUUID()
	fileNameExt := filepath.Ext(fileName)
	filename := fileNameUuid.String() + fileNameExt

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
		return
	}
	url := beego.AppConfig.String("qiniuUrl") + filename
	result["error"] = 0
	result["message"] = "上传成功"
	result["url"] = url
	c.Data["json"] = result
	c.ServeJSON()
	return
}
