```
├── api // 接口逻辑定义
│   ├── admin // 模块
│   │   ├── admin.go
│   │   ├── dept.go
│   │   ├── options.go
│   │   ├── permission.go
│   │   ├── role.go
│   │   └── types // 接口返回数据结构定义
│   │       ├── admin.go
│   │       ├── dept.go
│   │       ├── options.go
│   │       ├── permissions.go
│   │       └── role.go
│   ├── base.go // 接口基础类型和方法定义
│   └── types.go // 接口基础数据结构定义
├── api.md // swagger 文件基础数据
├── config.yaml.template // 配置文件模板
├── crontab // 定时任务模块
│   └── cron.go // 定时任务示例
├── dao // 数据库操作层
│   ├── gen // 代码自动生成器 https://gorm.io/gen/  可以在server 目录下执行   go run dao/gen/gen.go 来自动生成代码
│   │   └── gen.go
│   ├── init.go // 数据库初始配置
│   ├── migration // 数据迁移工具
│   │   ├── migration.go // 入口
│   │   ├── permissions.go // 权限点定义
│   │   ├── setting.go // 公用配置
│   │   └── version.go // 迁移版本,用来控制是否迁移数据库, 需要迁移需要修改版本号
│   ├── models // 数据库模型定义 https://gorm.io/docs/
│   │   ├── address.go
│   │   ├── base.go // 模型基础字段
│   │   ├── base_test.go
│   │   ├── casbin_rule.go
│   │   ├── dept.go
│   │   ├── login_log.go
│   │   ├── permissions.go
│   │   ├── product.go
│   │   ├── region.go
│   │   ├── role.go
│   │   ├── setting.go
│   │   ├── user.go
│   │   └── warehouse.go
│   ├── query // gorm/gen 自动生成的脚本
│   │   ├── address_relations.gen.go
│   │   ├── addresses.gen.go
│   │   ├── casbin_rule.gen.go
│   │   ├── depts.gen.go
│   │   ├── gen.go
│   │   ├── login_logs.gen.go
│   │   ├── permissions.gen.go
│   │   ├── regions.gen.go
│   │   ├── role_permissions.gen.go
│   │   ├── roles.gen.go
│   │   ├── settings.gen.go
│   │   ├── user_roles.gen.go
│   │   ├── users.gen.go
│   │   ├── warehouse_areas.gen.go
│   │   ├── warehouse_cells.gen.go
│   │   ├── warehouse_tenants.gen.go
│   │   └── warehouses.gen.go
│   └── repo // 数据库操作逻辑
│       ├── base.go // 基础数据定义
│       ├── casbin.go // 权限控制相关
│       ├── dept.go // 部门相关
│       ├── permission.go // 权限
│       ├── role.go // 角色
│       ├── setting.go // 设置
│       ├── user.go // 用户
│       └── warehouse.go // 仓库
├── global // 全局定义
│   └── global.go
├── go.mod
├── go.sum
├── gowatch.yml // gowatch 配置
├── locales // 国际化配置
│   ├── init.go
│   ├── keys.go
│   ├── lang // 配置这个就好(暂时没处理这块)
│   │   ├── en.json
│   │   └── zh.json
│   ├── lang.go
│   └── locales.json
├── logs
├── main.go // 程序入口
├── middleware  // 路由中间件 https://gin-gonic.com/docs/
│   ├── auth.go
│   ├── language.go
│   ├── limit.go
│   ├── log.go
│   └── logger.go
├── pkg // 相关包
│   ├── auth // 认证
│   │   ├── hmac.go
│   │   └── jwt.go
│   ├── cache // 缓存
│   │   └── cache.go
│   ├── conf // 配置
│   │   ├── config.go
│   │   ├── cors.go
│   │   ├── db.go
│   │   ├── mysql.go
│   │   ├── path.go
│   │   ├── pgsql.go
│   │   ├── redis.go
│   │   ├── sqlite.go
│   │   └── system.go
│   ├── constant // 常量定义
│   │   ├── constant.go
│   │   └── setting
│   │       └── setting.go
│   ├── e // 错误定义
│   │   ├── error.go
│   │   └── errors.go
│   ├── lock // 操作锁
│   │   ├── redis.go
│   │   └── redis_test.go
│   ├── logger // 日志
│   │   ├── hook.go
│   │   └── logger.go
│   ├── shutdown // 关闭信号监听
│   │   └── shutdown.go
│   ├── utils // 工具
│   │   ├── parse.go
│   │   ├── path.go
│   │   ├── strings.go
│   │   └── time.go
│   └── validators // 表单验证
│       ├── msg.go
│       └── validator.go
├── routers // 路由
│   ├── admin.go
│   ├── router.go
│   ├── swagger // swagger 接口文档文件 自动生成
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   └── swagger.go
├── tasks // 异步任务处理
│   ├── init.go
│   ├── permissions.go
│   ├── tasks.go
│   └── test.go
└── types // 通用类型定义
    └── base.go

```


- EXCEL 包
`go get github.com/xuri/excelize/v2`

- PDF 包
`https://wkhtmltopdf.org/` html to pdf
`https://github.com/jung-kurt/gofpdf`