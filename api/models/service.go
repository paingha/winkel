// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import "io"

//EmailParam - email service sending structure
type EmailParam struct {
	Template  string            `json:"template"`
	To        string            `json:"to"`
	Subject   string            `json:"subject"`
	BodyParam map[string]string `json:"body_param"`
}

//Message - Whatsapp message structure
type Message struct {
	Content string
	To      string
	Medium  string
}

//Push - Push message structure
type Push struct {
	Content map[string]string //map[string]string{"en": pushParam.Content},
	Players []string
	IsIOS   bool
}

//FileParam - file upload structure
type FileParam struct {
	Name   string
	File   io.Reader
	Medium string
}
