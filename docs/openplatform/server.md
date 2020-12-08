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

* 注意： 对于`VerifyTicket`事件，程序会默认存储cache中，且自动回复“SUCCESS”,其他事件，可通过自定义相关`Handler`来进行相关的逻辑处理，具体查看下方[自定义消息处理器](#handler)

#### 推送事件

* 开放平台在给第三方平台推送的有基本事件如下

>   * 授权成功   -    authorized
>   * 更新授权   -    updateauthorized
>   * 取消授权   -    unauthorized
>   * VerifyTicket - componetn_verify_ticket
>   * 快速注册小程序 - notify_third_fasteregister


#### <span id="handler">自定义消息处理器</span>

自定义消息处理器示例：

```golang
    import(
        oserver "gitee.com/wallesoft/ewa/openplatform/server"
        gserver "gitee.com/wallesoft/ewa/kernel/server"
    )
  ...
  server.Push(handler,oserver.EVENT_COMPONENT_VERIFY_TICKET)
  // handler实现 Handler 接口的
  ...
  或者
  server.PushFunc(func(m *gserver.Message)interface{}{
      //m.GetString("ComponentVerifyTicket")
      //m.GetInt("CreateTime")  
      return true
  },oserver.EVENT_COMPONENT_VERIFY_TICKET)
```
> :warning: 注意：
>   * 0. **handler的调用顺序为倒序，即 **`先添加的后调用`****
>   * 1. 最后一个非空返回值将作为最终应答给用户的消息内容，如果中间某一个 handler 返回值 **false**, 则将终止整个调用链，不会调用后续的 handlers。
>   * 2. 传入的自定义 Handler 需要实现接口 `gitee.com/wallesoft/ewa/kernel/server - Handler`。
>   * 3. 第三方平台的事件需要回复“success”，所以自定义事件处理时，返回值统一返回 **return true**即可，除非你知道自己在做什么（具体参考第1条）。