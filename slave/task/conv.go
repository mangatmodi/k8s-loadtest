package task

import (
	"container/ring"
	"crypto/tls"
	"time"

	"github.com/valyala/fasthttp"
)

var convTaskName = "tracker-conv"
var ClickIds = ring.New(1000)

func buildTrackerConvTask() {
	TrackerConvTask.ctx["http"] = &fasthttp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NoDefaultUserAgentHeader: false,
		MaxIdleConnDuration:      5 * time.Minute,
		ReadTimeout:              2 * time.Second,
		MaxConnsPerHost:          100,
	}
}

//Dummy
func makeConv() {
}
func getClick() string {
	value := ClickIds.Value
	if value == nil {
		return ""
	}
	return value.(string)
}
