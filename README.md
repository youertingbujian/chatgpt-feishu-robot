

## 1. 基本介绍

### 1.1 项目介绍

> Golang开发的ChatGPT，接入飞书自建应用机器人，支持单聊和群聊

## 2. 使用说明

```
- golang版本 >= v1.19
- IDE推荐：Goland
```

### 2.1 克隆项目

使用 `Goland` 等编辑工具，打开

```bash

# 克隆项目
git clone https://github.com/youertingbujian/chatgpt-feishu-robot.git
# 进入chatgpt-feishu-robot
cd chatgpt-feishu-robot

# 使用 go mod 并安装go依赖包
go mod tidy

# 修改config.go中对应的key

# 启动
go run main.go
```

### 2.2 创建飞书自建应用
```
1、获取app id和app secret
2、开启机器人
3、开启事件Encrypt Key
4、配置请求地址 "https://xxxxxxx/api/v1/webhook/event"
5、添加"接收消息"
6、开通"接收群聊中@机器人消息事件"、"获取用户发给机器人的单聊消息"、"读取用户发给机器人的单聊消息"、"获取用户在群组中@机器人的消息"权限
```

## 生产部署
### 1、nginx反向代理
```
后端端口 51515 （main.go中可自行修改）
```
### 2、生产部署文件
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

### 3、nohup方式部署amd64文件
```
sudo nohup ./bin/chatgpt  > nohup_chatgpt.log 2>&1 &
```
### 4、停止
```
ps -ef | grep chatgpt

kill -9 pid
```