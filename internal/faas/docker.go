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
