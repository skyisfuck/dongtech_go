****简介****

本项目集成了如下开源工具包：
cobra、
gin、
viper、
cron、
toml、
grpc、
postgres、
gorm、
pg、
uuid、
jwt、
set、
captcha、
email、
csv
......

****1.设置proxy****

export GO111MODULE=on

export GOPROXY=https://goproxy.io

****2.go get组件****

go get github.com/spf13/viper

go get -u github.com/spf13/cobra/cobra

go get github.com/gin-gonic/gin

go get github.com/kardianos/govendor

go get github.com/robfig/cron

go get github.com/BurntSushi/toml

go get github.com/pkg/errors

go get google.golang.org/grpc

go get github.com/golang/protobuf/protoc-gen-go

go get github.com/go-sql-driver/mysql

go get github.com/jinzhu/gorm

go get github.com/jinzhu/gorm/dialects/postgres

go get github.com/go-pg/pg

go get github.com/satori/go.uuid

go get github.com/dgrijalva/jwt-go

go get github.com/deckarep/golang-set

go get github.com/dchest/captcha

go get github.com/go-gomail/gomail

****3.问题****

1.ambiguous import: found github.com/ugorji/go/codec in multiple modules

解决：

go get github.com/ugorji/go@v1.1.2

2.生成 pb.go
  
命令行执行:protoc --go_out=plugins=grpc:. ServeRoute.proto

****4.运行****

启动main.go即可

****5.测试接口****

`1:gin接口测试`

http://localhost:9090/version 

`2:gin接口测试`

http://localhost:9090/config 

`3:gorm接口测试`

http://localhost:9090/getUser/1

`4:grpc接口测试`

http://localhost:9090/grpc

`5:uuid接口测试`

http://localhost:9090/uuid

`6:定时任务接口测试`

cron.go

`7:html接口测试`

http://localhost:9090/index

`8:jwt接口测试`

`8.1:jwt 登陆接口测试`
http://localhost:9090/login/dong/123456

`8.2:jwt 登陆验证接口测试`
http://localhost:9090/verify/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA1MTIyMTAsImlhdCI6MTU2MDUwODYxMCwidXNlcl9pZCI6MSwicGFzc3dvcmQiOiIxMjM0NTYiLCJ1c2VybmFtZSI6ImRvbmciLCJmdWxsX25hbWUiOiJkb25nIiwicGVybWlzc2lvbnMiOltdfQ.Esh1Zge0vO1BAW1GeR5wurWP3H1jUIaMf3tcSaUwkzA

`8.3:jwt 刷新token接口测试`
//http://localhost:9090/refresh/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA1MTIyNDMsImlhdCI6MTU2MDUwODYxMCwidXNlcl9pZCI6MSwicGFzc3dvcmQiOiIxMjM0NTYiLCJ1c2VybmFtZSI6ImRvbmciLCJmdWxsX25hbWUiOiJkb25nIiwicGVybWlzc2lvbnMiOltdfQ.Xkb_J8MWXkwGUcBF9bpp2Ccxp8nFPtRzFzOBeboHmg0
