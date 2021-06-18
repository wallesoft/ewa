### 公众号

#### 入门

> 微信公众号官方文档[https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html) ,建议仔细阅读


#### 开始
* 初始化
```golang

    import(
        ....
        "gitee.com/wallesoft/ewa/officialaccount"
    )

    func main() {
        oa := officialaccount.New(officialaccount.Config{
            AppID:  "your-appid", // 公众号appid 公众号平台获取
            Secret: "scret",      // 公众号secret 公众号平台获取，注意保存
        })
        // 之后的示例用 oa 标识，即不在进行初始化展示
    }
```
* :warning: 之后的示例用 oa 标识，即不在进行初始化展示