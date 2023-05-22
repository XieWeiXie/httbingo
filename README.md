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

- Dockerfile

> 多阶段构建

```shell

>> docker build -f Dockerfile -t httpbingo:v1 .

>> docker images | grep httpbin 

httpbin                                                       v1                                 71ac5a7bbce4        25 minutes ago      21.9MB

```


- Dive

> 镜像分层查看

```shell

>> dive httpbin:v1


┃ ● Layers ┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ │ Current Layer Contents ├──────────────────────────────────────────────────────────────
Cmp   Size  Command                                                                      Permission     UID:GID       Size  Filetree
    7.3 MB  FROM a30035fe3df17b5                                                         drwxr-xr-x         0:0     817 kB  ├── bin
    1.6 MB  apk update --no-cache && apk add --no-cache ca-certificates tzdata           -rwxrwxrwx         0:0        0 B  │   ├── arch → /bin/busybox
       0 B  #(nop) WORKDIR /app                                                          -rwxrwxrwx         0:0        0 B  │   ├── ash → /bin/busybox
     13 MB  #(nop) COPY dir:208dfa407a669df3c06fc31dd9e84f288481ee9f6d2647616d250c135bdd -rwxrwxrwx         0:0        0 B  │   ├── base64 → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── bbconfig → /bin/busybox
│ Layer Details ├─────────────────────────────────────────────────────────────────────── -rwxr-xr-x         0:0     817 kB  │   ├── busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── cat → /bin/busybox
Tags:   (unavailable)                                                                    -rwxrwxrwx         0:0        0 B  │   ├── chattr → /bin/busybox
Id:     a30035fe3df17b586df9ec4bf5996fa5920f686f42e3270038722884bcf977d5                 -rwxrwxrwx         0:0        0 B  │   ├── chgrp → /bin/busybox
Digest: sha256:bb01bd7e32b58b6694c8c3622c230171f1cec24001a82068a8d30d338f420d6c          -rwxrwxrwx         0:0        0 B  │   ├── chmod → /bin/busybox
Command:                                                                                 -rwxrwxrwx         0:0        0 B  │   ├── chown → /bin/busybox
#(nop) ADD file:7625ddfd589fb824ee39f1b1eb387b98f3676420ff52f26eb9d975151e889667 in /    -rwxrwxrwx         0:0        0 B  │   ├── cp → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── date → /bin/busybox
│ Image Details ├─────────────────────────────────────────────────────────────────────── -rwxrwxrwx         0:0        0 B  │   ├── dd → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── df → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── dmesg → /bin/busybox
Total Image size: 22 MB                                                                  -rwxrwxrwx         0:0        0 B  │   ├── dnsdomainname → /bin/busybox
Potential wasted space: 547 kB                                                           -rwxrwxrwx         0:0        0 B  │   ├── dumpkmap → /bin/busybox
Image efficiency score: 98 %                                                             -rwxrwxrwx         0:0        0 B  │   ├── echo → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── ed → /bin/busybox
Count   Total Space  Path                                                                -rwxrwxrwx         0:0        0 B  │   ├── egrep → /bin/busybox
    2        429 kB  /etc/ssl/certs/ca-certificates.crt                                  -rwxrwxrwx         0:0        0 B  │   ├── false → /bin/busybox
    2         93 kB  /lib/apk/db/installed                                               -rwxrwxrwx         0:0        0 B  │   ├── fatattr → /bin/busybox
    2         25 kB  /lib/apk/db/scripts.tar                                             -rwxrwxrwx         0:0        0 B  │   ├── fdflush → /bin/busybox
    2         288 B  /lib/apk/db/triggers                                                -rwxrwxrwx         0:0        0 B  │   ├── fgrep → /bin/busybox
    2         141 B  /etc/apk/world                                                      -rwxrwxrwx         0:0        0 B  │   ├── fsync → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── getopt → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── grep → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── gunzip → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── gzip → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── hostname → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── ionice → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── iostat → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── ipcalc → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── kbd_mode → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── kill → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── link → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── linux32 → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── linux64 → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── ln → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── login → /bin/busybox
                                                                                         -rwxrwxrwx         0:0        0 B  │   ├── ls → /bin/busybox

```


```shell

>> docker history 86b7571bf4e5

IMAGE               CREATED             CREATED BY                                      SIZE                COMMENT
86b7571bf4e5        3 hours ago         /bin/sh -c #(nop)  CMD ["/app/httpbingo"]       0B
aa8c73dde838        3 hours ago         /bin/sh -c #(nop)  EXPOSE 9091                  0B
d8319546655f        3 hours ago         /bin/sh -c #(nop) COPY dir:4331fcfade96bce65…   13MB
d8cd0444d48c        3 hours ago         /bin/sh -c #(nop) WORKDIR /app                  0B
4c8f3da5fbee        3 hours ago         /bin/sh -c apk update --no-cache && apk add …   1.64MB
5e2b554c1c45        12 days ago         /bin/sh -c #(nop)  CMD ["/bin/sh"]              0B
<missing>           12 days ago         /bin/sh -c #(nop) ADD file:7625ddfd589fb824e…   7.33MB

```

