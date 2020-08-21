### 服务端


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