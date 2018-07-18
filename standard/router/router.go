package router

import(
	"standard/controller" /*调用controller*/
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(){
	// Creates a default gin
	routes := gin.Default()
	routes.Delims("{[","]}")
	routes.Static("/static","static") /*代理静态目录*/
	routes.LoadHTMLGlob( "templates/*")/*加载html模版*/
	routes.GET( "/", controller.InitController) /*上传模板*/
	//路由分组
	apis := routes.Group( "/api")
	apis.GET( "/init", controller.GetModelController) /*初始化上传数据*/
	apis.POST( "/upload", controller.GetExcelController) /*导入excel*/
	apis.GET( "/introduce", controller.IntroduceModelController) /*上传标准化模板数据*/
	routes.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error": "404, page not exists!", })})
	routes.Run( ":3008") // listen and serve on 0.0.0.0:3008
}