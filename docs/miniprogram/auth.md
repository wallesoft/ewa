### 微信登录


#### 请求方法

* API: weapp.Session(ctx,code)

根据js_code获取用户session信息 [官方文档](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html)

#### 返回值
>   **❗❗❗特别说明❗❗❗**
>   **union_id** 不是每次都会有，只有在特定条件，详细查看[union_id机制说明](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/union-id.html)

kernel.http.ResponseData

#### 示例

```golang
    res := weapp.Session(ctx, code)
    
    res.Get("session_key").String()
    res.Get("openid").String()
    res.Get("errcode").Int()
    ....
    // 更多查看官方文档
    // 或者
    type SessionRes struct {
        SessionKey string   `json:"session_key"`
        UnionId string      `json:"union_id"`
        OpenId string       `json:"openid"`
        ErrCode int         `json:"errcode"`
        ErrMsg string       `json:"errmsg"`
    }
    var data *SessionRes
    if err := res.Scan(&data); err != nil {
        // do sth
    }
    
    // 获取session_key
    data.SessionKey
    
    // 获取OpenId
    data.OpenId
    
    // 注意类型，及json别名的设置要与官方返回的字段相同

```