
## TODO 
`lsof -i:8489 | grep gin-admin | awk '{print $2}' | xargs kill -9`

[x] 权限设计
[ ] 数据库设计
[ ] 前端页面

## 基础依赖
- mysql
- redis
- golang `https://go.dev/`
- node `https://nodejs.org/en/download/package-manager`
- 安装yarn `https://yarn.bootcss.com/docs/install/index.html#mac-stable`
- 安装gowatch `https://github.com/silenceper/gowatch`


## git克隆代码

## 运行后端

### 创建本地开发数据库

### 进入项目 -> server
- 1.安装go 项目依赖 `go get .`
- 2.复制 `server/config.yaml.template  -> server/config.yaml`
- 3.修改数据对应配置
- 运行go项目 `gowatch`

## 运行前端

### 进去前端项目  -> web
- 安装前端依赖 `yarn`
- 运行前端项目 `yarn watch`