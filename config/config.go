package config

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

var Sysconfig = &sysconfig{}

func init() {
	dir, err := os.Getwd()
	fmt.Println(dir)
	b, err := ioutil.ReadFile(dir + "/config.json")
	if err != nil {
		fmt.Println(err)
		panic("Sys config read err")
	}

	err = jsoniter.Unmarshal(b, Sysconfig)
	if err != nil {
		panic(err)
	}

}

type sysconfig struct {
	Port       string `json:"Port"`
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBIp       string `json:"DBIp"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
}
