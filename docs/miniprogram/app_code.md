### 获取小程序码
* 获取小程序码有两种，一种是小程序码，一种是二维码，其中小程序码中的get和二维码createQRCode是数量限制的，
业务适用不同，请注意：

#### 获取小程序码
* 此接口获取数量有限制，且获取的二维码永久有效，适用于需要小程序码数量较少的业务
```golang
    code := weapp.Get("path", g.Map{"width": 430})
    //返回的为 miniprogram/AppCode 
    //保存
    name, err := code.Save("save_path","file_name.png")
    //err 为struct类型 miniprogram/AppCodeError 可通过判断err是否为nil
    //判断保存状态，保存成功返回文件名称 string类型
```
#### 接口 GetUnlimit
* 此接口数量暂时没有限制，适用于临时业务或需要数量极多的业务场景
```golang
    code := weapp.GetUnlimit("scene", g.Map{"width": 430,"page":"pages/index"})
    //返回的为 miniprogram/AppCode , 具体第二个参数查看微信官方文档，有详细说明
    //保存
    name, err := code.Save("save_path","file_name.png")
    //err 为struct类型 miniprogram/AppCodeError 可通过判断err是否为nil
    //判断保存状态，保存成功返回文件名称 string类型
```
#### 接口 CreateQRCode
* 此接口的获取数量有限制
```golang
    code := weapp.CreateQrCode("path", 430)
    //返回的为 miniprogram/AppCode 
    //保存
    name, err := code.Save("save_path","file_name.png")
    //err 为struct类型 miniprogram/AppCodeError 可通过判断err是否为nil
    //判断保存状态，保存成功返回文件名称 string类型
```
