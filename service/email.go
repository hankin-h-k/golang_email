package service

import (
	"email/model"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

type emailClient struct {
	From string
	Dial *gomail.Dialer
}

var client *emailClient

func NewClient(host string, port int, user string, pass string, from string) *emailClient {

	if client != nil {
		return client
	}
	// 配置邮件参数
	d := gomail.NewDialer(host, port, user, pass)

	emailClient := emailClient{
		From: from,
		Dial: d,
	}
	client = &emailClient
	return &emailClient
}

func (c *emailClient) SendEmail(message *model.EmailMessage) bool {

	m := gomail.NewMessage()
	m.SetHeader("From", client.From)
	m.SetHeader("To", message.To)
	m.SetHeader("Subject", message.Subject)
	if message.TextContent != nil {
		m.SetBody("text/plain", *message.TextContent) // 纯文本
	}

	if message.HtmlContent != nil {
		m.SetBody("text/html", *message.HtmlContent) // html
	}

	if message.HtmlPath != nil {
		data, err := os.ReadFile(*message.HtmlPath)
		if err != nil {
			log.Println("读取文件出错:", err)
			return false
		}
		m.SetBody("text/html", string(data)) // html文件
	}

	for _, attachPath := range message.AttachPath {
		m.Attach(attachPath) // 附件地址
	}

	if err := client.Dial.DialAndSend(m); err != nil {
		log.Println("邮件发送失败：", err)
		return false
	}

	log.Println("邮件发送成功！")
	return true
}
