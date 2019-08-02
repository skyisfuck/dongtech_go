设置proxy

export GO111MODULE=on

export GOPROXY=https://goproxy.io

go get组件

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

问题：

ambiguous import: found github.com/ugorji/go/codec in multiple modules

解决：

go get github.com/ugorji/go@v1.1.2

启动main.go即可

测试接口 

localhost:9090/version 

localhost:9090/config

localhost:9090/getUser/1

localhost:9090/grpc
