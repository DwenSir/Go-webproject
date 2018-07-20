package controller

import (
	"github.com/gin-gonic/gin"
	"standard/controller/getModel"
	"standard/controller/introduceModel"
	"standard/controller/getExcel"
	"net/http"
	"os"
	"fmt"
	"io"
	"standard/pkg/pathJoin"
	"encoding/json"
)
//TODO:控制器 确定执行逻辑和返回对应的数据

func InitController(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", gin.H{ "title": "standardizationTools",})
}

//TODO:return初始化数据 模板信息和工具版本信息
func GetModelController(c *gin.Context){
	ver := getModel.VersionCfg()
	infos := getModel.ModelInfo()["model_info"]
	if ver == "None" || infos == nil{
//		返回错误信息
		c.JSON(http.StatusOK, gin.H{
			"status": 404,
			"version": "error",
			"model_info":"error",
		})
	}else {
//		常规返回
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"version": ver,/*版本号*/
			"model_info": infos, /*模板信息 返回的是map*/
		})
	}
}

//TODO:接收前端传过来的excel文件
func GetExcelController(c *gin.Context){
	file, header, err := c.Request.FormFile("file_name")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	filename := header.Filename
//	set excel save path
	pwd,_ := os.Getwd()
	driversPath := pathJoin.Join(pwd,"drivers")
	copyDoc := pathJoin.Join(driversPath,filename)
	out, err := os.Create(copyDoc)/*存放excel*/
	if err != nil {fmt.Println(err)}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {fmt.Println(err)}
	b := getExcel.GetXlsxInfo(filename)
	c.JSON(http.StatusCreated, b)
}

//TODO:return标准化模板信息
func IntroduceModelController(c *gin.Context){
	name := c.Query("file_name")
	data := introduceModel.GetModelInfo(name)
	c.JSON(http.StatusOK,data)
}


type s_data struct {
	Url       string   `json:"url" binding:"required"`
	Imgs      []string `json:"imgs"`
	Title     string   `json:"title" binding:"required"`
}

//TODO:快速绑定
func BindingInfo(c *gin.Context)  {
	var p s_data
	if c.BindJSON(&p) == nil {
		project := c.Param("project")
		taskid := c.Param("taskid")
		c.JSON(http.StatusOK, gin.H{
			"err_no":    "2001",
			"message":   "result save successful!",
			"project":   project,
			"taskid":    taskid,
			"title":     p.Title,
			"data":      func() string {
				json_str, _ := json.Marshal(p.Imgs)
				return string(json_str)
			}(),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"err_no":     "4001",
			"message":    "request's params wrong!",
		})
	}
//	binding.Binding()
}

//TODO:写回excel
func ToExcel(c *gin.Context)  {

}