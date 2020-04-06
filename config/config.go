package config

import (
	"encoding/json"
	"errors"
	"github.com/folk113/aliddns/log"
	"io/ioutil"
)

var AliAccessConfig = &Config{}
var Log *log.Logger

func init() {
	Log = log.NewLogger()

	data, err := ioutil.ReadFile("./config/config.json")
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
