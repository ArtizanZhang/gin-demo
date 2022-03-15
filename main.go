package main

import (
	"fmt"
	"github.com/ArtizanZhang/gin-demo/pkg/setting"
	"github.com/ArtizanZhang/gin-demo/routers"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"syscall"
)

func main() {
	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})

	//router := routers.InitRouter()
	//gin.SetMode(setting.RunMode)
	//
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//log.Fatal(s.ListenAndServe())

	gin.SetMode(setting.RunMode)
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf("localhost:%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(endPoint string) {
		log.Printf("Actual pid is %d %s ", syscall.Getpid(), endPoint)
	}

	log.Fatal(server.ListenAndServe())

}
