### 代小程序实现业务

#### 说明
* 在开始之前，需要仔细阅读官方文档[【代小程序实现业务】](https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Intro.html)

#### 开始
* 初始化小程序

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
        weapp := op.MiniProgram("authorizer_appid","authorizer_refresh_token")
    }
```
* 初始化后，同小程序相关的业务调用，通过 'weapp' 调用相关接口，调用方法同小程序接口，详细查看小程序相关文档 [小程序](/miniprogram/index)