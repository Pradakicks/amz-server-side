package Amz

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type AmzMonitor struct {
	Asin       string
	OfferId    string
	Category   string
	url        string
	payload    *strings.Reader
	available  bool
	Client     http.Client
	proxyList  []string
	proxyCount int
}

type postData struct {
	Asin     string
	OfferId  string
	Category string
	Time     int64
}
type Proxy struct {
	ip       string
	port     string
	userAuth string
	userPass string
}

func NewAmzMonitor(asin string, offerId string, category string) {
	m := AmzMonitor{}
	m.Asin = asin
	m.Category = category
	m.OfferId = offerId
	m.url = "https://www.amazon.com/gp/product/features/dp-fast-track/udp-ajax-handler/get-quantity-update-message.html?ie=UTF8"
	m.available = false
	m.Client = http.Client{Timeout: 10 * time.Second}
	file, err := os.Open("proxy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m.proxyList = append(m.proxyList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for {
		m.monitor()
	}
}
func (m *AmzMonitor) getProxy(proxyList []string) string {
	if m.proxyCount+1 == len(proxyList) {
		m.proxyCount = 0
	}
	m.proxyCount++
	return proxyList[m.proxyCount]
}
func (m *AmzMonitor) monitor() {

	currentProxy := m.getProxy(m.proxyList)
	splittedProxy := strings.Split(currentProxy, ":")
	proxy := Proxy{splittedProxy[0], splittedProxy[1], splittedProxy[2], splittedProxy[3]}
	prox1y := fmt.Sprintf("http://%s:%s@%s:%s", proxy.userAuth, proxy.userPass, proxy.ip, proxy.port)
	proxyUrl, err := url.Parse(prox1y)
	if err != nil {
		fmt.Println(err)
		return
	}
	defaultTransport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	m.Client.Transport = defaultTransport

	var currentAvailability bool

	req, err := http.NewRequest("POST", m.url, strings.NewReader(fmt.Sprintf("asin=%s&quantity=1&merchantId=A3KJ3YGUCB8ID4", m.Asin)))
	if err != nil {
		fmt.Println(err)
		return
	}
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
	req.Header.Set("Connection", "close")
	req.Close = true

	res, err := m.Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		test := fmt.Sprintf("Status Code : %d Asin : %s Body Length : %d Availability : %t Current : %t", res.StatusCode, m.Asin, len(string(body)), m.available, currentAvailability)
		fmt.Println(test)
	}()
	if res.StatusCode != 200 {
		fmt.Println(string(body))
	} else {
		fmt.Println(len(string(body)))
		if len(string(body)) > 150 {
			currentAvailability = true
		} else {
			currentAvailability = false
		}

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
