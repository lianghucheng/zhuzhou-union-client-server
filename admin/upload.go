package admin

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"zhuzhou-union-client-server/utils"
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
	file, err := fileHeader.Open()
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.returnFail()
	}

	url:=utils.UploadFile(fileHeader,data)

	c.returnSuccess(url)
	return
}

func (c *FileUploadController)returnSuccess(url string){
	result := make(map[string]interface{})
	result["error"] = 0
	result["message"] = "上传成功"
	result["url"] = url
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}

func (c *FileUploadController)returnFail(){
	result := make(map[string]interface{})
	result["error"] = 1
	result["message"] = "上传失败"
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}