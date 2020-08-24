### 服务端

#### 说明
* 由于各框架对 **`请求输入`** 和 **`请求输出`** 的实现各不相同，对于微信推送消息的处理将不包含相关请求输入及输出模块的实现，取而代之的是将相关处理逻辑排除在此框架外，以 Request* 及 Response* 结构体代替，具体结构及实现方法参考 [`ServerGuard`]() 模块详细介绍

#### 推送事件

* 开放平台在给第三方平台推送的有4个基本事件

    * 授权成功      authorized
    * 更新授权      updateauthorized
    * 取消授权      unauthorized
    * VerifyTicket componetn_verify_ticket


#### 用法示例

* 对于事件 **component_verify_ticket**,默认处理会将 **verify_ticket** 缓存，以下是部分代码

```golang
import (
    "gitee.com/wallesoft/ewa/openplatform/server"
)


    .....

    request := server.CreateRequest()
    // 将微信post过来的相关参数映射给request
    // reqeust.Timestamp
    // reqeust.Noce
    // reqeust.EncryptType
    // reqeust.MsgSignature
    // reqeust.BodyRaw
    message := server.GetMessageFromRequest(request)
    response := server.Serve(message)
    
    ....

```
#### 框架示例 [(goframe框架)](https://www.goframe.org/)
```golang
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

```
#### 自定义消息处理器