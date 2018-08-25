package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/pborman/uuid"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"path/filepath"
	"io/ioutil"
	"zhuzhou-union-client-server/utils"
)

type Ueditor struct {
	Common
}

type config_Json struct {
	ImageActionName     string   `json:"imageActionName"`
	ImageFieldName      string   `json:"imageFieldName"`
	ImageMaxSize        int      `json:"imageMaxSize"`
	ImageAllowFiles     []string `json:"imageAllowFiles"`
	ImageCompressEnable bool     `json:"imageCompressEnable"`
	ImageCompressBorder int      `json:"imageCompressBorder"`
	ImageInsertAlign    string   `json:"imageInsertAlign"`
	ImageUrlPrefix      string   `json:"imageUrlPrefix"`
	ImagePathFormat     string   `json:"imagePathFormat"`

	ScrawlActionName  string `json:"scrawlActionName"`
	ScrawlFieldName   string `json:"scrawlFieldName"`
	ScrawlPathFormat  string `json:"scrawlPathFormat"`
	ScrawlMaxSize     int    `json:" scrawlMaxSize"`
	ScrawlUrlPrefix   string `json:"scrawlUrlPrefix"`
	ScrawlInsertAlign string `json:"scrawlInsertAlign"`

	SnapscreenActionName  string `json:"snapscreenActionName"`
	SnapscreenPathFormat  string `json:"snapscreenPathFormat"`
	SnapscreenUrlPrefix   string `json:"snapscreenUrlPrefix"`
	SnapscreenInsertAlign string `json:"snapscreenInsertAlign"`

	VideoActionName string   `json:"videoActionName"`
	VideoFieldName  string   `json:"videoFieldName"`
	VideoPathFormat string   `json:"videoPathFormat"`
	VideoUrlPrefix  string   `json:"videoUrlPrefix"`
	VideoMaxSize    int      `json:"videoMaxSize"`
	VideoAllowFiles []string `json:"videoAllowFiles"`

	ImageManagerActionName  string   `json:"imageManagerActionName"`
	ImageManagerListPath    string   `json:"imageManagerListPath"`
	ImageManagerListSize    int      `json:"imageManagerListSize"`
	ImageManagerUrlPrefix   string   `json:"imageManagerUrlPrefix"`
	ImageManagerInsertAlign string   `json:"imageManagerInsertAlign"`
	ImageManagerAllowFiles  []string `json:"imageManagerAllowFiles"`

	FileActionName string   `json:"fileActionName"`
	FileFieldName  string   `json:"fileFieldName"`
	FileMaxSize    int      `json:"fileMaxSize"`
	FileAllowFiles []string `json:"fileAllowFiles"`
	FileUrlPrefix  string   `json:"fileUrlPrefix"`
	FilePathFormat string   `json:"filePathFormat"`

	FileManagerActionName string   `json:"fileManagerActionName"`
	FileManagerListPath   string   `json:"fileManagerListPath"`
	FileManagerUrlPrefix  string   `json:"fileManagerUrlPrefix"`
	FileManagerListSize   int      `json:"fileManagerListSize"`
	FileManagerAllowFiles []string `json:"fileManagerAllowFiles"`
}

type upload_res struct {
	Status   string `json:"state"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	Original string `json:"original"`
}

type list_st struct {
	Url string `json:"url"`
}

type list_res struct {
	Status string    `json:"state"`
	List   []list_st `json:"list"`
	Start  int       `json:"start"`
	Total  int       `json:"total"`
}

var configJson config_Json

func Init_Ueditor(config_path string) {

	file, err := os.Open(config_path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(file)

	json.Unmarshal(buf.Bytes(), &configJson)

	isRemoteUpload, _ := beego.AppConfig.Bool("isRemoteUpload")

	if !isRemoteUpload {
		os.MkdirAll(configJson.ImagePathFormat, 0777)
		os.MkdirAll(configJson.FilePathFormat, 0777)
		os.MkdirAll(configJson.VideoPathFormat, 0777)
		os.MkdirAll(configJson.SnapscreenPathFormat, 0777)
		os.MkdirAll(configJson.ScrawlPathFormat, 0777)
	}
}

func (this *Ueditor) U_Controller() {
	ueditorPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	Init_Ueditor(ueditorPath + "/conf/ueditor_config.json")
	var jsondata interface{}
	defer func() {
		this.Data["json"] = jsondata
		this.ServeJSON()
	}()
	action := this.GetString("action")
	beego.Debug(action)
	switch action {
	case "config":
		jsondata = configJson
	case configJson.ImageActionName:
		jsondata = upload_uuid_remote(this.Controller, configJson.ImageFieldName, configJson.ImagePathFormat)
	case configJson.ScrawlActionName:
		jsondata = upload_uuid_remote(this.Controller, configJson.ScrawlFieldName, configJson.ScrawlPathFormat)
	case configJson.VideoActionName:
		jsondata = upload_time_remote(this.Controller, configJson.VideoFieldName, configJson.VideoPathFormat)
	case configJson.FileActionName:
		jsondata = upload_uuid_remote(this.Controller, configJson.FileFieldName, configJson.FilePathFormat)
	case configJson.ImageManagerActionName:
		jsondata = list_file(this.Controller, configJson.ImageManagerListPath)
	case configJson.FileManagerActionName:
		jsondata = list_file(this.Controller, configJson.FileManagerListPath)
	default:
		map_ := make(map[string]interface{})
		map_["state"] = "错误"
		jsondata = map_
		beego.Debug("error")
	}
}

func upload_uuid(this beego.Controller, FieldName, PathFormat string) interface{} {
	var jsondata upload_res
	File_in, File_h, err := this.GetFile(FieldName)

	defer File_in.Close()
	if err != nil {
		beego.Error(err)
		jsondata.Status = "上传失败"
		return jsondata
	}
	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(File_h.Filename)
	this.SaveToFile(FieldName, path.Join(PathFormat, filename))
	jsondata.Status = "SUCCESS"
	jsondata.Original = File_h.Filename
	jsondata.Title = File_h.Filename
	jsondata.Url = filename
	return jsondata
}

func upload_uuid_remote(this beego.Controller, FieldName, PathFormat string) interface{} {
	var jsondata upload_res
	File_in, File_h, err := this.GetFile(FieldName)
	defer File_in.Close()
	if err != nil {
		beego.Error(err)
		jsondata.Status = "上传失败"
		return jsondata
	}
	fileByte, _ := ioutil.ReadAll(File_in)
	url, _ := utils.UploadFile(File_h.Filename, fileByte)

	jsondata.Status = "SUCCESS"
	jsondata.Original = File_h.Filename
	jsondata.Title = File_h.Filename
	jsondata.Url = url
	return jsondata
}

func upload_time(this beego.Controller, FieldName, PathFormat string) interface{} {
	var jsondata upload_res
	File_in, File_h, err := this.GetFile(FieldName)
	File_in.Close()
	if err != nil {
		beego.Error(err)
		jsondata.Status = "上传失败"
		return jsondata
	}
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + path.Ext(File_h.Filename)
	this.SaveToFile(FieldName, path.Join(PathFormat, filename))
	jsondata.Status = "SUCCESS"
	jsondata.Original = File_h.Filename
	jsondata.Title = File_h.Filename
	jsondata.Url = filename
	return jsondata
}

func upload_time_remote(this beego.Controller, FieldName, PathFormat string) interface{} {
	var jsondata upload_res
	File_in, File_h, err := this.GetFile(FieldName)
	File_in.Close()
	if err != nil {
		beego.Error(err)
		jsondata.Status = "上传失败"
		return jsondata
	}
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + path.Ext(File_h.Filename)
	fileByte, _ := ioutil.ReadAll(File_in)
	utils.UploadFile(File_h.Filename, fileByte)
	jsondata.Status = "SUCCESS"
	jsondata.Original = File_h.Filename
	jsondata.Title = File_h.Filename
	jsondata.Url = filename
	return jsondata
}

func upload_name(this beego.Controller, FieldName, PathFormat string) interface{} {
	var jsondata upload_res
	File_in, File_h, err := this.GetFile(FieldName)
	File_in.Close()
	if err != nil {
		beego.Error(err)
		jsondata.Status = "上传失败"
		return jsondata
	}
	filename := File_h.Filename
	f, err := os.Open(path.Join(PathFormat, filename))
	f.Close()
	if err == nil {
		jsondata.Status = "该文件名已存在"
		return jsondata
	}
	this.SaveToFile(FieldName, path.Join(PathFormat, filename))
	jsondata.Status = "SUCCESS"
	jsondata.Original = File_h.Filename
	jsondata.Title = File_h.Filename
	jsondata.Url = filename
	return jsondata
}

func upload_name_remote(this beego.Controller, FieldName, PathFormat string) interface{} {
	var jsondata upload_res
	File_in, File_h, err := this.GetFile(FieldName)
	File_in.Close()
	if err != nil {
		beego.Error(err)
		jsondata.Status = "上传失败"
		return jsondata
	}
	filename := File_h.Filename
	f, err := os.Open(path.Join(PathFormat, filename))
	f.Close()
	if err == nil {
		jsondata.Status = "该文件名已存在"
		return jsondata
	}
	fileByte, _ := ioutil.ReadAll(File_in)
	utils.UploadFile(File_h.Filename, fileByte)
	this.SaveToFile(FieldName, path.Join(PathFormat, filename))
	jsondata.Status = "SUCCESS"
	jsondata.Original = File_h.Filename
	jsondata.Title = File_h.Filename
	jsondata.Url = filename
	return jsondata
}

func list_file(this beego.Controller, Path string) interface{} {
	start, _ := this.GetInt("start")
	size, _ := this.GetInt("start")
	var jsondata list_res
	file, err := os.Open(Path)
	defer file.Close()
	if err != nil {
		beego.Error(err)
		jsondata.Status = "未知错误"
	}
	stat, err := file.Stat()
	if err != nil {
		beego.Error(err)
		jsondata.Status = "未知错误"
	}
	if stat.IsDir() {
		dirs, err := file.Readdir(0)
		if err != nil {
			beego.Error(err)
			jsondata.Status = "未知错误"
		}
		jsondata.Total = len(dirs)
		if jsondata.Total > start {
			jsondata.Start = start
		} else {
			jsondata.Start = 0
		}
		for index, fileInfo := range dirs {
			if index < jsondata.Start {
				continue
			}
			var t list_st
			t.Url = fileInfo.Name()
			jsondata.List = append(jsondata.List, t)
			if index < jsondata.Start+size {
				break
			}
		}
	}
	jsondata.Status = "SUCCESS"
	return jsondata
}
