package Amz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type AmzMonitor struct {
	Asin      string
	OfferId   string
	Category  string
	url       string
	payload   *strings.Reader
	available bool
	Client    http.Client
}

type postData struct {
	Asin     string
	OfferId  string
	Category string
	Time     int64
}

func NewAmzMonitor(asin string, offerId string, category string) {
	m := AmzMonitor{}
	m.Asin = asin
	m.Category = category
	m.OfferId = offerId
	m.url = "https://www.amazon.com/gp/product/features/dp-fast-track/udp-ajax-handler/get-quantity-update-message.html?ie=UTF8"
	m.available = false
	for {
		m.monitor()
	}
}

func (m *AmzMonitor) monitor() {
	var currentAvailability bool

	req, _ := http.NewRequest("POST", m.url, strings.NewReader(fmt.Sprintf("asin=%s&quantity=1&merchantId=A3KJ3YGUCB8ID4", m.Asin)))
	req.Header.Add("authority", "www.amazon.com")
	req.Header.Add("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`)
	req.Header.Add("dnt", "1")
	req.Header.Add("rtt", "50")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36")
	req.Header.Add("content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("accept", "text/html, */*; q=0.01")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("downlink", "10")
	req.Header.Add("ect", "4g")
	req.Header.Add("origin", "https://www.amazon.com")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("accept-language", "en-US,en;q=0.9")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	defer func() {
		test := fmt.Sprintf("Status Code : %d Asin : %s Body Length : %d Availability : %t Current : %t", res.StatusCode, m.Asin, len(string(body)), m.available, currentAvailability)
		fmt.Println(test)
	}()
	fmt.Println(len(string(body)))
	if len(string(body)) > 150 {
		currentAvailability = true
	} else {
		currentAvailability = false
	}

	if currentAvailability && !m.available {
		go m.sendRestock()
	}

	m.available = currentAvailability
}

func (m *AmzMonitor) sendRestock() {
	pay := postData{Time: time.Now().Unix(), Asin: m.Asin, OfferId: m.OfferId, Category: m.Category}
	url := "http://159.203.179.167:3030/"
	payload, err := json.Marshal(pay)
	fmt.Println(string(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Println(res, res.StatusCode)

}
