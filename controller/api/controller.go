package api

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/XieWeiXie/httpbingo/api/v1"
	"github.com/XieWeiXie/httpbingo/services/api"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"net/http"
	"strconv"
)

var ma = jsonpb.Marshaler{
	EnumsAsInts:  true,
	OrigName:     true,
	EmitDefaults: true,
}

type HttpBinUS struct {
}

func (a HttpBinUS) HttpMethodGet(ctx *gin.Context) {
	var req v1.HttpMethodGetReq
	id, _ := strconv.Atoi(ctx.Query("id"))
	req.Id = int32(id)
	if req.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("请求参数错误"),
		})
		return
	}
	service := api.Service{}
	response, err := service.HttpMethodGet(context.TODO(), &req)
	if err != nil {
		st := status.Convert(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("内部错误: %v", st.Message()),
		})
		return
	}
	ret, _ := ma.MarshalToString(response)
	ctx.JSONP(http.StatusOK, json.RawMessage(ret))
}

func (a HttpBinUS) HttpMethodPut(ctx *gin.Context) {
	var req v1.HttpMethodPutReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("请求参数错误: %v", err.Error()),
		})
		return
	}
	service := api.Service{}
	response, err := service.HttpMethodPut(context.TODO(), &req)
	if err != nil {
		st := status.Convert(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("内部错误: %v", st.Message()),
		})
		return
	}
	ret, _ := ma.MarshalToString(response)
	ctx.JSON(http.StatusOK, json.RawMessage(ret))
}

func (a HttpBinUS) HttpMethodDelete(ctx *gin.Context) {
	var req v1.HttpMethodDeleteReq
	id, _ := strconv.Atoi(ctx.Query("id"))
	req.Id = int32(id)
	if id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("请求参数错误"),
		})
		return
	}
	service := api.Service{}
	response, err := service.HttpMethodDelete(context.TODO(), &req)
	if err != nil {
		st := status.Convert(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("内部错误: %v", st.Message()),
		})
		return
	}
	ret, _ := ma.MarshalToString(response)
	ctx.JSON(http.StatusOK, json.RawMessage(ret))
}

func (a HttpBinUS) HttpMethodPost(ctx *gin.Context) {
	var req v1.HttpMethodPostReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("请求参数错误: %v", err.Error()),
		})
		return
	}
	service := api.Service{}
	response, err := service.HttpMethodPost(context.TODO(), &req)
	if err != nil {
		st := status.Convert(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("内部错误: %v", st.Message()),
		})
		return
	}
	ret, _ := ma.MarshalToString(response)
	ctx.JSON(http.StatusOK, json.RawMessage(ret))
}

func (a HttpBinUS) StatusCode(ctx *gin.Context) {
	var req v1.StatusCodeReq
	code, _ := strconv.Atoi(ctx.Query("code"))
	req.Code = int32(code)
	service := api.Service{}
	response, err := service.StatusCode(context.TODO(), &req)
	if err != nil {
		st := status.Convert(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("内部错误: %v", st.Message()),
		})
		return
	}
	ret, _ := ma.MarshalToString(response)
	ctx.JSON(http.StatusOK, json.RawMessage(ret))
}

func (a HttpBinUS) Headers(ctx *gin.Context) {
	var req v1.HeadersReq
	service := api.Service{}
	response, err := service.Headers(context.TODO(), &req)
	if err != nil {
		st := status.Convert(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("内部错误: %v", st.Message()),
		})
		return
	}

	var m = make(map[string]*anypb.Any)
	for k, v := range ctx.Request.Header {
		n, _ := json.Marshal(v)
		m[k] = &anypb.Any{Value: n}
	}
	response.Data, _ = anypb.New(&v1.HeadersReply{
		Headers: m,
	})
	ret, err := ma.MarshalToString(response)
	fmt.Println(err)
	ctx.JSON(http.StatusOK, json.RawMessage(ret))
}

func (a HttpBinUS) IP(ctx *gin.Context) {
	var req v1.IPReq
	service := api.Service{}
	response, err := service.IP(context.TODO(), &req)
	if err != nil {
		st := status.Convert(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("内部错误: %v", st.Message()),
		})
		return
	}
	ret, _ := ma.MarshalToString(response)
	ctx.JSON(http.StatusOK, json.RawMessage(ret))
}

func (a HttpBinUS) UserAgent(ctx *gin.Context) {
	var req v1.UserAgentReq
	service := api.Service{}
	response, err := service.UserAgent(context.TODO(), &req)
	if err != nil {
		st := status.Convert(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": fmt.Sprintf("内部错误: %v", st.Message()),
		})
		return
	}
	n, _ := anypb.New(&v1.UserAgentReply{UserAgent: ctx.Request.Header.Get("User-Agent")})
	response.Data = n
	ret, _ := ma.MarshalToString(response)
	ctx.JSON(http.StatusOK, json.RawMessage(ret))
}
