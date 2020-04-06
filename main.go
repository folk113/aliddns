package main

import (
	"github.com/folk113/aliddns/config"
	"github.com/folk113/aliddns/dns"
)

func main() {
	ip, err := dns.GetIp()
	if err != nil {
		return
	}
	config.Log.Infof("Get IP address:%s successfully", ip)
	dns.NewDns(config.AliAccessConfig.Domain, ip, config.AliAccessConfig.RRList).Bind()
}
