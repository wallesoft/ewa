### 服务端

#### 入门

--------
* 服务端主要涉及到微信推送消息/事件，通知等处理，涵盖公众号，小程序，第三方平台等，由于消息/事件等处理逻辑基本一致，将逻辑总结后统一放在此处说明处理
* 开始之前：**以下官方文档需要仔细阅读并理解其内容/特别注意其消息格式**
> * 公众号消息加解密说明：[官方文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Message_encryption_and_decryption_instructions.html)
> * 公众号接收普通消息: [官方文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_standard_messages.html)
> * 公众号接收事件消息: [官方文档](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_event_pushes.html)
> * 小程序消息推送: [官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html)
> * 小程序客服消息: [官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/customer-message/customer-message.html)
> * 小程序内容安全- 异步检查图片/音频通知: [官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/sec-check/security.mediaCheckAsync.html)
> * 第三方平台消息加解密: [官方文档](https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Message_Encryption/Message_encryption_and_decryption.html)
> * 第三平台其推送消息如 **快速创建小程序** **授权更新** 等等通知事件/消息在官方文档相关章节均有说明，:warning: 请仔细阅读
------

#### 开始

#### 简单示例：

* 自定义消息处理器(以接收第三方平台 component_verify_ticket 为例)：

```golang
    import(
        oserver "gitee.com/wallesoft/ewa/openplatform/server"
        gserver "gitee.com/wallesoft/ewa/kernel/server"
    )
  ...
  server.Push(handler,oserver.EVENT_COMPONENT_VERIFY_TICKET)
  // handler为实现 Handler 接口的
  // Handler接口具体查看 "https://godoc.org/gitee.com/wallesofte/ewa/kernel/server"
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