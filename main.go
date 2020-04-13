package main

import (
	"github.com/folk113/aliddns/config"
	"github.com/folk113/aliddns/dns"
	"os"
	"strings"
)

func main() {
	index := strings.LastIndex(os.Args[0], string(os.PathSeparator))
	path := os.Args[0][:index]
	config.Init(path)
	ip, err := dns.GetIp()
	if err != nil {
		return
	}
	config.Log.Infof("Get IP address:%s successfully", ip)
	dns.NewDns(config.AliAccessConfig.Domain, ip, config.AliAccessConfig.RRList).Bind()
}
