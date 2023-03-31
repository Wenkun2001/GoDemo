package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 自定义Go中间件 拦截器
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 通过自动逸的中间件，设置值，在后续处理中只要调用了这个中间件，就可以拿到这里的参数关于
		context.Set("userSession", "userid")
		context.Next() // 放行

		//context.Abort()	//阻止
	}
}

func main() {
	// 创建一个服务
	ginServer := gin.Default()
	//ginServer.Use(favicon.New("./favicon.ico"))

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")

	// 访问地址，处理我们的请求 Request Response
	// Gin RestFul
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,world"})
	})
	ginServer.POST("/user", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "post, user"})
	})

	// 响应一个页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		//context.JSON()   json数据
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "后台响应传递数据",
		})
	})

	// 接收前端传递过来的参数
	ginServer.GET("/user/info", myHandler(), func(context *gin.Context) {

		// 取出中间件中的值
		userSession := context.MustGet("userSession").(string)
		log.Println(userSession)

		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	ginServer.GET("user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 前端给后端传递json
	ginServer.POST("/json", func(context *gin.Context) {
		// request.body
		// []byte
		data, _ := context.GetRawData()

		var m map[string]interface{}
		// 包装为json数据[]byte
		// 序列化
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})

	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	// 路由
	ginServer.GET("/test", func(context *gin.Context) {
		// 重定向 301
		context.Redirect(http.StatusMovedPermanently, "https://baidu.com")
	})

	// 404 Not Found
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	//// 路由组
	//userGroup := ginServer.Group("/user") {
	//	userGroup.GET("/add")
	//	userGroup.GET("/login")
	//	userGroup.GET("/logout")
	//}

	// 服务器端口
	ginServer.Run(":8082")

}
