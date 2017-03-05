package spider

import (
	"github.com/wzdxt/lianjia-spider/models/ershoufang"
	"net/http"
	"io/ioutil"
	"github.com/wzdxt/lianjia-spider/inspector"
	"github.com/wzdxt/lianjia-spider/models/ershoufang_price"
	"github.com/wzdxt/lianjia-spider/models/xiaoqu"
)

func GetErshoufang(url string) *ershoufang.Ershoufang {
	resp, _ := http.Get(url)
	bytes, _ := (ioutil.ReadAll(resp.Body))
	return inspector.InspectErshoufang(string(bytes))
}

func FetchErshoufangFromUrl(url string) (*ershoufang.Ershoufang, *ershoufang_price.ErshoufangPrice) {
	return inspector.InspectErshoufangFromUrl(url)
}

func FetchXiaoquFromUrl(url string) *xiaoqu.Xiaoqu {
	return inspector.InspectXiaoquFromUrl(url)
}

