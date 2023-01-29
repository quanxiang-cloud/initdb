package main

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

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/quanxiang-cloud/initdb/internal/db"
	"github.com/vrischmann/envconfig"
)

func main() {
	var db db.DBConfig
	if err := envconfig.Init(&db); err != nil {
		fmt.Println(err)
	}
	/*
		dbPath, err := filepath.Abs("./migrations")
		if err != nil {
			fmt.Println(err)
		}
	*/
	dirs := []string{}
	files, err := ioutil.ReadDir("./migrations")
	if err != nil {
		fmt.Println(err)
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

	for _, d := range dirs {
		err = db.MigrateDB("./migrations/"+d, d)
	}

	fmt.Println(err)
}
