package admin

import (
	"github.com/astaxie/beego"
	"fmt"
)

type FileUploadController struct {
	beego.Controller
}

//@router /image/kindeditor/upload [*]
func (c *FileUploadController) Upload() {
	fmt.Println("upload success")
}
