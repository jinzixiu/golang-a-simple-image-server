## 一个简单的 图片服务器服务

###1.编译
`go build imageServer.go`

###2.执行
`nohup ./imageServer  1> server.out 2> server.err`

执行会自动产生目录

- /static/tools/
- /static/dev/
- /static/live/
- /static/identify/

```go
配置代码在 imageServer.go 文件开头
const(
	adomin="localhost:9090"
	toolsdir="/static/tools/"
	devdir="/static/dev/"
	livedir="/static/live/"
	identify ="/static/identify/"
)
```

### 前端使用方法
````
var form = new FormData();
form.append("uploadfile", "2.png");
form.append("publickey", "公钥");
form.append("secretkey", "秘钥");

var settings = {
  "async": true,
  "crossDomain": true,
  "url": "http://*********/tools/upload",
  "method": "POST",
  "processData": false,
  "contentType": false,
  "mimeType": "multipart/form-data",
  "data": form
}

$.ajax(settings).done(function (response) {
  console.log(response);
});
````



### [ tools目录 ] 上传图片获取url的接口

`POST - *********/tools/upload`

|  参数名 |  参数类型 |  描述  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
|  uploadfile  |  file  |  #  |  #  |
|  publickey  |  string  |  #  |  公钥  |
|  secretkey  |  string  |  #  |  秘钥  |

##### 返回值
```json
{
  "code": 200,
  "data": "*********:8080/static/tools/5ee56109-2c07-47bc-9677-caffbce80a89.png",
  "msg": "success"
}
```
##### 备注
```
这里写备注
```
------------


### [ dev目录 ] 上传图片获取url的接口

`POST - *********/dev/upload`

|  参数名 |  参数类型 |  描述  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
|  uploadfile  |  file  |  #  |  #  |
|  publickey  |  string  |  #  |  公钥  |
|  secretkey  |  string  |  #  |  秘钥  |

##### 返回值
```json
{
  "code": 200,
  "data": "*********:8080/static/dev/5ee56109-2c07-47bc-9677-caffbce80a89.png",
  "msg": "success"
}
```
##### 备注
```
这里写备注
```
------------


### [ live目录 ] 上传图片获取url的接口

`POST - *********/live/upload`

|  参数名 |  参数类型 |  描述  | 备注  |
| ------------ | ------------ | ------------ | ------------ |
|  uploadfile  |  file  |  #  |  #  |
|  publickey  |  string  |  #  |  公钥  |
|  secretkey  |  string  |  #  |  秘钥  |

##### 返回值
```json
{
  "code": 200,
  "data": "*********:8080/static/live/5ee56109-2c07-47bc-9677-caffbce80a89.png",
  "msg": "success"
}
```
##### 备注
```
这里写备注
```
------------