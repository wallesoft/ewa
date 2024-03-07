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
    }
```

* 注意： 对于`VerifyTicket`事件，程序会默认处理并缓存(缓存在文件中，如需要其他缓存方式如redis，内存等，可查看[gcache](https://www.goframe.org/os/gcache/index))，如果不处理其他事件可直接这样使用：

```golang
    server.Serve()
```

* 对于其他事件，可通过自定义相关`Handler`来进行相关的逻辑处理，具体查看下方[自定义消息处理器](#handler)

#### 推送事件

* 开放平台在给第三方平台推送的有基本事件如下

>   * 授权成功   -    authorized
>   * 更新授权   -    updateauthorized
>   * 取消授权   -    unauthorized
>   * VerifyTicket - component_verify_ticket
>   * 快速注册小程序 - notify_third_fasteregister


#### <span id="handler">自定义消息处理器</span>
#### 消息处理

具体内容到 [服务端](/server/index) 章节查看

