# httpbin-go

> Web 项目最佳通讯协议约定


### 项目约定

1、api: 协议约定
  - proto 约定
  - buf generate
2、controller: 接口层
  - http 实现层
3、services: 实现层
  - 具体实现逻辑层
4、main.go: 入口

```shell
>> tree .

├── README.md
├── api
│   └── v1
│       ├── api.pb.go
│       ├── api.proto
│       └── api_grpc.pb.go
├── buf.gen.yaml
├── controller
│   └── api
│       └── controller.go
├── go.mod
├── go.sum
├── main.go
├── openapi.yaml
└── services
    └── api
        ├── cache.go
        └── service.go

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


### 飞书事件订阅

- 机器人
- 权限申请

