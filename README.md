# golang_email

#### 介绍

发送邮箱信息


#### 安装教程

```
go get github.com/hankin-h-k/golang_email
```

#### 使用说明

1.  创建client

```
import "gitee.com/hankin_h/golang_email"

client := golang_email.service.NewClient(smtpHost, smtpPort, smtpUser, smtpPass, smtpPass)

var msg golang_email.model.EmailMessage
msg.To = "776514654@qq.com"
msg.Subject = "Go 邮件测试"
content := "这是一封来自 Go 的测试邮件"
msg.TextContent = &content

client.SendEmail(&msg)

```


