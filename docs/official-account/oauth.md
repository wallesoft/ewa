### 网页授权

#### 关于 OAuth2.0

OAuth是一个关于授权（authorization）的开放网络标准，在全世界得到广泛应用，目前的版本是2.0版。

> 摘自：[RFC 6749](https://datatracker.ietf.org/doc/rfc6749/?include_text=1)

步骤解释：

    （A）用户打开客户端以后，客户端要求用户给予授权。
    （B）用户同意给予客户端授权。
    （C）客户端使用上一步获得的授权，向认证服务器申请令牌。
    （D）认证服务器对客户端进行认证以后，确认无误，同意发放令牌。
    （E）客户端使用令牌，向资源服务器申请获取资源。
    （F）资源服务器确认令牌无误，同意向客户端开放资源。

关于 OAuth 协议我们就简单了解到这里，如果还有不熟悉的同学，请 [*百度](https://www.baidu.com/s?wd=oauth2.0)

#### 微信 OAuth

在微信里的 OAuth 其实有两种：[公众平台网页授权获取用户信息](http://mp.weixin.qq.com/wiki/9/01f711493b5a02f24b04365ac5d8fd95.html)、[开放平台网页登录](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419316505&token=&lang=zh_CN)。

它们的区别有两处，授权地址不同，`scope` 不同。

>  - **公众平台网页授权获取用户信息**
    **授权 URL**: `https://open.weixin.qq.com/connect/oauth2/authorize`
    **Scopes**: `snsapi_base` 与 `snsapi_userinfo`

>  - **开放平台网页登录**
    **授权 URL**: `https://open.weixin.qq.com/connect/qrconnect`
    **Scopes**: `snsapi_login`

他们的逻辑都一样：

  1. 用户尝试访问一个我们的业务页面，例如: `/user/profile`
  2. 如果用户已经登录，则正常显示该页面
  3. 系统检查当前访问的用户并未登录（从 session 或者其它方式检查），则跳转到**跳转到微信授权服务器**（上面的两种中一种**授权 URL**  ），并告知微信授权服务器我的**回调URL（redirect_uri=http://domain/example/callback)**，此时用户看到蓝色的授权确认页面（`scope` 为 `snsapi_base` 时不显示）
  4. 用户点击确定完成授权，浏览器跳转到**回调URL**: `callback` 并带上 `code`： `?code=CODE&state=STATE`。
  5. 在 `callback` 中得到 `code` 后，通过 `code` 再次向微信服务器请求得到 **网页授权 access_token** 与 `openid`
  6. 你可以选择拿 `openid` 去请求 API 得到用户信息（可选）
  7. 将用户信息写入 SESSION。
  8. 跳转到第 3 步写入的 `target_url` 页面（`/user/profile`）。

>
>
#### 逻辑组成

从上面我们所描述的授权流程来看，我们至少有3个页面：

  1. **业务页面**，也就是需要授权才能访问的页面。
  2. **发起授权页**，此页面其实可以省略，可以做成一个中间件，全局检查未登录就发起授权。
  3. **授权回调页**，接收用户授权后的状态，并获取用户信息，写入用户会话状态（SESSION）。

#### 开始之前

在开始之前请一定要记住，先登录公众号后台，找到**边栏 “开发”** 模块下的 **“接口权限”**，点击 **“网页授权获取用户基本信息”** 后面的修改，添加你的网页授权域名。

> 如果你的授权地址为：`http://www.abc.com/xxxxx`，那么请填写 `www.abc.com`，也就是说请填写与网址匹配的域名，前者如果填写 `abc.com` 是通过不了的。

#### 使用说明
    由于OAuth2.0授权的通用性，我们将网页授权单独拿了出来具体查看
    https://gitee.com/wallesoft/go/oauth2
    文档地址：https://pkg.go.dev/gitee.com/wallesoft/go/oauth2#WechatConfig

#### 发起授权

```golang
    //基本用法 scopes默认为 "snsapi_login"
    redirectUrl := oa.OAuth().GetAuthURL("http://domain/example/callback")

    // 设置scopes 默认scope 为 "snsapi_login"
    redirectUrl := oa.OAuth().Scope([]string{"snsapi_base"}).GetAuthURL("http://domain/example/callback")

    // 设置state参数 不设置为默认生成的随机字符串
    redirectUrl := oa.OAuth().Scope([]string{"snsapi_base"}).GetAuthURL("http://domain/example/callback","set_state_here")
```
* 获取到redirectUrl后，请自行进行302跳转

#### 获取以授权用户

```golang

    code := "微信回调URL携带的 code"
    user := oa.OAuth().GetUserFromCode(code)

```

 返回的user为gitee.com/wallesoft/go/oauth2 中 *oauth2.User 可参考文档

其中user可以用的方法有：

> user.GetID()         对应的微信的openid
>
> user.GetNickName()   对应微信的nickname
>
> user.GetName() 对应微信的nickname
>
> user.GetAvatar() 对应微信的headimg 头像地址
>
> user.GetRaw().MustToString() 原始信息 string格式
>
> user.GetAccessToken() \ user.GetRefreshToken() \user.GetExpiresIn()
> 
> 其他可用的方法查看oauth2.User文档进行参考


