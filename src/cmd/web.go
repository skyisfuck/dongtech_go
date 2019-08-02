package cmd

import (
	"dongtech_go/config"
	"dongtech_go/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
)

func startWeb(config *config.Config) {
	router := gin.Default()

	//打印panic error
	router.Use(serveRecover)

	//html
	router.LoadHTMLGlob("../templates/*")

	router.GET("/version", handler.Version)
	router.GET("/config", handler.PrintConfig)
	router.GET("/getUser/:id", handler.GetUser)
	router.GET("/index", handler.Index)
	router.GET("/uuid", handler.UUID)

	//jwt token
	router.GET("/login/:username/:password", handler.Login)
	router.GET("/verify/:token", handler.Verify)
	router.GET("/refresh/:token", handler.Refresh)
	router.GET("/sayHello/:token", handler.SayHello)

	router.GET("/grpc", Grpc(fmt.Sprintf(":%s", config.Grpc.Addr)))
	err := router.Run(fmt.Sprintf(":%s", config.Web.Addr))
	if err != nil {
		logrus.WithError(err).Println("web start failed")
		log.Fatalln(err)
		panic(err)
	}
}
