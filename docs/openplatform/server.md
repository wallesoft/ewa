### 服务端

#### 说明
* 由于各框架对 **`请求输入`** 和 **`请求输出`** 的实现各不相同，但基本都是在包 ['net/http'](https://godoc.org/net/http) 的基础上进行实现的，所以对请求的输入和输出分为 **`http.Request`** 和 **`http.ResponseWriter`**
具体的方法等参考['net/http'](https://godoc.org/net/http)
#### 开始
* 平台初始化配置
```golang
    import(
        "net/http"
        "gitee.com/wallesoft/ewa/openplatform"
    )
    func example(request *http.Request, writer *http.ResponseWriter) {
        // 具体参数配置信息查看[开放平台] open.weixin.qq.com
        op := openplatform.New(openplatform.Config{
            AppID:          "AppID",
            AppSecret:      "App secret",
            Token:          "token",
            EncodingAESKey: "Aes key", 
        })
        server := op.Server(request,writer)
        server.Serve()
    }
```

* 注意： 对于`VerifyTicket`事件，程序会默认存储cache中，且自动回复“SUCCESS”,其他事件，可通过自定义相关`Handler`来进行处理并回复，具体查看下方[自定义消息处理器](#handler)

#### 推送事件

* 开放平台在给第三方平台推送的有4个基本事件

    * 授权成功      authorized
    * 更新授权      updateauthorized
    * 取消授权      unauthorized
    * VerifyTicket componetn_verify_ticket


<!--
#### 用法示例

* 对于事件 **component_verify_ticket**,默认处理会将 **verify_ticket** 缓存，以下是示例
```golang

    .....

    request := server.NewRequest()
    // 将微信post过来的相关参数映射给request
    // reqeust.Timestamp
    // reqeust.Noce
    // reqeust.EncryptType
    // reqeust.MsgSignature
    // reqeust.RawBody
    // request.Uri // 非必须参数
    //message := server.GetMessageFromRequest(request)
    response := server.Serve(message)
    
    ....

```
#### 框架示例 [(goframe框架)](https://www.goframe.org/)
<!-- ```golang
package main 

import (
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
    "gitee.com/wallesoft/ewa/openplatform/server"
)
func main(){
    s := g.Server()
    s.BindHandler("/notify",func(r *ghttp.Request) {
        
        // 此处可以绑定相关信息处理函数
        // 详细看下一小节 自定义消息处理器
        
        request := server.CreateRequest()
        if err := r.Parse(&request); err != nil {
            // 错误处理
        }
        request.BodyRaw = r.GetBody()
        message := server.GetMessageFromRequest(request)
        response:= server.Serve(message)

        r.Reponse.WriteExit(response.GetContent())
    })
}

``` -->
#### <span id="handler">自定义消息处理器</span>