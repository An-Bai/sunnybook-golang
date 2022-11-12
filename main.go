package main

import (
	"fmt"
	"net/http"
	"sunnybook-golang/pkg/setting"
	"sunnybook-golang/routers"
)

func main() {
	// 调用routers的路由配置
	router := routers.InitRouter()

	// 不使用默认服务器router.Run()，使用我们自己配置的服务器参数启动http服务
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
