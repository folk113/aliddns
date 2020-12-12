package dns

import (
	"encoding/json"
	"github.com/folk113/aliddns/config"
	"io/ioutil"
	"net/http"
)

type IpAddress struct {
	IP      string `json:"ip"`
	GeoIp   string `json:"geo-ip"`
	ApiHelp string `json:"API Help"`
}

func GetIp() (string, error) {
	url := "https://ipv4.jsonip.com"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		config.Log.Error(err)
	}
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		config.Log.Error(err)
	}
	ipAddress := IpAddress{}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		config.Log.Error(err)
	}
	err = json.Unmarshal(data, &ipAddress)
	if err != nil {
		config.Log.Error(err)
	}
	return ipAddress.IP, err
}
