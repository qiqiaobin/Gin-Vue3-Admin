package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tadmin/conf"
	docs "tadmin/docs"
	"tadmin/model/orm"
	"tadmin/pkg/cache"
	"tadmin/pkg/logger"
	"tadmin/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 系统模块
// @version 1.0.0
// @description 接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	os.Mkdir("logs", 0755)
	//runtime.GOMAXPROCS(runtime.NumCPU())
	// 初始化config.yml配置文件映射信息
	conf.InitConfig()
	// 日志初始化
	logger.InitLogger()
	//数据库初始化
	if db, err := orm.InitGorm(); err != nil {
		fmt.Println("初始化数据库失败:", err)
	} else {
		orm.DB = db
	}

	cache.InitRedis()
	//初始化IP归属地查询服务
	//ip2region.InitIp2region()

	//初始化路由
	var routers = router.InitRouters()

	// 初始化swagger文档
	InitSwagger(routers)

	//启动服务（优雅关机）
	addr := fmt.Sprintf(":%d", conf.Config.System.Port)
	router := router.InitRouters()
	srv := http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("成功启动服务，地址 :%d...\n", conf.Config.System.Port)

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("server run failed, err: %v", err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个10秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	fmt.Println("开始关闭服务...")
	// 创建一个10秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 10秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过10秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("服务被强制关闭", err)
	}

	fmt.Println("服务成功退出")
}

// InitSwagger 初始化swagger文档
func InitSwagger(routers *gin.Engine) {
	//swagger文档
	docs.SwaggerInfo.Title = "gin-vue3-admin"
	docs.SwaggerInfo.Description = "Swagger Admin API"
	docs.SwaggerInfo.Version = "1.0"
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
