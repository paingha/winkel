// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mailer

import (
	"io/ioutil"
	"sync"

	"github.com/paingha/winkel/api/plugins"
)

var (
	//EmailTemplates - email template map
	EmailTemplates = map[string]string{}
)

func getFilesContents(files map[string]string) {
	var wg sync.WaitGroup
	var m sync.Mutex

	filesLength := len(files)
	wg.Add(filesLength)

	for key, file := range files {
		go func(key, file string) {
			content, err := ioutil.ReadFile(file)
			if err != nil {
				plugins.LogError("MailService", "error reading "+key+" email template file", err)
			}
			m.Lock()
			EmailTemplates[key] = string(content)
			m.Unlock()
			wg.Done()
		}(key, file)
	}

	wg.Wait()

}
