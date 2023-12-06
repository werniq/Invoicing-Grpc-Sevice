package main

import (
	"bytes"
	"embed"
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"html/template"
	"time"
)

//go:embed email-tmps
var emailTemplateFS embed.FS

func (app *Application) SendMail(
	from, to, subj, tmpl string,
	attachments []string,
	data interface{}) error {
	var err error

	templateToRender := fmt.Sprintf("email-tmps/%s.html.tmpl", tmpl)

	t, err := template.New("email-html").ParseFS(emailTemplateFS, templateToRender)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", data); err != nil {
		return err
	}

	formattedMsg := tpl.String()

	templateToRender = fmt.Sprintf("email-tmps/%s.plain.tmpl", tmpl)

	t, err = template.New("email-plain").ParseFS(emailTemplateFS, templateToRender)
	if err != nil {
		return err
	}

	if err = t.ExecuteTemplate(&tpl, "body", data); err != nil {
		return err
	}

	plainMsg := tpl.String()

	server := mail.NewSMTPClient()

	// necessary configurations
	server.Host = app.config.SMTP.host
	server.Port = app.config.SMTP.port
	server.Username = app.config.SMTP.user
	server.Password = app.config.SMTP.pwd

	// server options
	server.Encryption = mail.EncryptionTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.
		SetFrom(from).
		AddTo(to).
		SetSubject(subj)

	email.SetBody(mail.TextHTML, formattedMsg)
	email.AddAlternative(mail.TextPlain, plainMsg)

	if len(attachments) > 0 {
		for _, x := range attachments {
			email.AddAttachment(x)
		}
	}

	if err = email.Send(smtpClient); err != nil {
		return err
	}

	app.infoLogger.Println("Email successfully sent.")

	return nil
}
