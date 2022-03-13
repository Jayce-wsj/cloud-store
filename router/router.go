package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"workspace/cloud-pan/config"
	"workspace/cloud-pan/handler"
	"workspace/cloud-pan/util/logs"
)

func RouterRun(port int, args ...string) error {
	gin.SetMode(gin.ReleaseMode)

	//路由
	router := gin.Default()
	router.Use(Middelware())
	//加载模版文件
	router.LoadHTMLGlob("./static/view/*")
	//加载静态资源
	router.Static("/static", "./static")

	//定义GET方法
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			//改变量在"index.tmpl"中是有定义的，可以通过相同变量名拿到对应的数据哟~
			"title": "Hello World !",
		})
	})

	//主页面
	router.GET("/main/index", func(context *gin.Context) {
		//查看用户是否登陆，没登录跳转到登陆页，已登录显示主页
		loginUser, err := context.Cookie("username")
		logs.Debugf("user:%v,err:%v", loginUser, err)
		if loginUser == "" {
			context.Redirect(302, "/user/signin")
		} else {
			context.HTML(http.StatusOK, "home.html", gin.H{})
		}
	})

	//用户相关接口
	user := router.Group("/user")
	{
		//用户登录
		user.GET("/signin", func(context *gin.Context) {
			context.HTML(http.StatusOK, "signin.html", gin.H{})
		})
		user.POST("/dosignin", handler.SigninHandler)

		//用户注册
		user.GET("/signup", func(context *gin.Context) {
			context.HTML(http.StatusOK, "signup.html", gin.H{})
		})
		user.POST("/signup", handler.SignupHandler)
		user.POST("/info", handler.InfoHandler)
		user.GET("/logout", func(context *gin.Context) {
			context.SetCookie("username", "", -1, "/", context.Request.Host, false, true)
			context.Redirect(302, "/user/signin")
		})
	}

	//文件相关接口
	file := router.Group("/file")
	{
		//文件列表
		file.POST("/query", handler.FileQueryHandler)
		//文件上传
		file.GET("/upload", func(context *gin.Context) {
			context.HTML(http.StatusOK, "upload.html", gin.H{})
		})
		//文件列表
		file.POST("/oss", handler.FileOssHandler)
		//下载文件
		file.GET("/download", handler.DownloadURLHandler)
	}

	return router.Run(fmt.Sprintf(":%d", port))
}

//路由中间件
func Middelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin,Content-Length,Content-Type,X-Token,Access-Token,Token,Sign,Tm,UUID,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requeste")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
		//nginx 重复了, 上生产注释掉
		if config.RunMode == "debug" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
