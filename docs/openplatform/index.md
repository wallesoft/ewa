### 开放平台

#### 入门

----
* 涉及的接口相关信息参考：[授权流程技术说明 - 官方文档](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1453779503&token=&lang=)
----

#### 开始
* 平台初始化配置
```golang
    import(
        "net/http"
        "gitee.com/wallesoft/ewa/openplatform"
    )
    func example() {
        // 具体参数配置信息查看[开放平台] open.weixin.qq.com
        op := openplatform.New(openplatform.Config{
            AppID:          "AppID",
            AppSecret:      "App secret",
            Token:          "token",
            EncodingAESKey: "Aes key", 
        })
    }
```

#### 获取用户授权页 URL

```golang

url := op.GetPreAuthorizationUrl("https://test.com/callback")

```
:warning: 注意：
```
GetPreAuthorizationUrl(callback string,optional ...map[string]interface{}) string{}
```
还有可选参数, 具体查看[官方文档](https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Authorization_Process_Technical_Description.html)说明