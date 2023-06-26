### 入门

#### 在此之前

* 你已经阅读了[官方文档](https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/server-api-introduction),并理解文档当中的内容

#### 开始

* 初始化
```golang
    import "gitee.com/wallesoft/ewa/bytedance/miniprogram"
    func example() {
        miniapp := miniprogram.New(miniprogram.Config{
            AppID: "appid",
            Secret: "app secret"
            // 更多配置，请查看官方文档和miniprogram.Config
        })
    }
```

* 之后的相关接口调用均通过'miniapp'实例进行，不在详细说明