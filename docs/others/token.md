### Token 说明

### 要求
- 缓存方式
- 过期自动刷新

### 实现方式

## 前言
- 许多接口请求时需要携带token，例如微信端的access-token，其他如抖音的client-token等
- 请求携带的方式不同， 如url参数，header，body等

## 实现

- 基于goframe的[gclient](https://goframe.org/display/gf/HTTPClient),让其在请求前注入token，有了更多的灵活性
- 同样，为实现token过期，及缓存问题，goframe的[gcache](https://goframe.org/pages/viewpage.action?pageId=1114679),同样提供了超便捷的方案。


具体实现可以阅读源码，结合goframe文档的使用方法，自行实现。