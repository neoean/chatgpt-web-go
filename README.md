<div align="center">
<img src="./docs/openai.svg" style="width:64px;height:64px;margin:0 32px" alt="icon"/>

<h1 align="center">ChatGPT Web Go</h1>

A commercially-viable ChatGpt web application built with Go.
The corresponding front-end service is: [79E/ChatGPT-Web](https://github.com/79E/ChatGPT-Web/)

可部署商业化的 ChatGpt 网页应用。
对应的前端服务：[79E/ChatGPT-Web](https://github.com/79E/ChatGPT-Web/) 


[Issues](https://github.com/heimeropen/chatgpt-web-go/issues)


</div>

## 交流&赞助

<a href='https://t.me/+DDQufJfXm9s2OTQx' target='_blank'>
<img width='46%' style="border-radius: 12px;" src='https://www.helloimg.com/images/2023/06/20/otDPwM.png' />
</a>


## 🐶 演示
### 页面链接
TODO

如需帮助请提交 [Issues](https://github.com/heimeropen/chatgpt-web-go/issues) 或赞赏时留下联系方式。

### 页面截图

![cover](https://files.catbox.moe/tp963e.png)
![cover](https://files.catbox.moe/y5avbx.png)
![cover](https://files.catbox.moe/k16jsz.png)
![cover](https://files.catbox.moe/8o5oja.png)

## 🤖 主要功能

- 后台管理系统,可对用户,Token,商品,卡密等进行管理
- 精心设计的 UI，响应式设计
- 极快的首屏加载速度（~100kb）
- 支持Midjourney绘画和DALL·E模型绘画,GPT4等应用
- 海量的内置 prompt 列表，来自[中文](https://github.com/PlexPt/awesome-chatgpt-prompts-zh)和[英文](https://github.com/f/awesome-chatgpt-prompts)
- 一键导出聊天记录，完整的 Markdown 支持
- 支持自定义API地址（如：[openAI](https://api.openai.com) / [API2D](https://api2d.com/r/192767)）

## 🎮 开始使用
**1.先 `Fork` 本项目，然后克隆到本地。**
```
git clone https://github.com/heimeropen/chatgpt-web-go.git
```

**2.安装依赖**
```
# 倒入sql 
./model/sql/chatgpt.sql

# 编译&拉取依赖
go build
```

**3.运行**
```
# http server 
./cmd/server/main.go
```

**4.打包**
```
golang 正常打包
```

## ⛺️ 环境变量

前端项目需要修改配置 
#### `VITE_APP_REQUEST_HOST`


## 🚧 开发

> goland 


## 💰 赞助方



## 🧘 贡献者


## 📋 开源协议

[![License MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://github.com/79E/ChatGpt-Web/blob/master/license)
