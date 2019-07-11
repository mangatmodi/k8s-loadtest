package task

import (
	"container/ring"
	"crypto/tls"
	"fmt"
	"github.com/mangatmodi/k8s-loadtest/slave/util"
	"github.com/google/uuid"
	"github.com/myzhan/boomer"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
	"time"
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
func makeConv() {
	tracker, err := util.GetEnv("TRACKER_URL")
	tracker_url := "%s/conv?click_id=%s"
	client := TrackerConvTask.ctx["http"].(*fasthttp.Client)

	//no click with a conversion yet
	clickId, _ := uuid.NewRandom()
	test_url := fmt.Sprintf(tracker_url, tracker, clickId.String())

	request := fasthttp.AcquireRequest()
	request.SetRequestURI(test_url)
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()

	startTime := boomer.Now()
	err = client.Do(request, response)
	elapsed := boomer.Now() - startTime

	if err != nil {
		log.Fatalf("%v\n", err)
		boomer.RecordFailure(convTaskName, "error", 0.0, err.Error())
	} else {
		boomer.RecordSuccess(convTaskName, strconv.Itoa(response.StatusCode()), elapsed, int64(response.Header.ContentLength()))
	}

	fasthttp.ReleaseRequest(request)
}
func getClick() string {
	value := ClickIds.Value
	if value == nil {
		return ""
	}
	return value.(string)
}
