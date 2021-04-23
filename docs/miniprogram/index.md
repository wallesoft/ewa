### 入门

#### 在此之前

* 你已经阅读了[官方文档](https://developers.weixin.qq.com/miniprogram/dev/framework/)，并看懂理解文档中的内容。

#### 开始
* 初始化
```golang
    import(
        "net/http"
        "gitee.com/wallesoft/ewa/miniprogram"
    )
    func example(request *http.Request, writer *http.ResponseWriter) {
        // 具体参数配置信息查看[开放平台] open.weixin.qq.com
        weapp := miniprogram.New(miniprogram.Config{
            AppID:          "AppID",
            Secret:      "App secret",
        })
    }
```

* 之后相关接口调用都通过'weapp'实例进行，不在详细说明