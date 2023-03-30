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
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/quanxiang-cloud/initdb/internal/db"
	"github.com/quanxiang-cloud/initdb/internal/faas"
	"github.com/quanxiang-cloud/initdb/internal/search"
	"github.com/quanxiang-cloud/initdb/internal/web"
	"github.com/quanxiang-cloud/initdb/pkg/toolkits"
	"github.com/vrischmann/envconfig"
)

func main() {

	var db db.DBConfig
	var git faas.GitConf
	var docker faas.DockerConf
	var esSearch search.SearchConf
	if err := envconfig.Init(&db); err != nil {
		panic(err)
	}

	if err := envconfig.Init(&git); err != nil {
		log.Panicln(err)
	}
	if err := envconfig.Init(&docker); err != nil {
		log.Panicln(err)
	}
	if err := envconfig.Init(&esSearch); err != nil {
		log.Panicln(err)
	}
	dirs := []string{}
	files, err := ioutil.ReadDir("./migrations")
	if err != nil {
		toolkits.Error(err)
	}
	for _, f := range files {
		dirs = append(dirs, f.Name())
	}
	for {
		if db.Pong() != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	err = db.CreateDB(dirs)
	if err != nil {
		toolkits.Error(err)
	}

	for _, d := range dirs {
		err = db.MigrateDB("./migrations/"+d, d)
		if err != nil {
			toolkits.Error(err)
		}
	}

	// Init faas git and docker
	dockerSQL := docker.GenDockerSQL()
	for i := 0; i < 5; i++ {
		err = db.ExecSql("faas", dockerSQL)
		if err != nil {
			toolkits.Error(fmt.Sprintf("retry exec docker sql %d times", i), err)
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}

	gitSQL := git.GenGitSQL()
	for i := 0; i < 5; i++ {
		err = db.ExecSql("faas", gitSQL)
		if err != nil {
			toolkits.Error(fmt.Sprintf("retry exec gitlab sql %d times", i), err)
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}

	// init nav
	i := 0
	for {
		if err := web.InitNav(); err != nil {
			toolkits.Error(fmt.Sprintf("retry init nav %d times", i), err)
			time.Sleep(5 * time.Second)
			i++
			continue
		}
		break
	}

	// init artery
	i = 0
	for {
		if err := web.InitArtery(); err != nil {
			time.Sleep(5 * time.Second)
			toolkits.Error(fmt.Sprintf("retry init artery %d times", i), err)
			i++
			continue
		}
		break
	}
	for {
		var i int
		if err := esSearch.InitSearch("user"); err != nil {
			time.Sleep(5 * time.Second)
			toolkits.Error(fmt.Sprintf("retry init search user %d times", i), err)
			i++
			continue
		}
		break
	}

}
