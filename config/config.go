package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/folk113/aliddns/log"
	"io/ioutil"
	"os"
)

var AliAccessConfig = &Config{}
var Log *log.Logger

func Init(projectPath string) {
	Log = log.NewLogger()
	path := fmt.Sprintf("%s%cconfig%c", projectPath, os.PathSeparator,os.PathSeparator)
	data, err := ioutil.ReadFile(path + "config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, AliAccessConfig)
	if err != nil {
		panic(err)
	}
	if AliAccessConfig.Region == "" ||
		AliAccessConfig.AccessKeyID == "" ||
		AliAccessConfig.AccessKeySecrete == "" {
		panic(errors.New("配置信息为空"))
	}
}

type Config struct {
	Region           string   `json:"REGION"`
	AccessKeyID      string   `json:"ACCESS_KEY_ID"`
	AccessKeySecrete string   `json:"ACCESS_KEY_SECRET"`
	Domain           string   `json:"DOMAIN"`
	RRList           []string `json:"RRList"`
}
