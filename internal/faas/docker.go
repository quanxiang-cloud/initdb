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

import "fmt"

type DockerConf struct {
	Name       string `envconfig:"DOCKERNAME"`
	Registry   string `envconfig:"DOCKERREGISTRY"`
	NameSpace  string `envconfig:"DOCKERNAMESPACE"`
	DockerUser string `envconfig:"DOCKERUSER"`
	DockerPwd  string `envconfig:"DOCKERPWD"`
}

func (d *DockerConf) GenDockerSQL() string {
	return fmt.Sprintf("insert into dockers (id, host, user_name, name_space, secret, name) values('aZhvb2qR', '%s', '%s', '%s', '%s', '%s');",
		d.Registry, d.DockerUser, d.NameSpace, d.DockerPwd, d.Name)
}
