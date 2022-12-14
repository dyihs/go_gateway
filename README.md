项目来源：https://github.com/e421083458/gin_scaffold
### 安装软件依赖
```
git clone git@github.com:e421083458/gin_scaffold.git
cd gin_scaffold
go mod tidy
```
确保正确配置了 conf/mysql_map.toml、conf/redis_map.toml：

### 运行

```
❯ go run main.go
------------------------------------------------------------------------
[INFO]  config=./conf/dev/
[INFO]  start loading resources.
[INFO]  success loading resources.
------------------------------------------------------------------------
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
[GIN-debug] GET    /ping                     --> go_gateway/router.InitRouter.func1 (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] GET    /demo/index               --> go_gateway/controller.(*DemoController).Index-fm (7 handlers)
[GIN-debug] GET    /demo/bind                --> go_gateway/controller.(*DemoController).Bind-fm (7 handlers)
[GIN-debug] POST   /demo/bind                --> go_gateway/controller.(*DemoController).Bind-fm (7 handlers)
[GIN-debug] PUT    /demo/bind                --> go_gateway/controller.(*DemoController).Bind-fm (7 handlers)
[GIN-debug] PATCH  /demo/bind                --> go_gateway/controller.(*DemoController).Bind-fm (7 handlers)
[GIN-debug] HEAD   /demo/bind                --> go_gateway/controller.(*DemoController).Bind-fm (7 handlers)
[GIN-debug] OPTIONS /demo/bind                --> go_gateway/controller.(*DemoController).Bind-fm (7 handlers)
[GIN-debug] DELETE /demo/bind                --> go_gateway/controller.(*DemoController).Bind-fm (7 handlers)
[GIN-debug] CONNECT /demo/bind                --> go_gateway/controller.(*DemoController).Bind-fm (7 handlers)
[GIN-debug] TRACE  /demo/bind                --> go_gateway/controller.(*DemoController).Bind-fm (7 handlers)
[GIN-debug] GET    /demo/dao                 --> go_gateway/controller.(*DemoController).Dao-fm (7 handlers)
[GIN-debug] GET    /demo/redis               --> go_gateway/controller.(*DemoController).Redis-fm (7 handlers)
[GIN-debug] POST   /api/login                --> go_gateway/controller.(*ApiController).Login-fm (7 handlers)
[GIN-debug] GET    /api/loginout             --> go_gateway/controller.(*ApiController).LoginOut-fm (7 handlers)
[GIN-debug] GET    /api/user/listpage        --> go_gateway/controller.(*ApiController).ListPage-fm (8 handlers)
[GIN-debug] POST   /api/user/add             --> go_gateway/controller.(*ApiController).AddUser-fm (8 handlers)
[GIN-debug] POST   /api/user/edit            --> go_gateway/controller.(*ApiController).EditUser-fm (8 handlers)
[GIN-debug] POST   /api/user/remove          --> go_gateway/controller.(*ApiController).RemoveUser-fm (8 handlers)
[GIN-debug] POST   /api/user/batchremove     --> go_gateway/controller.(*ApiController).RemoveUser-fm (8 handlers)
2022/12/01 17:52:38  [INFO] HttpServerRun::8880
```
### swagger文档生成

https://github.com/swaggo/swag/releases

**下载对应操作系统的执行文件到 $GOPATH/bin 下面**

```
~
❯ ll -r $GOPATH/bin
total 425520
-rwxr-xr-x  1 shiyd  staff    13M 12  1 00:19 swag
```

**设置接口文档参考：**

```
// ListPage godoc
// @Summary 测试数据绑定
// @Description 测试数据绑定
// @Tags 用户
// @ID /demo/bind
// @Accept  json
// @Produce  json
// @Param polygon body dto.DemoInput true "body"
// @Success 200 {object} middleware.Response{data=dto.DemoInput} "success"
// @Router /demo/bind [post]
```

生成接口文档：`swag init`
然后启动服务器：`go run main.go`，浏览地址: http://127.0.0.1:8880/swagger/index.html
