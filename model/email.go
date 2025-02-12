package model

type EmailMessage struct {
	To          string   // 被通知邮箱
	Subject     string   // 主题
	TextContent *string  // 文本内容
	HtmlContent *string  // HTML内容
	HtmlPath    *string  // HTML文件路径
	AttachPath  []string // 附件路径
}
