package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"os"
	"path"
	"strings"
)

type RetJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Url  string      `json:"url"`
}

type PublicController struct {
	beego.Controller
}

func (this *PublicController) sendJson(code int, msg string, data interface{}, url string) {
	ret := RetJson{1, "", "", ""}
	ret.Code = code
	ret.Msg = msg
	ret.Data = data
	ret.Url = url

	this.Data["json"] = &ret
	this.ServeJSON()
	this.StopRun()
}

// 初始化redis
func getRedis() (redis cache.Cache) {
	host := beego.AppConfig.String("RedisHost")
	port := beego.AppConfig.String("RedisPort")

	redis, err := cache.NewCache("redis", `{"key":"novel","conn":"`+host+`:`+port+`","dbNum":"0"}`)

	if err != nil {
		beego.Info("init redis error")
		panic("init redis error")
		return
	}
	return
}

func Struct2Json(data interface{}) (jsonData string) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	jsonData = string(jsonBytes)
	return
}

// 保存图片
func (this *PublicController) saveImage(formName, customPath, fileName string) (realPath string, err error) {
	// 配置中定义的目录
	basePath := beego.AppConfig.String("ImagePath")
	// 实际目录
	filePath := basePath + customPath

	// 判断没有该目录，没有则创建
	err = createFile(filePath)
	if err != nil {
		return
	}

	// 从上传中获取相应文件信息
	f, h, _ := this.GetFile(formName)

	tempName := h.Filename
	arr := strings.Split(tempName, ":")
	formFileaName := arr[len(arr)-1]
	fileArr := strings.Split(formFileaName, ".")
	suffix := fileArr[len(fileArr)-1]

	realPath = path.Join(filePath, fileName+"."+suffix)

	f.Close()
	err = this.SaveToFile(formName, realPath)
	return
}

// 判断目录是否存在
func isExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 递归创建文件
func createFile(filePath string) error {

	if !isExists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

// 调试
func showDebug(msg interface{}) {
	beego.Debug(msg)
}
