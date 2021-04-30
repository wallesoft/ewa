### 微信支付

#### 之前

----
* 采用微信支付v3.0开发，在开始之前，请仔细阅读[微信支付apiv3.0相关文档](https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pages/index.shtml)、[接口规则](https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay-1.shtml)以及常见问题等，其中的一些术语及相关技术不在此进行解释
----

#### 开始

* 初始化配置
```golang
    import (
        "gitee.com/wallesoft/ewa/payment"
    )
    func main() {
        pay := payment.New(payment.Config{
            AppID:          "your-appid",  // 直连商户申请的公众号或移动应用appid。
            MchID:          "your-mchid", // 商户mchid
            SerialNo:       "your-serialno",// 证书编号注意是平台证书编号不是商户证书编号
            PriCertPath:    "to-your-private-cert-path", //商户私钥证书地址 绝对地址！！！
            PubCertPath:    "to-your-public-cert-path", //商户公钥证书地址  绝对地址！！！
            ApiV3Key:       "apiv3key", //商户apiv3秘钥
    
            PFCertSavePath: "wechat-platform-saver-folder-path", //平台证书下载存放地址 绝对地址！！！且到目录即可 默认：/etc/wechatpay/
        })
    }
    
 ```   
 > 商户平台公共证书注意首次需要下载，否则容易形成“死循环”具体解释查看官方文档及首次下载证书说明：https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml

#### 获取平台证书列表

* 获取平台可用证书列表并将证书保存到配置的[平台证书下载存放地址], 确保配置的地址目录可写