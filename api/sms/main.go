// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sms

import (
	"log"

	"github.com/paingha/winkel/api/config"
	"github.com/sfreiberg/gotwilio"
	"github.com/sirupsen/logrus"
)

var (
	cfg config.SystemConfig
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//Messages - messages struct
type Messages struct {
	Content string
	To      string
	Medium  string
}

func init() {
	logrus.Info("Starting service ... ")
	err := config.InitConfig(&cfg)
	if err != nil {
		logrus.Fatalf("load config %v", err)
	}
}

func SendSMS(smsParam Messages) {
	twilio := gotwilio.NewTwilioClient(cfg.TwilioAccountSid, cfg.TwilioAuthToken)
	sendingSms(smsParam, twilio)
}

//SendWhatsapp - Sends Whatsapp text message
func (m *Messages) SendWhatsapp(credentials *gotwilio.Twilio) (*gotwilio.SmsResponse, error) {
	resp, _, err := credentials.SendWhatsApp(cfg.SenderPhone, m.To, m.Content, "", "")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//SendSms - Sends Sms text message
func (m *Messages) SendSms(credentials *gotwilio.Twilio) (*gotwilio.SmsResponse, error) {
	resp, _, err := credentials.SendSMS(cfg.SenderPhone, m.To, m.Content, "", "")
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func sendingSms(smsParam Messages, twilio *gotwilio.Twilio) {
	logrus.Infof("Sending to %s ... ", smsParam.To)
	if smsParam.Medium == "whatsapp" {
		if _, err := smsParam.SendWhatsapp(twilio); err != nil {
			logrus.Errorf("whatsapp sending error: ...", err.Error())
		}
		logrus.Info("whatsapp message sent... ")
	} else {
		if _, err := smsParam.SendSms(twilio); err != nil {
			logrus.Errorf("sms sending error: ...", err.Error())
		}
		logrus.Info("sms sent... ")
	}

}
