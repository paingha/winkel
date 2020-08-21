// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mailer

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/paingha/winkel/api/config"
	"github.com/paingha/winkel/api/plugins"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	cfg config.SystemConfig
)

//EmailParam - struct for email message
type EmailParam struct {
	Template  string            `json:"template"`
	To        string            `json:"to"`
	Subject   string            `json:"subject"`
	BodyParam map[string]string `json:"body_param"`
}

func handleError(err error, msg string) {
	if err != nil {
		plugins.LogFatal("MailService", msg, err)
	}
}

func init() {
	files := map[string]string{
		"TemplateVerifyEmail": "./mailer/templates/verify.html",
		"TemplateResetEmail":  "./mailer/templates/password-reset.html",
	}
	getFilesContents(files)
}

//SendNow - sends email through a go routine
func SendNow(emailParam EmailParam) {
	sendEmail(emailParam)
}

func sendEmail(emailParam EmailParam) {
	var bodyTemplate string
	plugins.LogInfo("MailService: Sending mail to ", emailParam.To)
	if t, ok := EmailTemplates[emailParam.Template]; ok {
		bodyTemplate = t
	} else {
		bodyTemplate = EmailTemplates["TemplateVerifyEmail"]
	}
	parsedTemplate, err := template.New("template").Parse(bodyTemplate)
	if err != nil {
		plugins.LogError("MailService", "error parsing email template", err)
	}
	var buf bytes.Buffer
	err = parsedTemplate.Execute(&buf, emailParam.BodyParam)

	from := mail.NewEmail("Winkel", cfg.SenderEmail)
	subject := emailParam.Subject
	to := mail.NewEmail(fmt.Sprintf("%s %s ", emailParam.BodyParam["first_name"], emailParam.BodyParam["last_name"]), emailParam.To)

	htmlContent := buf.String()
	message := mail.NewSingleEmail(from, subject, to, "text", htmlContent)

	client := sendgrid.NewSendClient(cfg.SendgridAPIKey)
	response, err := client.Send(message)
	fmt.Println(response.StatusCode)
	if err != nil {
		plugins.LogError("MailService", "error sending email", err)
	} else {
		plugins.LogInfo("MailService: Email sent successfully", "200")
	}
}
