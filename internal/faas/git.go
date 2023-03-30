/*
Copyright 2022 QuanxiangCloud Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package faas

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type GitConf struct {
	Name             string `envconfig:"default=rsa,GITNAME"`
	Host             string `envconfig:"default=localhost,GITHOST"`
	KnownHostsScan   string `envconfig:"KNOWNHOSTSSCAN"`
	GitSSHHost       string `envconfig:"GITSSHHOST"`
	GitSSHPort       int    `envconfig:"GITSSHPORT"`
	GitToken         string `envconfig:"GITTOKEN"`
	GitSshPrivateKey string `envconfig:"GITSSHPRIVATEKEY"`
}

func (g *GitConf) GenGitSQL() string {
	knowHosts := fmt.Sprintf("ssh://git@%s:%d/", g.GitSSHHost, g.GitSSHPort)
	sshKey := DecodeBase64String(g.GitSshPrivateKey)
	return fmt.Sprintf("insert into gits (id, host, token, name, known_hosts, key_scan_known_hosts, ssh) values('%s', '%s','%s', '%s','%s', '%s', '%s');",
		"mzHjx1QR", g.Host, g.GitToken, g.Name, knowHosts, g.KnownHostsScan, string(sshKey))

}

func DecodeBase64String(enc string) string {
	reader := strings.NewReader(enc)
	decoder := base64.NewDecoder(base64.RawStdEncoding, reader)
	buf := make([]byte, 1024)
	dst := ""
	for {
		n, err := decoder.Read(buf)
		dst += string(buf[:n])
		if err != nil || n == 0 {
			break
		}
	}
	return dst
}
