package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/XieWeiXie/httpbingo/controller/api"
	"github.com/gin-gonic/gin"
	"github.com/larksuite/oapi-sdk-gin"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func main() {
	s := gin.New()
	s.Use(gin.Recovery(), gin.Logger(), gin.ErrorLogger())

	client := lark.NewClient("", "")

	handler := dispatcher.NewEventDispatcher("verificationToken", "eventEncryptKey").
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			fmt.Println(larkcore.Prettify(event))
			fmt.Println(event.RequestId())
			tenantKey := event.TenantKey()
			resp, err := client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
				ReceiveIdType(larkim.ReceiveIdTypeOpenId).
				Body(larkim.NewCreateMessageReqBodyBuilder().
					MsgType(larkim.MsgTypePost).
					ReceiveId("ou_c245b0a7dff2725cfa2fb104f8b48b9d").
					Content("text").
					Build()).
				Build(), larkcore.WithTenantKey(tenantKey))
			fmt.Println(resp, err)
			return nil
		}).
		OnP2MessageReadV1(func(ctx context.Context, event *larkim.P2MessageReadV1) error {
			fmt.Println(larkcore.Prettify(event))
			fmt.Println(event.RequestId())
			return nil
		})

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
	httpMethodGroup.PUT("put", controller.HttpMethodPut)
	httpMethodGroup.DELETE("delete", controller.HttpMethodDelete)
	httpMethodGroup.POST("post", controller.HttpMethodPost)

	statusCodeGroup := v1.Group("statusCodes")
	statusCodeGroup.GET("code", controller.StatusCode)

	headersGroup := v1.Group("headers")
	headersGroup.GET("header", controller.Headers)

	ipGroup := v1.Group("ips")
	ipGroup.GET("ip", controller.IP)

	userAgentGroup := v1.Group("userAgents")
	userAgentGroup.GET("userAgent", controller.UserAgent)

	webhook := v1.Group("feishu")
	webhook.POST("webhook/event", sdkginext.NewEventHandlerFunc(handler))

	_ = s.Run(":9091")

}
