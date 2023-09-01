## 全局公共参数
#### 全局Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
token | 233091959b615823d9d95ab916851b21 | 用户Token
#### 全局Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 全局Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 全局认证方式
```text
noauth
```
#### 全局预执行脚本
```javascript
暂无预执行脚本
```
#### 全局后执行脚本
```javascript
暂无后执行脚本
```
## /卡密充值
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:3200/api/use_carmi

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 233091959b615823d9d95ab916851b21 | String | 是 | 用户token
#### 请求Body参数
```javascript
{
    "carmiHandlers": "q1"
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
carmi | q1 | String | 是 | 卡密
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": [],
	"message": "充值成功"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
message | 充值成功 | String | 
## /获取code
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:3200/api/send_sms?source=youemail@qq.com

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
source | youemail@qq.com | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"message": "发送成功，请注意查收！",
	"data": {}
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 
message | 发送成功，请注意查收！ | String | 
data | - | Object | 
#### 错误响应示例
```javascript
{
	"code": -1,
	"message": "请使用邮箱登录！",
	"data": {}
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | -1 | Integer | 
message | 请使用邮箱登录！ | String | 
data | - | Object | 
## /登录/注册
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:3200/api/login

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数
```javascript
{
    "account":"youemail@qq.com",
    "code": "014198"
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
account | youemail@qq.com | String | 是 | 可以是邮箱或者手机号
code | 123456 | Integer | 是 | 验证码
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"user_info": {
			"id": "49454828484562944",
			"account": "you@qq.com",
			"nickname": "Chat用户",
			"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
			"role": "user ｜ administrator",
			"integral": 650,
			"vip_expire_time": "2023-06-24",
			"svip_expire_time": "2023-06-24",
			"password": "6c37",
			"ip": "127.0.0.1",
			"status": 1,
			"create_time": "2023-05-17 11:15:50",
			"update_time": "2023-05-28 15:56:22",
			"is_signin": 1
		},
		"token": "0342a173a0d99c24077f5eb66348b5ef"
	},
	"message": "登陆成功"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 返回数据
data.user_info | - | Object | 
data.user_info.id | 49454828484562944 | String | 
data.user_info.account | you@qq.com | String | 可以是邮箱或者手机号
data.user_info.nickname | Chat用户 | String | 
data.user_info.avatar | https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png | String | 
data.user_info.role | user ｜ administrator | String | 角色 // administrator 为管理员
data.user_info.integral | 650 | Integer | 积分
data.user_info.vip_expire_time | 2023-06-24 | String | 
data.user_info.svip_expire_time | 2023-06-24 | String | 
data.user_info.password | 6c37 | String | 
data.user_info.ip | 127.0.0.1 | String | 
data.user_info.status | 1 | Integer | 
data.user_info.create_time | 2023-05-17 11:15:50 | String | 
data.user_info.update_time | 2023-05-28 15:56:22 | String | 
data.user_info.is_signin | 1 | Integer | 当日是否签到
data.token | 0342a173a0d99c24077f5eb66348b5ef | String | 认证令牌
message | 登陆成功 | String | 
#### 错误响应示例
```javascript
{
	"code": -1,
	"data": [],
	"message": "请先发送验证码"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | -1 | Integer | 
data | - | Array | 
message | 请先发送验证码 | String | 
## /获取用户信息
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:3200/api/user/info

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | ae04a0328c868f188c198c43e06df3e2 | String | 是 | 用户Token
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"id": "49454828484562944",
		"account": "you@qq.com",
		"nickname": "Chat用户",
		"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
		"role": "user ｜ administrator",
		"integral": 650,
		"vip_expire_time": "2023-06-24",
		"svip_expire_time": "2023-06-24",
		"password": "6c37",
		"ip": "127.0.0.1",
		"status": 1,
		"create_time": "2023-05-17 11:15:50",
		"update_time": "2023-05-28 15:56:22",
		"is_signin": 1
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.id | 49454828484562944 | String | 
data.account | 201444307@qq.com | String | 可以是邮箱或者手机号
data.nickname | Chat用户 | String | 
data.avatar | https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png | String | 
data.role | administrator | String | 
data.integral | 650 | Integer | 积分
data.vip_expire_time | 2023-06-24 | String | 
data.svip_expire_time | 2023-06-24 | String | 
data.password | 6c34d65f8a75e9ac5ab455cdec2062f7 | String | 
data.ip | 127.0.0.1 | String | 
data.status | 1 | Integer | 
data.create_time | 2023-05-17 11:15:50 | String | 
data.update_time | 2023-05-28 15:56:22 | String | 
data.is_signin | 1 | Integer | 是否签到
message | - | String | 
#### 错误响应示例
```javascript
{
	"code": -1,
	"data": [],
	"message": "请登陆"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | -1 | Integer | 验证码
data | - | Array | 
message | 请登陆 | String | 
## /绘画
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:3200/api/images/generations

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | 用户Token
#### 请求Body参数
```javascript
{
  "prompt": "画一个中国龙",
  "n": 1,
  "size": "1024x1024",
  "response_format": "url"
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": [
		{
			"url": ""
		}
	],
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
data.url | - | String | 
message | - | String | 
#### 错误响应示例
```javascript
{
	"code": -1,
	"data": [],
	"message": ""
}
```
## /产品列表
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:3200/api/product?page_size=1

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 90f7882930465c5542eadae1dbc14ec8 | String | 是 | -
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
page_size | 1 | String | 是 | 单次获取数量
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"products": [
			{
				"id": "51616604789673984",
				"title": "1元366分",
				"price": 100,
				"original_price": 200,
				"value": 366,
				"badge": "尝鲜",
				"type": "integral",
				"level": 2,
				"status": 1,
				"create_time": "2023-05-23 10:25:57",
				"update_time": "2023-05-26 13:05:02"
			},
			{
				"id": 50420988902379520,
				"title": "10元15天",
				"price": 1000,
				"original_price": 2000,
				"value": 10,
				"badge": "会员",
				"type": "day",
				"level": 1,
				"status": 1,
				"create_time": "2023-05-20 03:15:01",
				"update_time": "2023-05-26 13:01:41"
			},
			{
				"id": 4342432,
				"title": "3元3天",
				"price": 300,
				"original_price": 500,
				"value": 3,
				"badge": "会员",
				"type": "day",
				"level": 1,
				"status": 1,
				"create_time": "2023-05-17 19:49:25",
				"update_time": "2023-05-26 13:01:40"
			},
			{
				"id": 345678,
				"title": "3元1000积分",
				"price": 300,
				"original_price": 500,
				"value": 1000,
				"badge": "特惠",
				"type": "integral",
				"level": 1,
				"status": 1,
				"create_time": "2023-05-17 17:07:53",
				"update_time": "2023-05-26 13:01:37"
			}
		],
		"pay_types": [
			"alipay",
			"wxpay",
			"qqpay"
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.products | - | Array | 
data.products.id | 51616604789673984 | String | 
data.products.title | 1元366分 | String | 
data.products.price | 100 | Integer | 
data.products.original_price | 200 | Integer | 
data.products.value | 366 | Integer | 
data.products.badge | 尝鲜 | String | 
data.products.type | integral | String | 
data.products.level | 2 | Integer | 
data.products.status | 1 | Integer | 
data.products.create_time | 2023-05-23 10:25:57 | String | 
data.products.update_time | 2023-05-26 13:05:02 | String | 
data.pay_types | alipay | Array | 
message | - | String | 
## /获取消费记录
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:3200/api/turnover?page=1&pageSIze=10

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 233091959b615823d9d95ab916851b21 | String | 是 | 用户Token
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
page | 1 | String | 是 | -
pageSIze | 10 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 33,
		"rows": [
			{
				"id": "50257850848514048",
				"user_id": "49454828484562944",
				"value": "10积分",
				"describe": "签到奖励",
				"create_time": "2023-05-19 16:26:46",
				"update_time": "2023-05-19 16:26:46"
			},
			{
				"id": "50254395463438336",
				"user_id": "49454828484562944",
				"value": "-1积分",
				"describe": "对话",
				"create_time": "2023-05-19 16:13:02",
				"update_time": "2023-05-19 16:13:02"
			},
			{
				"id": 50252779276472320,
				"user_id": "49454828484562944",
				"value": "-1积分",
				"describe": "对话",
				"create_time": "2023-05-19 16:06:37",
				"update_time": "2023-05-19 16:06:37"
			},
			{
				"id": 50252664373514240,
				"user_id": "49454828484562944",
				"value": "-1积分",
				"describe": "对话",
				"create_time": "2023-05-19 16:06:09",
				"update_time": "2023-05-19 16:06:09"
			},
			{
				"id": "50213819577798656",
				"user_id": "49454828484562944",
				"value": "-1积分",
				"describe": "对话",
				"create_time": "2023-05-19 13:31:48",
				"update_time": "2023-05-19 13:31:48"
			},
			{
				"id": 50212154862735360,
				"user_id": "49454828484562944",
				"value": "-2积分",
				"describe": "对话",
				"create_time": "2023-05-19 13:25:10",
				"update_time": "2023-05-19 13:25:10"
			},
			{
				"id": "50212125896871936",
				"user_id": "49454828484562944",
				"value": "-1积分",
				"describe": "对话",
				"create_time": "2023-05-19 13:25:04",
				"update_time": "2023-05-19 13:25:04"
			},
			{
				"id": "50208144046952448",
				"user_id": "49454828484562944",
				"value": "-1积分",
				"describe": "对话",
				"create_time": "2023-05-19 13:09:14",
				"update_time": "2023-05-19 13:09:14"
			},
			{
				"id": "50208098165460992",
				"user_id": "49454828484562944",
				"value": "-1积分",
				"describe": "对话",
				"create_time": "2023-05-19 13:09:03",
				"update_time": "2023-05-19 13:09:03"
			},
			{
				"id": 50207988325027840,
				"user_id": "49454828484562944",
				"value": "-2积分",
				"describe": "对话",
				"create_time": "2023-05-19 13:08:37",
				"update_time": "2023-05-19 13:08:37"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.count | 33 | Integer | 
data.rows | - | Array | 
data.rows.id | 50257850848514048 | String | 
data.rows.user_id | 49454828484562944 | String | 
data.rows.value | 10积分 | String | 
data.rows.describe | 签到奖励 | String | 
data.rows.create_time | 2023-05-19 16:26:46 | String | 
data.rows.update_time | 2023-05-19 16:26:46 | String | 
message | - | String | 
## /签到
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:3200/api/signin

#### 请求方式
> POST

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | 用户Token
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": [],
	"message": "签到成功 +3积分"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
message | 签到成功 +3积分 | String | 
## /对话
```text
暂无描述
```
#### 接口状态
> 已完成

#### 接口URL
> http://127.0.0.1:3200/api/chat/completions

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 478da26028ad39dd8e8297c43560b736 | String | 是 | 用户Token
#### 请求Body参数
```javascript
{
    "prompt": "你好",
    "parentMessageId": "e24c1448-16bf-482e-bff3-db1bb9e21b39",
    "options": {
        "model": "gpt-3.5-turbo",
        "temperature": 0,
        "presence_penalty": 0,
        "frequency_penalty": 0,
        "max_tokens": 2000
    }
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
prompt | 你好 | String | 是 | 话
parentMessageId | e24c1448-16bf-482e-bff3-db1bb9e21b39 | String | 是 | 会话ID
options | - | Object | 是 | 一些参数 比如温度 模型
options.model | gpt-3.5-turbo | String | 是 | 模型
options.temperature | 0 | Integer | 是 | 输出随机性
options.presence_penalty | 0 | Integer | 是 | 惩罚性质
options.frequency_penalty | 0 | Integer | 是 | 惩罚频率
options.max_tokens | 2000 | Integer | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"role": "assistant",
	"id": "chatcmpl-783vCoRLqyxviZJ001ILo6vyFpFM2",
	"parentMessageId": "1b1d43d6-56ce-448f-847d-f4df61951958",
	"content": "",
	"model": "gpt-4-0314",
	"dateTime": "2023-02-03 12:33:33",
	"segment": "'start' | 'text' | 'stop'"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
role | assistant | String | 角色
id | chatcmpl-783vCoRLqyxviZJ001ILo6vyFpFM2 | String | 
parentMessageId | 1b1d43d6-56ce-448f-847d-f4df61951958 | String | 会话ID
content | - | String | 内容
model | gpt-4-0314 | String | 模型
dateTime | 2023-02-03 12:33:33 | String | 
segment | 'start' | 'text' | 'stop' | String | 开始 中间 结束
## /获取用户当月签到记录
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/signin/list

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 2ab4e7a57550e844ce094a6dcb36780c | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /订单创建
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/pay/precreate

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 3672bddd96e31ce753acb38813819037 | String | 是 | -
#### 请求Body参数
```javascript
{
    "pay_type": "alipay",
    "product_id": 50420988902379520,
    "quantity": 1
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"order_id": "51064445698314240",
		"pay_url": "https://qr.alipay.com/bax02138sch9ehchnpge30c5"
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 
data | - | Object | 
data.order_id | 51064445698314240 | String | 
data.pay_url | https://qr.alipay.com/bax02138sch9ehchnpge30c5 | String | 
message | - | String | 
## /重置用户密码
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/user/password

#### 请求方式
> PUT

#### Content-Type
> form-data

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /配置文件
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/config

#### 请求方式
> GET

#### Content-Type
> form-data

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/卡密管理
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/卡密管理/卡密列表
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/carmi?page=1&page_size=10

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
page | 1 | String | 是 | -
page_size | 10 | String | 是 | 单次获取数量
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 60,
		"rows": [
			{
				"id": "50365654087237632",
				"user_id": null,
				"ip": null,
				"key": "ac8ac5639a9ea49ce048f5b5ae89aff8",
				"value": "10",
				"type": "integral",
				"end_time": "2023-05-19",
				"status": 2,
				"create_time": "2023-05-19 23:35:09",
				"update_time": "2023-05-20 00:00:02"
			},
			{
				"id": "50365654087241728",
				"user_id": null,
				"ip": null,
				"key": "d53305cdba64d75fdb93423947fc0155",
				"value": "10",
				"type": "integral",
				"end_time": "2023-05-19",
				"status": 2,
				"create_time": "2023-05-19 23:35:09",
				"update_time": "2023-05-20 00:00:02"
			},
			{
				"id": "50365654087245824",
				"user_id": null,
				"ip": null,
				"key": "cee676ed05acabc5b544ad3414d3673f",
				"value": "10",
				"type": "integral",
				"end_time": "2023-05-18",
				"status": 2,
				"create_time": "2023-05-19 23:35:09",
				"update_time": "2023-05-19 23:38:00"
			},
			{
				"id": 50365654087249920,
				"user_id": null,
				"ip": null,
				"key": "d0d60cfa3e09760b43c11ad3378420ca",
				"value": "10",
				"type": "integral",
				"end_time": "2023-05-18",
				"status": 2,
				"create_time": "2023-05-19 23:35:09",
				"update_time": "2023-05-19 23:38:00"
			},
			{
				"id": "50360647338164224",
				"user_id": null,
				"ip": null,
				"key": "c72d23328dfd6f812ff47ee50e8b560b",
				"value": "10",
				"type": "integral",
				"end_time": "2023-05-20",
				"status": 2,
				"create_time": "2023-05-19 23:15:14",
				"update_time": "2023-05-21 00:00:01"
			},
			{
				"id": "50360149126152192",
				"user_id": null,
				"ip": null,
				"key": "56101bcdd8b21603cba084d9705d86ec",
				"value": "10",
				"type": "integral",
				"end_time": "",
				"status": 0,
				"create_time": "2023-05-19 23:13:16",
				"update_time": "2023-05-19 23:13:16"
			},
			{
				"id": "50360149126217728",
				"user_id": null,
				"ip": null,
				"key": "22cf4d493bc7e520068002434376009f",
				"value": "10",
				"type": "integral",
				"end_time": "",
				"status": 0,
				"create_time": "2023-05-19 23:13:16",
				"update_time": "2023-05-19 23:13:16"
			},
			{
				"id": "50360149126156288",
				"user_id": null,
				"ip": null,
				"key": "7875e826ee182ce797cdbbf11dd3ae56",
				"value": "10",
				"type": "integral",
				"end_time": "",
				"status": 0,
				"create_time": "2023-05-19 23:13:16",
				"update_time": "2023-05-19 23:13:16"
			},
			{
				"id": "50360149126221824",
				"user_id": null,
				"ip": null,
				"key": "0e198132cdb8476a90f0f949aee672b6",
				"value": "10",
				"type": "integral",
				"end_time": "",
				"status": 0,
				"create_time": "2023-05-19 23:13:16",
				"update_time": "2023-05-19 23:13:16"
			},
			{
				"id": "50360149126160384",
				"user_id": null,
				"ip": null,
				"key": "881e96e45ee6d5e78755aeb1c5707c9f",
				"value": "10",
				"type": "integral",
				"end_time": "",
				"status": 0,
				"create_time": "2023-05-19 23:13:16",
				"update_time": "2023-05-19 23:13:16"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.count | 60 | Integer | 
data.rows | - | Array | 
data.rows.id | 50365654087237632 | String | 
data.rows.user_id | null | Null | 
data.rows.ip | null | Null | 
data.rows.key | ac8ac5639a9ea49ce048f5b5ae89aff8 | String | 
data.rows.value | 10 | String | 
data.rows.type | integral | String | 
data.rows.end_time | 2023-05-19 | String | 
data.rows.status | 2 | Integer | 
data.rows.create_time | 2023-05-19 23:35:09 | String | 
data.rows.update_time | 2023-05-20 00:00:02 | String | 
message | - | String | 
## /Admin/卡密管理/删除卡密
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/carmi/50745921733918720

#### 请求方式
> DELETE

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | 用户Token
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": 1,
	"message": "删除成功"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | 1 | Integer | 
message | 删除成功 | String | 
## /Admin/卡密管理/生成卡密
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/carmi

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Body参数
```javascript
{
    "type": "integral",
    "end_time": "",
    "quantity": 1,
    "reward": 10
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
type | integral | String | 是 | 类型 积分 和 天数
end_time | - | String | 是 | 卡密截止日期
quantity | 1 | Integer | 是 | 生成数量
reward | 10 | Integer | 是 | 奖励数量
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": [
		{
			"id": "50746199967268864",
			"key": "8fb9e1af5adedacd914c8a4f778294e4",
			"type": "integral",
			"end_time": "",
			"value": 10,
			"status": 0
		}
	],
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
data.id | 50746199967268864 | String | 
data.key | 8fb9e1af5adedacd914c8a4f778294e4 | String | 
data.type | integral | String | 类型 积分 和 天数
data.end_time | - | String | 卡密截止日期
data.value | 10 | Integer | 
data.status | 0 | Integer | 
message | - | String | 
#### 错误响应示例
```javascript
{
	"code": -1,
	"data": [],
	"message": "拒绝访问"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | -1 | Integer | 验证码
data | - | Array | 
message | 拒绝访问 | String | 
## /Admin/卡密管理/检查卡密
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/carmi/check

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | gsdgsdgsdgsdgsdgsdgsdg | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/用户管理
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/用户管理/删除用户
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/user/50260043739697152

#### 请求方式
> DELETE

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": 1,
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | 1 | Integer | 
message | - | String | 
## /Admin/用户管理/用户列表
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/user?page=1&page_size=10

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
page | 1 | String | 是 | -
page_size | 10 | String | 是 | 单次获取数量
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 33,
		"rows": [
			{
				"id": "50607418651971584",
				"account": "1159217956@qq.com",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": 10,
				"subscribe": "2023-05-10",
				"ip": "140.83.51.187",
				"status": 1,
				"create_time": "2023-05-20 15:35:49",
				"update_time": "2023-05-21 00:00:06"
			},
			{
				"id": "50535861028130816",
				"account": "1107590384@qq.com",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": 10,
				"subscribe": "2023-05-19",
				"ip": "103.167.134.150",
				"status": 1,
				"create_time": "2023-05-20 10:51:28",
				"update_time": "2023-05-20 18:56:35"
			},
			{
				"id": "50515411749310464",
				"account": "1991336263@qq.com",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": 9,
				"subscribe": "2023-05-19",
				"ip": "180.120.86.12",
				"status": 1,
				"create_time": "2023-05-20 09:30:12",
				"update_time": "2023-05-20 18:56:34"
			},
			{
				"id": "50380307794235392",
				"account": "huateng618@yeah.net",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": 0,
				"subscribe": "2023-05-19",
				"ip": "223.104.66.220",
				"status": 1,
				"create_time": "2023-05-20 00:33:21",
				"update_time": "2023-05-20 18:56:34"
			},
			{
				"id": "50357285670621184",
				"account": "121044767@qq.com",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": 0,
				"subscribe": "2023-05-18",
				"ip": "120.85.100.214",
				"status": 1,
				"create_time": "2023-05-19 23:01:52",
				"update_time": "2023-05-20 18:56:34"
			},
			{
				"id": "50339754931130368",
				"account": "2476173156@qq.com",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": 10,
				"subscribe": "2023-05-18",
				"ip": "171.108.138.60",
				"status": 1,
				"create_time": "2023-05-19 21:52:13",
				"update_time": "2023-05-20 18:56:33"
			},
			{
				"id": "50338559730651136",
				"account": "1146544859@qq.com",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": 0,
				"subscribe": "2023-05-18",
				"ip": "58.48.27.46",
				"status": 1,
				"create_time": "2023-05-19 21:47:28",
				"update_time": "2023-05-20 18:56:33"
			},
			{
				"id": "50321886659219456",
				"account": "2948847465@qq.com",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": -5,
				"subscribe": "2023-05-18",
				"ip": "223.104.111.73",
				"status": 1,
				"create_time": "2023-05-19 20:41:12",
				"update_time": "2023-05-20 18:56:33"
			},
			{
				"id": "50316933005840384",
				"account": "fanshuai8@gmail.com",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": 9,
				"subscribe": "2023-05-18",
				"ip": "60.176.185.14",
				"status": 1,
				"create_time": "2023-05-19 20:21:31",
				"update_time": "2023-05-20 18:56:32"
			},
			{
				"id": "50291314268311552",
				"account": "wanghzyzhh@gmail.com",
				"nickname": "Chat用户",
				"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
				"role": "user",
				"integral": 10,
				"subscribe": "2023-05-18",
				"ip": "113.12.232.137",
				"status": 1,
				"create_time": "2023-05-19 18:39:43",
				"update_time": "2023-05-20 18:56:32"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.count | 33 | Integer | 
data.rows | - | Array | 
data.rows.id | 50607418651971584 | String | 
data.rows.account | 1159217956@qq.com | String | 可以是邮箱或者手机号
data.rows.nickname | Chat用户 | String | 
data.rows.avatar | https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png | String | 
data.rows.role | user | String | 角色
data.rows.integral | 10 | Integer | 积分
data.rows.subscribe | 2023-05-10 | String | 
data.rows.ip | 140.83.51.187 | String | 
data.rows.status | 1 | Integer | 
data.rows.create_time | 2023-05-20 15:35:49 | String | 
data.rows.update_time | 2023-05-21 00:00:06 | String | 
message | - | String | 
## /Admin/用户管理/修改用户信息
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/user

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Body参数
```javascript
{
    "account": "1159217956@qq.com",
    "role": "user",
    "status": 1,
    "nickname": "Chat用户",
    "avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
    "integral": 10,
    "subscribe": "2023-05-10",
    "id": "50607418651971584"
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": [
		{
			"id": "50607418651971584",
			"account": "1159217956@qq.com",
			"avatar": "https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png",
			"integral": 10,
			"nickname": "Chat用户",
			"role": "user",
			"subscribe": "2023-05-10"
		},
		false
	],
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
data.id | 50607418651971584 | String | 
data.account | 1159217956@qq.com | String | 可以是邮箱或者手机号
data.avatar | https://u1.dl0.cn/icon/1682426702646avatarf3db669b024fad66-1930929abe2847093.png | String | 
data.integral | 10 | Integer | 积分
data.nickname | Chat用户 | String | 
data.role | user | String | 角色
data.subscribe | 2023-05-10 | String | 订阅
message | - | String | 
## /Admin/商品管理
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/商品管理/商品列表
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/products?page=1&page_size=10

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
page | 1 | String | 是 | -
page_size | 10 | String | 是 | 单次获取数量
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 3,
		"rows": [
			{
				"id": 50420988902379520,
				"title": "测试001",
				"price": 100,
				"original_price": 0,
				"integral": 0,
				"badge": "测试",
				"day": 7,
				"status": 1,
				"create_time": "2023-05-20 03:15:01",
				"update_time": "2023-05-20 03:20:53"
			},
			{
				"id": 4342432,
				"title": "3天1",
				"price": 300,
				"original_price": 500,
				"integral": null,
				"badge": "会员",
				"day": 3,
				"status": 1,
				"create_time": "2023-05-17 19:49:25",
				"update_time": "2023-05-20 03:13:58"
			},
			{
				"id": 345678,
				"title": "1000积分",
				"price": 300,
				"original_price": 500,
				"integral": 1000,
				"badge": "特惠",
				"day": null,
				"status": 1,
				"create_time": "2023-05-17 17:07:53",
				"update_time": "2023-05-17 17:09:07"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.count | 3 | Integer | 
data.rows | - | Array | 
data.rows.id | 50420988902379520 | Integer | 
data.rows.title | 测试001 | String | 标题
data.rows.price | 100 | Integer | 价格（分）
data.rows.original_price | 0 | Integer | 原价（分）
data.rows.integral | 0 | Integer | 积分
data.rows.badge | 测试 | String | 角标
data.rows.day | 7 | Integer | 天数
data.rows.status | 1 | Integer | 
data.rows.create_time | 2023-05-20 03:15:01 | String | 
data.rows.update_time | 2023-05-20 03:20:53 | String | 
message | - | String | 
## /Admin/商品管理/新增商品
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/products

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Body参数
```javascript
{
    "title": "测试",
    "badge": "会员",
    "status": 1,
    "price": 1000,
    "original_price": 1223,
    "integral": 1000,
    "day": 0
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
title | 测试 | String | 是 | -
badge | 会员 | String | 是 | -
status | 1 | Integer | 是 | -
price | 1000 | Integer | 是 | -
original_price | 1223 | Integer | 是 | -
integral | 1000 | Integer | 是 | -
day | 0 | Integer | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"id": "50747888841527296",
		"title": "测试",
		"price": 1000,
		"original_price": 1223,
		"integral": 1000,
		"badge": "会员",
		"day": 0,
		"status": 1
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.id | 50747888841527296 | String | 
data.title | 测试 | String | 
data.price | 1000 | Integer | 
data.original_price | 1223 | Integer | 
data.integral | 1000 | Integer | 积分
data.badge | 会员 | String | 
data.day | 0 | Integer | 
data.status | 1 | Integer | 
message | - | String | 
## /Admin/商品管理/修改商品
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/products

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Body参数
```javascript
{
    "title": "测试",
    "badge": "会员",
    "status": 1,
    "price": 1000,
    "original_price": 12232,
    "integral": 1000,
    "day": 0,
    "id": "50747888841527296"
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
title | 测试 | String | 是 | -
badge | 会员 | String | 是 | -
status | 1 | Integer | 是 | -
price | 1000 | Integer | 是 | -
original_price | 12232 | Integer | 是 | -
integral | 1000 | Integer | 是 | -
day | 0 | Integer | 是 | -
id | 50747888841527296 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": [
		{
			"id": "50747888841527296",
			"title": "测试",
			"price": 1000,
			"original_price": 12232,
			"integral": 1000,
			"badge": "会员",
			"day": 0,
			"status": 1
		},
		false
	],
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
data.id | 50747888841527296 | String | 
data.title | 测试 | String | 
data.price | 1000 | Integer | 
data.original_price | 12232 | Integer | 
data.integral | 1000 | Integer | 积分
data.badge | 会员 | String | 
data.day | 0 | Integer | 
data.status | 1 | Integer | 
message | - | String | 
## /Admin/商品管理/删除商品
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/products/50747888841527296

#### 请求方式
> DELETE

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Body参数
```javascript

```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
title | 测试 | String | 是 | -
badge | 会员 | String | 是 | -
status | 1 | Integer | 是 | -
price | 1000 | Integer | 是 | -
original_price | 12232 | Integer | 是 | -
integral | 1000 | Integer | 是 | -
day | 0 | Integer | 是 | -
id | 50747888841527296 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": 1,
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | 1 | Integer | 
message | - | String | 
## /Admin/会话管理
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/会话管理/对话列表
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/messages?page=1&page_size=10

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
page | 1 | String | 是 | -
page_size | 10 | String | 是 | 单次获取数量
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 150,
		"rows": [
			{
				"id": "50743051148070912",
				"content": "你好",
				"user_id": "49454828484562944",
				"role": "user",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "e24c1448-16bf-482e-bff3-db1bb9e21b39",
				"status": 1,
				"create_time": "2023-05-21 00:34:47",
				"update_time": "2023-05-21 00:34:47"
			},
			{
				"id": "50743051148075008",
				"content": "你好！有什么我可以帮助你的吗？",
				"user_id": "49454828484562944",
				"role": "assistant",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "e24c1448-16bf-482e-bff3-db1bb9e21b39",
				"status": 1,
				"create_time": "2023-05-21 00:34:47",
				"update_time": "2023-05-21 00:34:47"
			},
			{
				"id": "50742652101988352",
				"content": "你好",
				"user_id": "49454828484562944",
				"role": "user",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "e24c1448-16bf-482e-bff3-db1bb9e21b39",
				"status": 1,
				"create_time": "2023-05-21 00:33:11",
				"update_time": "2023-05-21 00:33:11"
			},
			{
				"id": "50742652101992448",
				"content": "你好！有什么我可以帮助你的吗？",
				"user_id": "49454828484562944",
				"role": "assistant",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "e24c1448-16bf-482e-bff3-db1bb9e21b39",
				"status": 1,
				"create_time": "2023-05-21 00:33:11",
				"update_time": "2023-05-21 00:33:11"
			},
			{
				"id": "50720371879448576",
				"content": "python 实现冒泡排序",
				"user_id": "50714609354543104",
				"role": "user",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "5010d6c2-c70d-4c2b-916b-9ea6f87fb67e",
				"status": 1,
				"create_time": "2023-05-20 23:04:49",
				"update_time": "2023-05-20 23:04:49"
			},
			{
				"id": "50720371879452672",
				"content": "以下是 Python 实现冒泡排序的代码：\n\npython\ndef bubble_sort(arr):\n    n = len(arr)\n    for i in range(n):\n        for j in range(0, n-i-1):\n            if arr[j] > arr[j+1]:\n                arr[j], arr[j+1] = arr[j+1], arr[j]\n\narr = [64, 34, 25, 12, 22, 11, 90]\nbubble_sort(arr)\nprint(\"排序后的数组：\")\nfor i in range(len(arr)):\n    print(\"%d\" % arr[i])\n\n\n输出结果：\n\n\n排序后的数组：\n11\n12\n22\n25\n34\n64\n90\n\n\n在这个实现中，我们使用了两个嵌套的循环来遍历数组。外层循环控制排序的轮数，内层循环控制每轮排序中的比较和交换操作。如果相邻的两个元素顺序不对，就交换它们的位置。这样，每轮排序都会把当前未排序部分的最大值移动到数组的末尾。最终，整个数组就被排序了。",
				"user_id": "50714609354543104",
				"role": "assistant",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "5010d6c2-c70d-4c2b-916b-9ea6f87fb67e",
				"status": 1,
				"create_time": "2023-05-20 23:04:49",
				"update_time": "2023-05-20 23:04:49"
			},
			{
				"id": "50720286881878016",
				"content": "给出傅里叶变换公式，用markdown渲染",
				"user_id": "50714609354543104",
				"role": "user",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "5010d6c2-c70d-4c2b-916b-9ea6f87fb67e",
				"status": 1,
				"create_time": "2023-05-20 23:04:21",
				"update_time": "2023-05-20 23:04:21"
			},
			{
				"id": "50720286881882112",
				"content": "傅里叶变换公式：\n\n$$\nF(\\omega) = \\int_{-\\infty}^{\\infty} f(t) e^{-i\\omega t} dt\n$$\n\n其中，$f(t)$ 是时间域函数，$F(\\omega)$ 是频率域函数，$i$ 是虚数单位，$\\omega$ 是角频率。",
				"user_id": "50714609354543104",
				"role": "assistant",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "5010d6c2-c70d-4c2b-916b-9ea6f87fb67e",
				"status": 1,
				"create_time": "2023-05-20 23:04:21",
				"update_time": "2023-05-20 23:04:21"
			},
			{
				"id": "50580592315600896",
				"content": "扫雷代码",
				"user_id": null,
				"role": "user",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "cf7f53b7-9dae-404c-a1db-69e9c18cf0c7",
				"status": 1,
				"create_time": "2023-05-20 13:49:36",
				"update_time": "2023-05-20 13:49:36"
			},
			{
				"id": "50580592315604992",
				"content": "抱歉，我是一个语言模型AI，我不能编写代码。但是，以下是一个简单的扫雷游戏的伪代码，供您参考：\n\n1. 初始化游戏界面，包括雷区大小、雷数等参数。\n2. 随机生成雷区，将雷随机分布在游戏界面中。\n3. 根据雷区中每个格子周围的雷数，标记出每个格子的数字。\n4. 玩家点击某个格子，如果该格子是雷，则游戏结束；如果该格子不是雷，则显示该格子的数字。\n5. 如果玩家点击的格子周围没有雷，则自动展开周围的格子，直到周围有雷为止。\n6. 玩家标记某个格子为雷，如果标记正确，则游戏继续；如果标记错误，则游戏结束。\n7. 当所有非雷格子都被揭开时，游戏胜利。\n\n以上是一个简单的扫雷游戏的伪代码，具体实现可能会有所不同。",
				"user_id": null,
				"role": "assistant",
				"frequency_penalty": 0,
				"max_tokens": 2000,
				"model": "gpt-3.5-turbo",
				"presence_penalty": 0,
				"temperature": 0,
				"parent_message_id": "cf7f53b7-9dae-404c-a1db-69e9c18cf0c7",
				"status": 1,
				"create_time": "2023-05-20 13:49:36",
				"update_time": "2023-05-20 13:49:36"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.count | 150 | Integer | 
data.rows | - | Array | 
data.rows.id | 50743051148070912 | String | 
data.rows.content | 你好 | String | 内容
data.rows.user_id | 49454828484562944 | String | 
data.rows.role | user | String | 角色
data.rows.frequency_penalty | 0 | Integer | 
data.rows.max_tokens | 2000 | Integer | 
data.rows.model | gpt-3.5-turbo | String | 模型
data.rows.presence_penalty | 0 | Integer | 
data.rows.temperature | 0 | Integer | 
data.rows.parent_message_id | e24c1448-16bf-482e-bff3-db1bb9e21b39 | String | 
data.rows.status | 1 | Integer | 
data.rows.create_time | 2023-05-21 00:34:47 | String | 
data.rows.update_time | 2023-05-21 00:34:47 | String | 
message | - | String | 
## /Admin/签到记录
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/签到记录/签到列表
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/signin?page=1&page_size=10

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
page | 1 | String | 是 | -
page_size | 10 | String | 是 | 单次获取数量
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 23,
		"rows": [
			{
				"id": "50742286476120064",
				"user_id": "49454828484562944",
				"ip": "127.0.0.1",
				"status": 1,
				"create_time": "2023-05-21 00:31:44",
				"update_time": "2023-05-21 00:31:44"
			},
			{
				"id": 50714662366351360,
				"user_id": "50714609354543104",
				"ip": "::1",
				"status": 1,
				"create_time": "2023-05-20 22:41:58",
				"update_time": "2023-05-20 22:41:58"
			},
			{
				"id": "50607633421307904",
				"user_id": "50607418651971584",
				"ip": "140.83.51.187",
				"status": 1,
				"create_time": "2023-05-20 15:36:40",
				"update_time": "2023-05-20 15:36:40"
			},
			{
				"id": "50536372896796672",
				"user_id": "50535861028130816",
				"ip": "103.167.134.150",
				"status": 1,
				"create_time": "2023-05-20 10:53:30",
				"update_time": "2023-05-20 10:53:30"
			},
			{
				"id": "50515538610229248",
				"user_id": "50515411749310464",
				"ip": "180.120.86.12",
				"status": 1,
				"create_time": "2023-05-20 09:30:43",
				"update_time": "2023-05-20 09:30:43"
			},
			{
				"id": "50357371406389248",
				"user_id": "50357285670621184",
				"ip": "120.85.100.214",
				"status": 1,
				"create_time": "2023-05-19 23:02:13",
				"update_time": "2023-05-19 23:02:13"
			},
			{
				"id": "50339814414749696",
				"user_id": "50339754931130368",
				"ip": "171.108.138.60",
				"status": 1,
				"create_time": "2023-05-19 21:52:27",
				"update_time": "2023-05-19 21:52:27"
			},
			{
				"id": "50322737821913088",
				"user_id": "50321886659219456",
				"ip": "223.104.111.73",
				"status": 1,
				"create_time": "2023-05-19 20:44:36",
				"update_time": "2023-05-19 20:44:36"
			},
			{
				"id": "50317029940400128",
				"user_id": "50316933005840384",
				"ip": "60.176.185.14",
				"status": 1,
				"create_time": "2023-05-19 20:21:55",
				"update_time": "2023-05-19 20:21:55"
			},
			{
				"id": "50291381356204032",
				"user_id": "50291314268311552",
				"ip": "113.12.232.137",
				"status": 1,
				"create_time": "2023-05-19 18:40:00",
				"update_time": "2023-05-19 18:40:00"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.count | 23 | Integer | 
data.rows | - | Array | 
data.rows.id | 50742286476120064 | String | 
data.rows.user_id | 49454828484562944 | String | 
data.rows.ip | 127.0.0.1 | String | 
data.rows.status | 1 | Integer | 
data.rows.create_time | 2023-05-21 00:31:44 | String | 
data.rows.update_time | 2023-05-21 00:31:44 | String | 
message | - | String | 
## /Admin/消费记录
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/消费记录/用户消费记录
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/turnover?page=1&page_size=10

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
page | 1 | String | 是 | -
page_size | 10 | String | 是 | 单次获取数量
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 112,
		"rows": [
			{
				"id": 50743053190696960,
				"user_id": "49454828484562944",
				"value": "-1积分",
				"describe": "对话",
				"create_time": "2023-05-21 00:34:47",
				"update_time": "2023-05-21 00:34:47"
			},
			{
				"id": 50742654165585920,
				"user_id": "49454828484562944",
				"value": "-1积分",
				"describe": "对话",
				"create_time": "2023-05-21 00:33:12",
				"update_time": "2023-05-21 00:33:12"
			},
			{
				"id": "50742286476120064",
				"user_id": "49454828484562944",
				"value": "3积分",
				"describe": "签到奖励",
				"create_time": "2023-05-21 00:31:45",
				"update_time": "2023-05-21 00:31:45"
			},
			{
				"id": "50720409754013696",
				"user_id": "50714609354543104",
				"value": "-7积分",
				"describe": "对话",
				"create_time": "2023-05-20 23:04:49",
				"update_time": "2023-05-20 23:04:49"
			},
			{
				"id": "50720297443135488",
				"user_id": "50714609354543104",
				"value": "-3积分",
				"describe": "对话",
				"create_time": "2023-05-20 23:04:22",
				"update_time": "2023-05-20 23:04:22"
			},
			{
				"id": 50714662366351360,
				"user_id": "50714609354543104",
				"value": "3积分",
				"describe": "签到奖励",
				"create_time": "2023-05-20 22:41:58",
				"update_time": "2023-05-20 22:41:58"
			},
			{
				"id": 50714611699159040,
				"user_id": "50714609354543104",
				"value": "10积分",
				"describe": "注册奖励",
				"create_time": "2023-05-20 22:41:46",
				"update_time": "2023-05-20 22:41:46"
			},
			{
				"id": "50607633421307904",
				"user_id": "50607418651971584",
				"value": "10积分",
				"describe": "签到奖励",
				"create_time": "2023-05-20 15:36:41",
				"update_time": "2023-05-20 15:36:41"
			},
			{
				"id": "50580688050589696",
				"user_id": "50321886659219456",
				"value": "-7积分",
				"describe": "对话",
				"create_time": "2023-05-20 13:49:36",
				"update_time": "2023-05-20 13:49:36"
			},
			{
				"id": "50580559574863872",
				"user_id": "50321886659219456",
				"value": "-5积分",
				"describe": "对话",
				"create_time": "2023-05-20 13:49:06",
				"update_time": "2023-05-20 13:49:06"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.count | 112 | Integer | 
data.rows | - | Array | 
data.rows.id | 50743053190696960 | Integer | 
data.rows.user_id | 49454828484562944 | String | 
data.rows.value | -1积分 | String | 
data.rows.describe | 对话 | String | 
data.rows.create_time | 2023-05-21 00:34:47 | String | 
data.rows.update_time | 2023-05-21 00:34:47 | String | 
message | - | String | 
## /Admin/消费记录/删除消费记录
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/turnover/50743053190696960

#### 请求方式
> DELETE

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": 1,
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | 1 | Integer | 
message | - | String | 
#### 错误响应示例
```javascript
{
	"code": -1,
	"data": [],
	"message": "拒绝访问"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | -1 | Integer | 验证码
data | - | Array | 
message | 拒绝访问 | String | 
## /Admin/消费记录/修改消费记录
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/turnover

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | -
#### 请求Body参数
```javascript
{
    "id": 50743053190696960,
    "user_id": "49454828484562944",
    "value": "-1积分",
    "describe": "对话",
    "create_time": "2023-05-21 00:34:47",
    "update_time": "2023-05-21 00:34:47"
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
    "code": 0,
    "data": [
        {
            "id": 50743053190696960,
            "value": "-1积分",
            "describe": "对话",
            "user_id": "49454828484562944"
        },
        true
    ],
    "message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
data.id | 50743053190696960 | Integer | 
data.value | -1积分 | String | 
data.describe | 对话 | String | 
data.user_id | 49454828484562944 | String | 
message | - | String | 
#### 错误响应示例
```javascript
{
	"code": -1,
	"data": [],
	"message": "拒绝访问"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | -1 | Integer | 验证码
data | - | Array | 
message | 拒绝访问 | String | 
## /Admin/Token管理
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/Token管理/检查Token
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/token/check

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | eaec72ab8082e1afa15520573b7faa5b | String | 是 | -
#### 请求Body参数
```javascript
{
    "host":"https://proxy.proxy.com",
    "key": "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"
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
host | https://caloi.top/openai | String | 是 | -
key | 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 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"status": 0,
		"hard_limit_usd": 5,
		"total_usage": 0.23
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.status | 0 | Integer | 
data.hard_limit_usd | 5 | Integer | 
data.total_usage | 0.23 | Number | 
message | - | String | 
#### 错误响应示例
```javascript
{
	"code": -1,
	"data": [],
	"message": "拒绝访问"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | -1 | Integer | 验证码
data | - | Array | 
message | 拒绝访问 | String | 
## /Admin/Token管理/Token列表
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/token?page=1&page_size=20

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | 用户Token
#### 请求Query参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
page | 1 | String | 是 | -
page_size | 20 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 2,
		"rows": [
			{
				"id": "50543019736174592",
				"key": "sk-XSNNkBQ4P8zT3BlbkFJX3zgjESYYLxVxeC6rFCe",
				"host": "https://api.openai-proxy.com",
				"remarks": "代理",
				"usage": 0.42,
				"limit": 5.00004,
				"status": 1,
				"create_time": "2023-05-20 11:19:56",
				"update_time": "2023-05-20 15:04:43"
			},
			{
				"id": 3456787654,
				"key": "fk192767-83uXEbVjcR3Go",
				"host": "https://opi.ad.net",
				"remarks": "A",
				"usage": 0,
				"limit": 0,
				"status": 0,
				"create_time": "2023-05-19 15:56:14",
				"update_time": "2023-05-20 15:04:44"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.count | 2 | Integer | 
data.rows | - | Array | 
data.rows.id | 50543019736174592 | String | 
data.rows.key | sk-XSNNkBQ4P8zT3BlbkFJX3zgjESYYLxVxeC6rFCe | String | 
data.rows.host | https://api.openai-proxy.com | String | 
data.rows.remarks | 代理 | String | 
data.rows.usage | 0.42 | Number | 
data.rows.limit | 5.00004 | Number | 
data.rows.status | 1 | Integer | 
data.rows.create_time | 2023-05-20 11:19:56 | String | 
data.rows.update_time | 2023-05-20 15:04:43 | String | 
message | - | String | 
## /Admin/Token管理/新增Token
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/token

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | 用户Token
#### 请求Body参数
```javascript
{
    "host": "https://caloi.top/openai",
    "key": "XSNNkBQ4PRrfHy66H78zT3BlbkFJX3zgjESYYLxVxeC6rFCe",
    "status": 1
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"id": "50750490736070656",
		"key": "XSNNkBQ4PRrfHy66H78zT3BlbkFJX3zgjESYYLxVxeC6rFCe",
		"host": "https://caloi.top/openai",
		"status": 1
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Object | 
data.id | 50750490736070656 | String | 
data.key | XSNNkBQ4PRrfHy66H78zT3BlbkFJX3zgjESYYLxVxeC6rFCe | String | 
data.host | https://caloi.top/openai | String | 
data.status | 1 | Integer | 
message | - | String | 
## /Admin/Token管理/修改Token
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/token

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Body参数
```javascript
{
    "host": "https://caloi.top/openai",
    "key": "XSNNkBQ4PRrfHy66H78zT3BlbkFJX3zgjESYYLxVxeC6rFCe",
    "status": 0,
    "id": "50750490736070656"
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
host | https://caloi.top/openai | String | 是 | -
key | XSNNkBQ4PRrfHy66H78zT3BlbkFJX3zgjESYYLxVxeC6rFCe | String | 是 | -
status | 0 | Integer | 是 | -
id | 50750490736070656 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": [
		{
			"id": "50750490736070656",
			"key": "XSNNkBQ4PRrfHy66H78zT3BlbkFJX3zgjESYYLxVxeC6rFCe",
			"host": "https://caloi.top/openai",
			"status": 0
		},
		false
	],
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
data.id | 50750490736070656 | String | 
data.key | XSNNkBQ4PRrfHy66H78zT3BlbkFJX3zgjESYYLxVxeC6rFCe | String | 
data.host | https://caloi.top/openai | String | 
data.status | 0 | Integer | 
message | - | String | 
## /Admin/Token管理/删除Token
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/token/50750490736070656

#### 请求方式
> DELETE

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | 用户Token
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": 1,
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | 1 | Integer | 
message | - | String | 
## /Admin/系统配置
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/系统配置/获取配置
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/config

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | 用户Token
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": [
		{
			"id": 1,
			"name": "signin_reward",
			"value": "3",
			"remarks": "签到奖励",
			"create_time": "2023-05-19 16:21:12",
			"update_time": "2023-05-20 17:27:07"
		},
		{
			"id": 2,
			"name": "register_reward",
			"value": "10",
			"remarks": "注册奖励",
			"create_time": "2023-05-19 16:21:49",
			"update_time": "2023-05-20 17:27:07"
		}
	],
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
data.id | 1 | Integer | 
data.name | signin_reward | String | 
data.value | 3 | String | 
data.remarks | 签到奖励 | String | 
data.create_time | 2023-05-19 16:21:12 | String | 
data.update_time | 2023-05-20 17:27:07 | String | 
message | - | String | 
## /Admin/系统配置/修改配置
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/config

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | b6404772785343665a52818f226ba276 | String | 是 | 用户Token
#### 请求Body参数
```javascript
{
    "register_reward": "10",
    "signin_reward": "3"
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
register_reward | 10 | String | 是 | -
signin_reward | 3 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": [],
	"message": "修改成功"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 验证码
data | - | Array | 
message | 修改成功 | String | 
## /Admin/支付配置
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/支付配置/获取支付渠道列表
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/payment

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 93fbbd25e371c2d9d275b1e5e3c6a5cc | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 2,
		"rows": [
			{
				"id": "51294279280103424",
				"name": "易支付",
				"channel": "yipay",
				"types": "alipay,wxpay",
				"params": "{\"pid\":\"1\",\"key\":\"5HIY5yiCdlDDlYFI\",\"api\":\"https://you.pay.com\",\"return_url\":\"\"}",
				"status": 1,
				"create_time": "2023-05-22 13:05:09",
				"update_time": "2023-05-22 13:16:48"
			},
			{
				"id": 24234242,
				"name": "支付宝-当面付",
				"channel": "alipay",
				"types": "alipay",
				"params": "{\"appId\":\"20122\",\"keyType\":\"PKCS8\",\"privateKey\":\"MIIEvQjR+ihDU=\",\"alipayPublicKey\":\"MII4MKxjjaOeAOQIDAQAB\"}",
				"status": 1,
				"create_time": "2023-05-21 21:04:51",
				"update_time": "2023-05-21 21:04:51"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 
data | - | Object | 
data.count | 2 | Integer | 
data.rows | - | Array | 
data.rows.id | 51294279280103424 | String | 
data.rows.name | 易支付 | String | 
data.rows.channel | yipay | String | 
data.rows.types | alipay,wxpay | String | 
data.rows.params | {"pid":"1","key":"5HIY5yiCdlDDlYFI","api":"https://you.pay.com","return_url":""} | String | 
data.rows.status | 1 | Integer | 
data.rows.create_time | 2023-05-22 13:05:09 | String | 
data.rows.update_time | 2023-05-22 13:16:48 | String | 
message | - | String | 
## /Admin/支付配置/新增支付渠道
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/payment

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 93fbbd25e371c2d9d275b1e5e3c6a5cc | String | 是 | -
#### 请求Body参数
```javascript
{
    "name": "易支付",
    "status": 0,
    "channel": "yipay",
    "types": "wxpay",
    "params": "{\"pid\":\"2222\",\"key\":\"333\",\"api\":\"444\",\"return_url\":\"5555\"}"
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
name | 易支付 | String | 是 | -
status | 0 | Integer | 是 | -
channel | yipay | String | 是 | -
types | wxpay | String | 是 | -
params | {"pid":"2222","key":"333","api":"444","return_url":"5555"} | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/支付配置/修改支付渠道
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/payment

#### 请求方式
> PUT

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 93fbbd25e371c2d9d275b1e5e3c6a5cc | String | 是 | -
#### 请求Body参数
```javascript
{
    "name": "易支付",
    "status": 0,
    "channel": "yipay",
    "types": "wxpay",
    "params": "{\"pid\":\"2222\",\"key\":\"333\",\"api\":\"444\",\"return_url\":\"5555\"}",
    "id": 546374734683674
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
name | 易支付 | String | 是 | -
status | 0 | Integer | 是 | -
channel | yipay | String | 是 | -
types | wxpay | String | 是 | -
params | {"pid":"2222","key":"333","api":"444","return_url":"5555"} | String | 是 | -
id | 546374734683674 | Integer | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/支付配置/删除渠道
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/payment/35625377

#### 请求方式
> DELETE

#### Content-Type
> form-data

#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | 93fbbd25e371c2d9d275b1e5e3c6a5cc | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/支付订单
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/支付订单/订单列表
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/orders

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
token | c22d69f0d8cdc60b6f64774f0a06bace | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 0,
	"data": {
		"count": 12,
		"rows": [
			{
				"id": 51370316848435200,
				"trade_no": null,
				"pay_type": "alipay",
				"product_id": 50420988902379520,
				"product_title": "测试001",
				"trade_status": "TRADE_AWAIT",
				"user_id": "49454828484562944",
				"product_info": "{\"id\":50420988902379520,\"title\":\"测试001\",\"price\":100,\"original_price\":0,\"integral\":0,\"badge\":\"测试\",\"day\":7,\"status\":1,\"create_time\":\"2023-05-20 03:15:01\",\"update_time\":\"2023-05-20 03:20:53\"}",
				"channel": "yipay",
				"params": "{\"order_id\":\"51370316848435200\",\"product_id\":50420988902379520,\"user_id\":\"49454828484562944\",\"payment_id\":\"51294279280103424\"}",
				"payment_id": "51294279280103424",
				"payment_info": "{\"id\":\"51294279280103424\",\"name\":\"易支付\",\"channel\":\"yipay\",\"types\":\"alipay,wxpay\",\"params\":\"{\\\"pid\\\":\\\"13\\\",\\\"key\\\":\\\"5HIFI\\\",\\\"api\\\":\\\"htom\\\",\\\"return_url\\\":\\\"\\\"}\",\"status\":1,\"create_time\":\"2023-05-22 13:05:09\",\"update_time\":\"2023-05-22 17:19:59\"}",
				"money": 1,
				"notify_info": null,
				"pay_url": null,
				"ip": "127.0.0.1",
				"create_time": "2023-05-22 18:07:18",
				"update_time": "2023-05-22 18:07:18"
			}
		]
	},
	"message": ""
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 0 | Integer | 
data | - | Object | 
data.count | 12 | Integer | 
data.rows | - | Array | 
data.rows.id | 51370316848435200 | Integer | 
data.rows.trade_no | null | Null | 
data.rows.pay_type | alipay | String | 
data.rows.product_id | 50420988902379520 | Integer | 
data.rows.product_title | 测试001 | String | 
data.rows.trade_status | TRADE_AWAIT | String | 
data.rows.user_id | 49454828484562944 | String | 
data.rows.product_info | {"id":50420988902379520,"title":"测试001","price":100,"original_price":0,"integral":0,"badge":"测试","day":7,"status":1,"create_time":"2023-05-20 03:15:01","update_time":"2023-05-20 03:20:53"} | String | 
data.rows.channel | yipay | String | 
data.rows.params | {"order_id":"51370316848435200","product_id":50420988902379520,"user_id":"49454828484562944","payment_id":"51294279280103424"} | String | 
data.rows.payment_id | 51294279280103424 | String | 
data.rows.payment_info | {"id":"51294279280103424","name":"易支付","channel":"yipay","types":"alipay,wxpay","params":"{\"pid\":\"13\",\"key\":\"5HIFI\",\"api\":\"htom\",\"return_url\":\"\"}","status":1,"create_time":"2023-05-22 13:05:09","update_time":"2023-05-22 17:19:59"} | String | 
data.rows.money | 1 | Integer | 
data.rows.notify_info | null | Null | 
data.rows.pay_url | null | Null | 
data.rows.ip | 127.0.0.1 | String | 
data.rows.create_time | 2023-05-22 18:07:18 | String | 
data.rows.update_time | 2023-05-22 18:07:18 | String | 
message | - | String | 
## /Admin/通知配置
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/通知配置/获取通知列表
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/notification

#### 请求方式
> GET

#### Content-Type
> form-data

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/通知配置/新增通知
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/notification

#### 请求方式
> POST

#### Content-Type
> form-data

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/通知配置/修改通知
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/notification

#### 请求方式
> PUT

#### Content-Type
> form-data

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /Admin/通知配置/删除通知
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/admin/notification/333454365436

#### 请求方式
> DELETE

#### Content-Type
> form-data

#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /绘画
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /绘画/sd文字生图片
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200https://api.stability.ai/v1/generation/stable-diffusion-v1-5/text-to-image

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
Content-Type | application/json | String | 是 | -
Accept | application/json | String | 是 | -
Authorization | Bearer sk- | String | 是 | -
#### 请求Body参数
```javascript
{"text_prompts":[{"text":"A lighthouse on a cliff"}],"cfg_scale":7,"clip_guidance_preset":"FAST_BLUE","height":512,"width":512,"samples":1,"steps":30,"style_preset":"3d-model"}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
text_prompts | - | Array | 是 | 描述
text_prompts.text | A lighthouse on a cliff | String | 是 | 描述
cfg_scale | 7 | Integer | 是 | 较高的值使您的图像更接近您的提示
clip_guidance_preset | FAST_BLUE | String | 是 | 速度
height | 512 | Integer | 是 | 高度
width | 512 | Integer | 是 | 宽度
samples | 1 | Integer | 是 | 生成数量
steps | 30 | Integer | 是 | 运行步骤
style_preset | 3d-model | String | 是 | 风格预设
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /绘画/sd2文字生图片
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200https://stablediffusionapi.com/api/v3/text2img

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Body参数
```javascript
{
  "key": "",
  "prompt": "ultra realistic close up portrait ((beautiful pale cyberpunk female with heavy black eyeliner))",
  "negative_prompt": null,
  "width": "512",
  "height": "512",
  "samples": "1",
  "num_inference_steps": "20",
  "safety_checker": "no",
  "enhance_prompt": "yes",
  "seed": null,
  "guidance_scale": 7.5,
  "multi_lingual": "no",
  "panorama": "no",
  "self_attention": "no",
  "upscale": "no",
  "embeddings_model": "embeddings_model_id",
  "webhook": null,
  "track_id": null
}
```
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
key | 1MxfzJQKR3ITk86mc89aVr32mM6Vjf4lv4PbrZzQS9v6foZd3h0g9rN0wzwe | String | 是 | -
prompt | ultra realistic close up portrait ((beautiful pale cyberpunk female with heavy black eyeliner)) | String | 是 | -
negative_prompt | null | Null | 是 | -
width | 512 | String | 是 | 宽度
height | 512 | String | 是 | 高度
samples | 1 | String | 是 | 生成数量
num_inference_steps | 20 | String | 是 | -
safety_checker | no | String | 是 | -
enhance_prompt | yes | String | 是 | -
seed | null | Null | 是 | -
guidance_scale | 7.5 | Number | 是 | -
multi_lingual | no | String | 是 | -
panorama | no | String | 是 | -
self_attention | no | String | 是 | -
upscale | no | String | 是 | -
embeddings_model | embeddings_model_id | String | 是 | -
webhook | null | Null | 是 | -
track_id | null | Null | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /绘画/openai文字生成图片
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200https://proxy.aizj.top/v1/images/generations

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
Authorization | Bearer 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 | String | 是 | -
#### 请求Body参数
```javascript
{
    "prompt": "熊猫",
    "response_format": "b64_json"
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /绘画/sd图片生图片
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200https://api.stability.ai/v1/generation/stable-diffusion-v1-5/image-to-image

#### 请求方式
> POST

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
Content-Type | application/json | String | 是 | -
Accept | application/json | String | 是 | -
Authorization | - | String | 是 | -
#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
text_prompts[0][text] | A lighthouse on acliff | String | 是 | -
cfg_scale | 7 | String | 是 | -
clip_guidance_preset | FAST_BLUE | String | 是 | -
height | 512 | String | 是 | -
width | 512 | String | 是 | -
samples | 1 | String | 是 | -
steps | 30 | String | 是 | -
style_preset | 3d-model | String | 是 | -
init_image | ["/Users/qiaoyue/Downloads/05c3260d-6a2e-4aa5-82f0-e952f2a5fa10-0.png"] | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /绘画/新建接口
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200https://api.stability.ai/v1/user/balance

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
Authorization | Bearer sk-fxVEA6IZLlmj6wPwCTGxRv8qfYnSSI79YxpIzCTTkbn4ztFT | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /绘画/上传图片
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200/api/upload

#### 请求方式
> POST

#### Content-Type
> form-data

#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
file | ["/Users/qy/Downloads/0cb0a1356df472940ba4fa76a44591e15776.png"] | String | 是 | -
file | - | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /openai
```text
暂无描述
```
#### Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /openai/获取模型
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200https://proxy.aizj.top/v1/models

#### 请求方式
> GET

#### Content-Type
> form-data

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
Authorization | Bearer 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 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /openai/对话
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200https://proxy.aizj.top/v1/chat/completions

#### 请求方式
> POST

#### Content-Type
> json

#### 请求Header参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
Authorization | Bearer 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 | String | 是 | -
#### 请求Body参数
```javascript
{
    "model": "gpt-3.5-turbo-16k",
    "messages": [{"role": "user", "content": "Hello!"}]
}
```
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /openai/新建接口
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200https://cli.im/Api/Browser/deqr

#### 请求方式
> POST

#### Content-Type
> form-data

#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
data | https://api.xunhupay.com/payments/wechat/qrcode?id=20233703604&nonce_str=6491456378&time=1687965443&appid=201906158184&hash=a8561c8ea773658b9cd5841089ce884f | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /openai/新建接口
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:3200https://qrdetector-api.cli.im/v1/detect_binary

#### 请求方式
> POST

#### Content-Type
> form-data

#### 请求Body参数
参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述
--- | --- | --- | --- | ---
image_path | https://api.xunhupay.com/qrcode/201906158184.html?data=d2VpeGluOi8vd3hwYXkvYml6cGF5dXJsP3ByPTdkcVlmWGR6eg==&nonce_str=5619761847&time=1687964157&hash=c3434ffb7d1a5c837efb56d1f1187555 | String | 是 | -
remove_background | 1 | String | 是 | -
#### 认证方式
```text
noauth
```
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```