package task

import (
	"crypto/tls"
	"fmt"
	"github.com/mangatmodi/k8s-loadtest/slave/data"
	"github.com/mangatmodi/k8s-loadtest/slave/util"
	"github.com/myzhan/boomer"
	"github.com/valyala/fasthttp"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

var taskName = "tracker-click"
var convRate = 1 // percentage

func buildTrackerClickTask() {
	links, err := data.GetOfferPublisherLink()
	if err != nil {
		panic(fmt.Sprintf("Error while getting offer publisher link: %+v\n", err))
	}
	TrackerClickTask.Data = links
	TrackerClickTask.ctx["http"] = &fasthttp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NoDefaultUserAgentHeader: false,
		MaxIdleConnDuration:      5 * time.Minute,
		ReadTimeout:              2 * time.Second,
		MaxConnsPerHost:          1000,
	}
}
func makeClick() {
	tracker, err := util.GetEnv("TRACKER_URL")
	tracker_url := "%s/click?offer_id=%s&pub_id=%s"
	germanIp := "213.61.59.250"
	links := TrackerClickTask.Data.([]*data.OfferPublisherLink)
	test_url := fmt.Sprintf(tracker_url, tracker, links[0].OfferId, links[0].PublisherId)
	client := TrackerClickTask.ctx["http"].(*fasthttp.Client)

	request := fasthttp.AcquireRequest()
	request.SetRequestURI(test_url)
	request.Header.SetMethod("GET")
	request.Header.Set("HTTP_X_FORWARDED_FOR", germanIp)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 7.1.1; G8231 Build/41.2.A.0.219; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/59.0.3071.125 Mobile Safari/537.36")

	response := fasthttp.AcquireResponse()
	startTime := boomer.Now()
	err = client.Do(request, response)
	elapsed := boomer.Now() - startTime

	if err != nil {
		log.Fatalf("%v\n", err)
		boomer.RecordFailure(taskName, "error", 0.0, err.Error())
	} else {
		boomer.RecordSuccess(taskName, strconv.Itoa(response.StatusCode()), elapsed, int64(response.Header.ContentLength()))
	}

	if rand.Intn(100) < (convRate + 1) {
		byteArray := response.Header.Peek("Location")
		location := string(byteArray[:])
		addToClickRing(location)
	}
	fasthttp.ReleaseRequest(request)
}
func addToClickRing(location string) {
	ur, err := url.ParseRequestURI(location)
	if err != nil {
		log.Println(`Unable to get ClickId from url: %s`, location)
		return
	}
	q := ur.Query()
	clickId := q.Get("click_id")
	if clickId == "" {
		return
	}
	ClickIds.Next()
	ClickIds.Value = clickId

}
