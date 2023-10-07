package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/baidubce/bce-sdk-go/services/bcm"
)

func main() {
	ak := "xxxx"
	sk := "xxxx"
	endpoint := "http://xxxx:8869"
	bcmClient, _ := bcm.NewClient(ak, sk, endpoint)

	dimensions := map[string]string{
		"InstanceId": "xxx",
	}
	req := &bcm.GetMetricDataRequest{
		UserId:         "xxx",
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
