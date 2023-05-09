package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/XieWeiXie/httpbingo/controller/api"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.New()
	s.Use(gin.Recovery(), gin.Logger(), gin.ErrorLogger())

	v1 := s.Group("http-bin/v1")
	controller := api.HttpBinUS{}

	heartBeatGroup := v1.Group("heart")
	heartBeatGroup.Any("beat", func(context *gin.Context) {
		context.JSONP(http.StatusOK, gin.H{
			"code":    0,
			"message": "",
			"data":    fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05")),
		})
	})

	httpMethodGroup := v1.Group("httpMethods")
	httpMethodGroup.GET("get", controller.HttpMethodGet)
	httpMethodGroup.GET("put", controller.HttpMethodPut)
	httpMethodGroup.GET("delete", controller.HttpMethodDelete)
	httpMethodGroup.GET("post", controller.HttpMethodPost)

	statusCodeGroup := v1.Group("statusCodes")
	statusCodeGroup.GET("code", controller.StatusCode)

	headersGroup := v1.Group("headers")
	headersGroup.GET("header", controller.Headers)

	ipGroup := v1.Group("ips")
	ipGroup.GET("ip", controller.IP)

	userAgentGroup := v1.Group("userAgents")
	userAgentGroup.GET("userAgent", controller.UserAgent)

	_ = s.Run(":9091")

}
