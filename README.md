# 网址

部署在我的服务器上了
`outsiders.top`
80端口

# 接口

## 注册

| 注册        |             |
| ----------- | ----------- |
| 方法        | `POST`      |
| URL         | `/api/user` |
| 是否需要jwt | 否          |
| 请求与响应  | `json`      |

### 请求json格式：

```json
{
    "name": "test",
    "phoneNum": "13651132812",
    "password":"123456",
    "sex": 0
}
```

### 响应json格式：

```json
{
    "msg": "phoneNum used",
    "status_code": 400
}
```

## 登录


| 登录        |              |
| ----------- | ------------ |
| 方法        | `POST`       |
| URL         | `/api/login` |
| 是否需要jwt | 否           |
| 请求与响应  | `json`       |

### 请求json格式：

```json
{
	"phoneNum":"13651132812",
    "password":"123456"
}
```

### 响应json格式：

```json
{
    "msg": "ok",
    "status_code": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZU51bSI6IjEzNjUxMTMyODEyIn0.yrkdzLLxv6rxauPAacchYs6Um_4mH115mq7-_IxnReY"
}
```
**token**即jwt，需前段储存。（没有过期时间）

## 获取用户信息
| 获取用户信息 |                          |
| ------------ | ------------------------ |
| 方法         | `GET`                    |
| URL          | `/api/user`              |
| 是否需要jwt  | 是                       |
| 请求与响应   | 请求body为空，响应`json` |

### 响应json格式：

```json
{
    "userInfo": {
        "name": "test",
        "phoneNum": "13651132812",
        "sex": 0
    },
    "userDiet": {
        "meals": [
            {
                "name": "first test",
                "date": 1234566789,
                "time": 2,
                "food": [
                    "aaa",
                    "bbb"
                ],
                "placeCategory": 0,
                "placeExact": "811"
            },
            {
                "name": "second test",
                "date": 134566789,
                "time": 1,
                "food": [
                    "aaa",
                    "bbb"
                ],
                "placeCategory": 0,
                "placeExact": "811"
            }
        ],
        "num": 2
    }
}
```


## 添加饮食记录
| 添加饮食记录 |                            |
| ------------ | -------------------------- |
| 方法         | `POST`                     |
| URL          | `/api/user/<手机号>/meals` |
| 是否需要jwt  | 是                         |
| 请求与响应   | `json`                     |

### 请求json格式：

```json
{
	"name":"third test",
	"date":134566789, //（这里时间传yyyy-mm-dd字符串即可）
	"time":1,
	"food":[
		"dddddd",
		"bbb"
	],
	"foodCategory":2,
	"placeExact":"abc"
}
```

### 响应json格式：

```json
{
    "userInfo": {
        "name": "test",
        "phoneNum": "13651132812",
        "sex": 0
    },
    "userDiet": {
        "meals": [
            {
                "name": "first test",
                "date": 1234566789,
                "time": 2,
                "food": [
                    "aaa",
                    "bbb"
                ],
                "placeCategory": 0,
                "placeExact": "811"
            },
            {
                "name": "second test",
                "date": 134566789,
                "time": 1,
                "food": [
                    "aaa",
                    "bbb"
                ],
                "placeCategory": 0,
                "placeExact": "811"
            },
            {
                "name": "third test",
                "date": 134566789,
                "time": 1,
                "food": [
                    "dddddd",
                    "bbb"
                ],
                "placeCategory": 0,
                "placeExact": "abc"
            }
        ],
        "num": 3
    }
}
```

## 获取饮食记录
| 获取饮食记录 |                            |
| ------------ | -------------------------- |
| 方法         | `GET`                      |
| URL          | `/api/user/<手机号>/meals` |
| 请求与响应   | 请求body为空，响应`json`   |

### 响应json格式：

```json
{
    "meals": [
        {
            "name": "first test",
            "date": 1234566789,
            "time": 2,
            "food": [
                "aaa",
                "bbb"
            ],
            "placeCategory": 0,
            "placeExact": "811"
        },
        {
            "name": "second test",
            "date": 134566789,
            "time": 1,
            "food": [
                "aaa",
                "bbb"
            ],
            "placeCategory": 0,
            "placeExact": "811"
        },
        {
            "name": "third test",
            "date": 134566789,
            "time": 1,
            "food": [
                "dddddd",
                "bbb"
            ],
            "placeCategory": 0,
            "placeExact": "abc"
        }
    ],
    "num": 3
}
```

## 上传图片

| 上传图片 |                            |
| ------------ | -------------------------- |
| 方法         | `POST`                      |
| URL          | `/api/photo` |
| 请求与响应   | 请求body为图片，响应`json`   |

### 响应json格式：

```json
{
    "res":[],
    "statue_code":200
}
```

## 推荐

| 推荐 |                            |
| ------------ | -------------------------- |
| 方法         | `GET`                      |
| URL          | `/api/user/<手机号>/recommend` |
| 请求与响应   | 请求body为空，响应`json`   |

### 响应json格式：

```json
{
    "res":[],
    "statue_code":200
}
```



# `JWT`的使用

在请求的header中加入`Authorization`头部，内容为`token <xxxxx>`，其中xxxxx为登录后获得的token（注意中间有个空格）
`JWT`没有过期时间

# 错误信息

有可能会返回错误，一般也是json格式

```json
{
    "msg":"xxxx",
    "status_code":404
}
```