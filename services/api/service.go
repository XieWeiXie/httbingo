package api

import (
	"context"
	"github.com/XieWeiXie/httpbin-go/api/v1"
	"net"
)

type Service struct{}

func (s Service) HttpMethodGet(ctx context.Context, req *v1.HttpMethodGetReq) (*v1.ResponseReply, error) {
	var reply = new(v1.ResponseReply)
	cache, err := defaultCache.GetCache(int(req.Id))
	if err != nil {
		return reply, err
	}
	reply.Code = 0
	reply.Data = &v1.HttpMethodGetReply{
		Id:   int32(cache.Id),
		Name: cache.Name,
	}
	return reply, nil
}

func (s Service) HttpMethodPut(ctx context.Context, req *v1.HttpMethodPutReq) (*v1.ResponseReply, error) {
	var reply = new(v1.ResponseReply)
	cache, err := defaultCache.PutCache(int(req.Id), req.Name)
	if err != nil {
		return reply, err
	}
	reply.Data = &v1.HttpMethodPutReply{
		Id:   int32(cache.Id),
		Name: cache.Name,
	}
	return reply, nil
}

func (s Service) HttpMethodDelete(ctx context.Context, req *v1.HttpMethodDeleteReq) (*v1.ResponseReply, error) {
	var reply = new(v1.ResponseReply)
	_, err := defaultCache.DeleteCache(int(req.Id))
	if err != nil {
		return reply, err
	}
	reply.Data = &v1.HttpMethodDeleteReply{}
	return reply, nil

}

func (s Service) HttpMethodPost(ctx context.Context, req *v1.HttpMethodPostReq) (*v1.ResponseReply, error) {
	var reply = new(v1.ResponseReply)
	cache, err := defaultCache.PostCache(req.Name)
	if err != nil {
		return reply, err
	}
	reply.Data = &v1.HttpMethodPostReply{
		Id:   int32(cache.Id),
		Name: cache.Name,
	}
	return reply, nil
}

func (s Service) StatusCode(ctx context.Context, req *v1.StatusCodeReq) (*v1.ResponseReply, error) {
	var reply = new(v1.ResponseReply)
	reply.Data = &v1.StatusCodeReply{Code: int32(req.Code)}
	return reply, nil
}

func (s Service) Headers(ctx context.Context, req *v1.HeadersReq) (*v1.ResponseReply, error) {
	var reply = new(v1.ResponseReply)
	reply.Data = &v1.HeadersReply{Headers: make(map[string]string)}
	return reply, nil
}

func (s Service) IP(ctx context.Context, req *v1.IPReq) (*v1.ResponseReply, error) {
	var reply = new(v1.ResponseReply)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return reply, err
	}
	var origin string
	for _, value := range addrs {
		if ip, ok := value.(*net.IPNet); ok && !ip.IP.IsLoopback() && ip.IP.To4() != nil {
			origin = ip.IP.String()
			break
		}
	}
	reply.Data = &v1.IPReply{Origin: origin}
	return reply, nil
}

func (s Service) UserAgent(ctx context.Context, req *v1.UserAgentReq) (*v1.ResponseReply, error) {
	var reply = new(v1.ResponseReply)
	reply.Data = &v1.UserAgentReply{UserAgent: ""}
	return reply, nil
}
