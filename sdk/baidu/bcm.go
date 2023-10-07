package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/baidubce/bce-sdk-go/services/bcm"
)

func main() {
	ak := "8600e073a6524e4498846b5b0f8799b6"
	sk := "41115cd08ebe47b19a5729c3a2c441eb"
	endpoint := "http://xxxx:8869"
	bcmClient, _ := bcm.NewClient(ak, sk, endpoint)

	dimensions := map[string]string{
		"InstanceId": "4fbc1dfa-eec1-4d04-9d3f-96da349164b5",
	}
	req := &bcm.GetMetricDataRequest{
		UserId:         "453bf9588c9e488f9ba2c984129090dc",
		Scope:          "BCE_BCC",
		MetricName:     "vCPUUsagePercent",
		Dimensions:     dimensions,
		Statistics:     strings.Split(bcm.Average+","+bcm.SampleCount+","+bcm.Sum+","+bcm.Minimum+","+bcm.Maximum, ","),
		PeriodInSecond: 60,
		StartTime:      time.Now().UTC().Add(-2 * time.Hour).Format("2006-01-02T15:04:05Z"),
		EndTime:        time.Now().UTC().Add(-1 * time.Hour).Format("2006-01-02T15:04:05Z"),
	}
	resp, err := bcmClient.GetMetricData(req)
	fmt.Printf("err is %v\n", err)
	content, _ := json.Marshal(resp)
	fmt.Println(string(content))
}
