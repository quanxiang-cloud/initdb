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

package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DBConfig struct {
	Type     string `envconfig:"default=mysql,DBTYPE"`
	Host     string `envconfig:"default=localhost,DBHOST"`
	UserName string `envconfig:"default=root,DBUSER"`
	Password string `envconfig:"DBPASSWORD"`
}

func (d *DBConfig) MigrateDB(dbPaths, dbName string) (err error) {
	conn, err := sql.Open(d.Type, d.dsn(dbName))
	if err != nil {
		fmt.Println(1)
		return err
	}
	defer conn.Close()
	driver, err := mysql.WithInstance(conn, &mysql.Config{})
	if err != nil {
		fmt.Println(2)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", dbPaths), "mysql", driver)
	if err != nil {
		fmt.Println(3)
		return err
	}
	err = m.Up()
	if err != nil {
		fmt.Println(4)
		return err
	}
	return err
}

func (d *DBConfig) CreateDB(dbNames []string) error {
	conn, err := sql.Open("mysql", d.dsn(""))
	if err != nil {
		return err
	}
	defer conn.Close()
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	for _, v := range dbNames {
		res, err := conn.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default character set utf8mb4 collate utf8mb4_unicode_ci;", v))
		if err != nil {
			return err
		}
		no, err := res.RowsAffected()
		if err != nil {
			fmt.Printf("Error %s when fetching rows", err)
			return err
		}
		fmt.Printf("rows affected %d\n", no)
	}
	return err
}

func (d *DBConfig) ExecSql(dbName, sqlStr string) error {
	conn, err := sql.Open("mysql", d.dsn(dbName))
	if err != nil {
		return err
	}
	defer conn.Close()
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := conn.ExecContext(ctx, sqlStr)
	if err != nil {
		fmt.Println("Failed to exec sql:  ", sqlStr)
		return err
	}
	no, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error %s when fetching rows", err)
	}
	fmt.Printf("rows affected %d", no)
	return err
}

func (d *DBConfig) dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true", d.UserName, d.Password, d.Host, dbName)
}

func (d *DBConfig) Pong() error {
	conn, _ := sql.Open("mysql", d.dsn(""))
	return conn.Ping()
}
