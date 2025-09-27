package service

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/smtp"
	"shop/code/config"
	"shop/code/errs"
	"shop/code/repository"
	"shop/code/utils"
	"strings"
)

type EmailInfo struct {
	from       string
	password   string
	smtpServer string
}
type CodeService struct {
	repo      repository.CodeRepository
	EmailInfo *EmailInfo
	tm        config.TemplateManager
}

func NewCodeService(repo repository.CodeRepository, tm config.TemplateManager) *CodeService {
	return &CodeService{repo: repo,
		EmailInfo: &EmailInfo{
			from:       config.Cf.Smtp.From,
			password:   config.Cf.Smtp.Password,
			smtpServer: config.Cf.Smtp.Host,
		},
		tm: tm,
	}
}

func (svc *CodeService) VerifyCode(ctx context.Context, addr string, subject string, code string) (bool, error) {
	codes, err := svc.repo.Get(ctx, addr, subject)
	if err != nil {
		return false, errs.ErrWrongCode
	}
	if codes == code {
		_ = svc.repo.Delete(ctx, addr, subject)
		return true, nil
	} else {
		return false, errs.ErrWrongCode
	}
}

func (svc *CodeService) SendCode(ctx context.Context, addr string, subject string) error {
	templatedetail, err := svc.tm.GetTemplate(subject)
	if err != nil {
		return errors.New("未找到对应业务")
	}
	captcha, err := utils.RandomDigitCaptcha(templatedetail.CodeLength)
	if err != nil {
		return err
	}
	err = svc.repo.Create(ctx, addr, subject, captcha, templatedetail.ExpireSeconds)
	if errors.Is(err, errs.ErrTooManyRequest) {
		zap.L().Info(fmt.Sprintf("用户 %s 请求发送 %s 服务验证码还未过期", addr, subject))
		return errs.ErrTooManyRequest
	}
	var msg string
	msg, err = svc.tm.RenderContent(subject, captcha)
	if err != nil {
		return err
	}
	if strings.Contains(addr, "@") {
		return svc.SendCodeByEmail(ctx, addr, subject, captcha, msg)
	} else {
		return svc.SendCodeByPhone(ctx, addr, subject, captcha, msg)
	}
}
func (svc *CodeService) SendCodeByPhone(ctx context.Context, phone string, subject string, code string, message string) error {
	panic("implement me")
}

func (svc *CodeService) SendCodeByEmail(ctx context.Context, email string, subject string, code string, message string) error {
	return svc.SendEmail(ctx, email, subject, message)
}

func (svc *CodeService) SendEmail(ctx context.Context, sendTo string, subject string, body string) error {
	from := svc.EmailInfo.from
	password := svc.EmailInfo.password // 邮箱授权码
	smtpServer := svc.EmailInfo.smtpServer

	// 设置 PlainAuth
	auth := smtp.PlainAuth("", from, password, "smtp.qq.com")

	// 创建 tls 配置
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.qq.com",
	}

	// 连接到 SMTP 服务器
	conn, err := tls.Dial("tcp", smtpServer, tlsconfig)
	if err != nil {
		return fmt.Errorf("TLS 连接失败: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, "smtp.qq.com")
	if err != nil {
		return fmt.Errorf("SMTP 客户端创建失败: %v", err)
	}
	defer client.Quit()

	// 使用 auth 进行认证
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("认证失败: %v", err)
	}

	// 设置发件人和收件人
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("发件人设置失败: %v", err)
	}
	if err = client.Rcpt(sendTo); err != nil {
		return fmt.Errorf("收件人设置失败: %v", err)
	}

	// 写入邮件内容
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("数据写入失败: %v", err)
	}
	defer wc.Close()

	msg := []byte("From: shop团队<" + from + ">\r\n" +
		"To: " + sendTo + "\r\n" +
		"Subject: 你的一次性代码" + "\r\n" +
		"\r\n" +
		body + "\r\n")
	_, err = wc.Write(msg)
	if err != nil {
		return fmt.Errorf("消息发送失败: %v", err)
	}

	return nil
}
