# httpbin-go

> Web 项目最佳接口协议约定


### 项目约定

1、api: 协议约定
  - proto 约定
  - buf generate

2、controller: 接口层
  - http 实现层

3、services: 实现层
  - 具体实现逻辑层

4、main.go: 入口

5、third_party: 第三方 proto

6、openapi.yaml: 接口定义文件

7、s.yaml: serverless 配置文件

```shell
>> tree .
├── README.md
├── api
│   └── v1
│       ├── api.pb.go
│       ├── api.proto
│       └── api_grpc.pb.go
├── buf.gen.yaml
├── buf.work.yaml
├── buf.yaml
├── controller
│   └── api
│       └── controller.go
├── go.mod
├── go.sum
├── main.go
├── openapi.yaml
├── s.yaml
├── services
│   └── api
│       ├── cache.go
│       └── service.go
└── third_party
    ├── README.md
    ├── errors
    │   └── errors.proto
    ├── google
    │   ├── api
    │   │   ├── annotations.proto
    │   │   ├── client.proto
    │   │   ├── field_behavior.proto
    │   │   ├── http.proto
    │   │   └── httpbody.proto
    │   └── protobuf
    │       ├── any.proto
    │       ├── api.proto
    │       ├── compiler
    │       │   └── plugin.proto
    │       ├── descriptor.proto
    │       ├── duration.proto
    │       ├── empty.proto
    │       ├── field_mask.proto
    │       ├── source_context.proto
    │       ├── struct.proto
    │       ├── timestamp.proto
    │       ├── type.proto
    │       └── wrappers.proto
    ├── openapi
    │   └── v3
    │       ├── annotations.proto
    │       └── openapi.proto
    └── validate
        ├── README.md
        └── validate.proto
```


### Protobuf

```shell
buf.work.yaml
buf.gen.yaml
buf.yaml

>> buf generate

```


### 内网穿透

- cpolar

```shell
>> cpolar authtoken *****
>> cpolar http 9091

>> curl https://317723e6.r3.cpolar.cn/httpbin-v1/heart/beat | jq .

{
  "code": 0,
  "data": "2023-05-10 22:28:32",
  "message": ""
}

```


- ngrok


```shell
>> ngrok config add-authtoken **
>> ngrok http 9091

>> curl https://f695-2409-895a-6a0-14e7-2d-73cd-ffac-7d17.ngrok-free.app/http-bin/v1/heart/beat

{
  "code": 0,
  "data": "2023-05-10 22:51:28",
  "message": ""
}

```
