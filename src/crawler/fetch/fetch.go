package fetch

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

// 定时器，减慢采集速度
var rateLimiter = time.Tick(time.Second)

var cookieStr string = "__cm_warden_uid=72a4a73ca844495be77e1373ea17bb78cookie;__cm_warden_upi=MTE5LjEzOS4xOTcuMjE3"

var uid_name string = `'uid_name'.*\:.*\"(.*)\"`
var uid_value string = `'uid_value'.*\:.*\"(.*)\"`
var upi_name string = `'upi_name'.*\:.*\"(.*)\"`
var upi_value string = `'upi_value'.*\:.*\"(.*)\"`

func checkRobot(resp *http.Response) bool {
	contens, _ := ioutil.ReadAll(resp.Body)

	res := regexp.MustCompile(uid_name)
	uid_name_matcher := res.FindSubmatch(contens)

	// 重新把内容放回resp.Body
	buf := bytes.NewBuffer(contens)
	resp.Body = ioutil.NopCloser(buf)

	if len(uid_name_matcher) <= 0 {
		return false
	}

	isUidEmpty := string(uid_name_matcher[0])

	if isUidEmpty == "" {
		return false
	}
	uid_name_str := string(uid_name_matcher[1])

	uid_value_res := regexp.MustCompile(uid_value)
	uid_value_matcher := uid_value_res.FindSubmatch(contens)

	uid_value_str := string(uid_value_matcher[1])

	upi_name_res := regexp.MustCompile(upi_name)
	upi_name_matcher := upi_name_res.FindSubmatch(contens)

	upi_name_str := string(upi_name_matcher[1])

	upi_value_res := regexp.MustCompile(upi_value)
	upi_value_matcher := upi_value_res.FindSubmatch(contens)

	upi_value_str := string(upi_value_matcher[1])

	cookieStr = uid_name_str + "=" + uid_value_str + ";" + upi_name_str + "=" + upi_value_str

	return true
}

func Fetch(url string) (*http.Response, error) {

	<-rateLimiter

	temp, err := http_get(url)

	if checkRobot(temp) {
		fmt.Println("is robet, restart ......")
		temp, err = http_get(url)
	}

	return temp, err
}

func http_get(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.62 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Cookie", cookieStr)
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "www.80s.tw")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	//req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:65.0) Gecko/20100101 Firefox/65.0")
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}

	//e := determineEncoding(resp.Body)
	//
	//utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())

	return resp, err
}

/**
自动编码
*/
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReaderSize(r, 512).Peek(512)
	if err != nil {
		log.Printf("fetcher err %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
