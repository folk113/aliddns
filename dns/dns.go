package dns

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/folk113/aliddns/config"
	"time"
)

type Dns struct {
	client *alidns.Client
	IP     string
	Domain string
	RRList []string
}

func NewDns(domain, ip string, rrList []string) *Dns {
	if domain == "" || ip == "" || len(rrList) == 0 {
		panic(fmt.Errorf("domain ip or rr cannot be empty"))
	}
	client, err := alidns.NewClientWithAccessKey(config.AliAccessConfig.Region,
		config.AliAccessConfig.AccessKeyID,
		config.AliAccessConfig.AccessKeySecrete)
	if err != nil {
		panic(fmt.Errorf("new alidns client failed: %v", err))
	}
	return &Dns{
		client: client,
		IP:     ip,
		Domain: domain,
		RRList: rrList,
	}
}

type Record struct {
	ShouldUpdate bool
	RecordId     string
	RecordValue  string
	RR           string
}

func (dns *Dns) Bind() {
	recordResp, err := dns.findRecords()
	if err != nil {
		return
	}
	records := recordResp.DomainRecords.Record
	recordArr := make([]Record, len(dns.RRList))

	for _, record := range records {
		for index, rr := range dns.RRList {
			if record.RR == rr {
				// 如果找到RR和输入里的rr相同的记录，则更新这条记录的解析。反之则添加一条新解析
				recordArr[index].ShouldUpdate = true
				recordArr[index].RecordId = record.RecordId
				recordArr[index].RecordValue = record.Value
				recordArr[index].RR = rr
				break
			}
		}
	}

	// add
	for _, rrAddr := range recordArr {
		if !rrAddr.ShouldUpdate {
			config.Log.Infof("add domain record")
			if _, err := dns.addRecord(rrAddr.RR); err != nil {
				config.Log.Errorf("add domain %s:%s failed", rrAddr.RecordValue, rrAddr.RR)
			}
		} else {
			if rrAddr.RecordValue == dns.IP {
				config.Log.Infof("ip not changed, no need updating")
			} else {
				_, err := dns.updateRecord(rrAddr.RR,rrAddr.RecordId)
				// 本来一个条件表达式搞定的问题，导致写下一堆完成本来应该一行搞定的事情
				var result string
				if err != nil {
					result = "failed"
				} else {
					result = "success"
				}
				config.Log.Errorf("update domain %s:%s %s", rrAddr.RecordValue, rrAddr.RR, result)
			}
		}
	}
}

func (dns *Dns) findRecords() (*alidns.DescribeDomainRecordsResponse, error) {
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = dns.Domain
	resp, err := dns.client.DescribeDomainRecords(request)
	if err != nil {
		// try to fix timeout issue
		if clientErr, ok := err.(*errors.ClientError); ok && clientErr.ErrorCode() == errors.TimeoutErrorCode {
			// retry
			config.Log.Errorf("timeout. retry...", clientErr)
			time.Sleep(time.Second)
			return dns.findRecords()
		}
		config.Log.Errorf("finding records failed", err)
		return nil, fmt.Errorf("finding records failed: %v", err)
	}
	return resp, nil
}

func (dns *Dns) addRecord(rr string) (*alidns.AddDomainRecordResponse, error) {
	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = dns.Domain
	request.Type = "A"
	request.RR = rr
	request.Value = dns.IP
	resp, err := dns.client.AddDomainRecord(request)
	if err != nil {
		config.Log.Errorf("adding record failed", err)
		return nil, fmt.Errorf("adding record failed: %v", err)
	}
	config.Log.Infof(fmt.Sprintf(`set ip of '%s.%s' to %s`, rr, dns.Domain, dns.IP))
	return resp, nil
}

func (dns *Dns) updateRecord(rr, recordId string) (*alidns.UpdateDomainRecordResponse, error) {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.RecordId = recordId
	request.Type = "A"
	request.RR = rr
	request.Value = dns.IP
	resp, err := dns.client.UpdateDomainRecord(request)
	if err != nil {
		config.Log.Errorf("updating record failed", err)
		return nil, fmt.Errorf("updating record:%s failed reason: %v", recordId, err)
	}
	config.Log.Infof(fmt.Sprintf(`set ip of '%s.%s' to %s`, rr, dns.Domain, dns.IP))
	return resp, nil
}
