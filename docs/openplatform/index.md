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

---
#### 获取用户授权页 URL (mobile) 
```golang

url := op.GetMobilePreAuthorizationUrl("https://test.com/callback")

```


#### 使用授权码换取接口调用凭据和授权信息

在用户在授权页授权流程完成后，授权页会自动跳转进入回调URI，并在URL参数中返回授权码和过期时间，如：(https://test.com/callback?auth_code=xxx&expires_in=600)

```golang
   res :=  op.HandleAuthorize(authcode string);
```



#### 获取授权方的帐号基本信息

```golang
info := op.GetAuthorizer(appid string);
```
> appid 为授权appid
#### 获取授权方的选项设置信息

```golang
info := op.GetAuthorizerOption(appid string,name string);
```

#### 设置授权方的选项信息

```golang
 res := op.SetAuthorizerOption(appid string, name string, value string);
```

> 该API用于获取授权方的公众号或小程序的选项设置信息，如：地理位置上报，语音识别开关，多客服开关。注意，获取各项选项设置信息，需要有授权方的授权，详见权限集说明。


#### 获取已授权的授权方列表

```golang
list := op.GetAuthorizers(offset int, count int)
```

#### :man:另外

> 如果想获取 access_token 等数据，用以下方法

##### 获取verify_ticket

```golang
ticket := op.GetVerifyTicket()
```
##### 获取component_access_token

```golang
token := op.GetAccessToken()
```
##### 获取pre_auth_code

```golang

code, err := op.GetPreAuthCode()

```