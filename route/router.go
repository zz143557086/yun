package route

import (
	hdl "filestore-server/handler"

	"github.com/gin-gonic/gin"
)

// Router ：路由规则定义
func Router() *gin.Engine {
	// 创建一个默认的 Gin 引擎
	router := gin.Default()
	router.Use(hdl.Cors())
	// 静态资源处理
	router.Static("/static/", "./static")

	// 不需要验证的接口

	// 用户注册页面
	router.GET("/user/signup", hdl.SignupHandler)
	// 用户登录页面
	router.GET("/user/signin", hdl.SigninHandler)
	// 处理用户注册请求
	router.POST("/user/signup", hdl.DoSignupHandler)
	// 处理用户登录请求
	router.POST("/user/signin", hdl.DoSignInHandler)
	// 检查用户是否已存在
	router.GET("/user/exists", hdl.UserExistsHandler)

	// 添加认证中间件以确保需要身份验证的接口受到保护

	// 文件上传页面
	router.GET("/file/upload", hdl.UploadHandler)
	// 处理文件上传请求
	router.POST("/file/upload", hdl.DoUploadHandler)
	// 文件上传成功页面
	router.GET("/file/upload/suc", hdl.UploadSucHandler)
	// 获取文件元信息
	router.GET("/file/meta", hdl.GetFileMetaHandler)
	// 处理文件查询请求
	router.POST("/file/query", hdl.FileQueryHandler)
	// 文件下载页面
	router.GET("/file/download", hdl.DownloadHandler)
	// 处理文件下载请求
	router.POST("/file/download", hdl.DownloadHandler)
	// 处理文件元信息更新请求
	router.POST("/file/update", hdl.FileMetaUpdateHandler)
	// 处理文件删除请求
	router.POST("/file/delete", hdl.FileDeleteHandler)
	// 处理获取文件下载链接请求
	router.POST("/file/downloadurl", hdl.DownloadURLHandler)

	// 秒传接口
	router.POST("/file/fastupload", hdl.TryFastUploadHandler)

	// 分块上传接口

	//初始化分块信息
	router.POST("/file/mpupload/init", hdl.InitialMultipartUploadHandler)
	//上传分块
	router.POST("/file/mpupload/uppart", hdl.UploadPartHandler)
	//通知分块上传完成
	router.POST("/file/mpupload/complete", hdl.CompleteUploadHandler)

	// 用户相关接口

	// 处理获取用户信息请求
	router.POST("/user/info", hdl.UserInfoHandler)

	return router
}
