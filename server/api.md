### vben-admin 1.0.0 API接口文档

- web 端接口前缀: /api/v2

> 接口用户认证用cookie/query/header中的 x-token 进行用户认证
> 
> header 中必带信息
> X-TOKEN: token 信息
>

#### 返回数据格式
```
{
    "code":"0",
    "msg":"Success",
    "data": {
        "key":"value"
    }
}
```
> code 为0时为正常返回  返回其他时为异常返回
> 接口文档描述中一般只描述data结构


### 返回码
```
// CodeSuccess 请求成功
CodeSuccess = 0
// CodeCreated 已创建
CodeCreated = 201

// CodeRedirect 301 跳转
CodeRedirect = 301
// CodeUnauthorized 未登录
CodeUnauthorized = 401
// CodeForbidden 禁止访问
CodeForbidden = 403
// CodeNotFound 资源未找到
CodeNotFound = 404

//CodeParamErr 各种奇奇怪怪的参数错误
CodeParamErr       = 40001
CodeUsernamePwdErr = 40002

// CodeSignExpired 签名过期
CodeSignExpired = 40109
// CodeExist 数据已存在
CodeExist               = 40002
CodeInsufficientBalance = 40003

// CodeSystemError 系统错误
CodeSystemError = 50001

```

