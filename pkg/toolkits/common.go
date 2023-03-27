package toolkits

import (
	"log"
	"net/http"
	"strings"
)

func DoPost(url, data string) error {
	dataReq := strings.NewReader(data)
	req, _ := http.NewRequest("POST", url, dataReq)
	req.Header.Add("Content-Type", "application/json")
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func DoPUT(url, data string) error {
	dataReq := strings.NewReader(data)
	req, _ := http.NewRequest("PUT", url, dataReq)
	req.Header.Add("Content-Type", "application/json")
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func Error(v ...interface{}) {
	log.SetPrefix("ERROR: ")
	log.Println(v...)
}

func Info(v ...interface{}) {
	log.SetPrefix("INFO: ")
	log.Println(v...)
}
