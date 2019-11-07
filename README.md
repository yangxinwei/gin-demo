###依赖文件
```
export GO111MODULE=on
go mod vendor
```

###swagger
https://github.com/swaggo/gin-swagger
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles

swag init

```

####文档访问地址
http://localhost:8080/swagger/index.html

###xorm
```
go get -u github.com/go-xorm/xorm
go get -u github.com/go-xorm/cmd

cd $GOPATH/src/github.com/go-xorm/cmd/xorm
xorm help reverse
```
example:
1、整个数据库所有表
```
xorm reverse mysql "{username}:{password}@({host}:{port})/{dbname}?charset=utf8" templates/goxorm
```
会在默认路径./models 目录下 生成整个数据库所有表的对应文件

2、指定表和导出路径
```
xorm reverse mysql "{username}:{password}@({host}:{port})/{dbname}?charset=utf8" templates/goxorm entity  "{tablename}"
```

会在./entity 目录下 生成 {tablename}.go 文件