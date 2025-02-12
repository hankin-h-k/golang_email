package main

import (
	"github.com/hankin-h-k/golang_email/model"
	"github.com/hankin-h-k/golang_email/service"
)

func main() {
	smtpHost := "smtp.126.com" // SMTP 服务器地址（如 smtp.gmail.com）
	smtpPort := 25             // SMTP 端口（Gmail 为 587）
	smtpUser := "hankin_h"     // SMTP 用户名（通常为邮箱地址）
	smtpPass := ""             // SMTP 密码或应用专用密码
	smtpFrom := "hankin_h@126.com"
	d := service.NewClient(smtpHost, smtpPort, smtpUser, smtpPass, smtpFrom)

	var msg model.EmailMessage
	msg.To = "776514654@qq.com"
	msg.Subject = "Go 邮件测试"
	content := "这是一封来自 Go 的测试邮件"
	msg.TextContent = &content
	// htmlContent := "<h1>HTML 内容</h1><p>这是一封 HTML 邮件</p>"
	// msg.HtmlContent = &htmlContent
	// htmlPath := fmt.Sprintf("%s/email/index.html", str)
	// msg.HtmlPath = &htmlPath

	// str, _ := os.Getwd()
	// attachPath1 := fmt.Sprintf("%s/email/index.html", str)
	// attachPath2 := fmt.Sprintf("%s/email/index.html", str)
	// attachPath3 := fmt.Sprintf("%s/email/index.html", str)

	// msg.AttachPath = []string{attachPath1, attachPath2, attachPath3}
	d.SendEmail(&msg)
}
