package inspector

import (
	ershoufang_repo "github.com/wzdxt/lianjia-spider/models/ershoufang/repo"
	"log"
	"regexp"
	"strconv"
	"math"
	ershoufang_price_repo "github.com/wzdxt/lianjia-spider/models/ershoufang_price/repo"
	xiaoqu_repo "github.com/wzdxt/lianjia-spider/models/xiaoqu/repo"
	"github.com/wzdxt/lianjia-spider/models/ershoufang"
	"github.com/wzdxt/lianjia-spider/models/ershoufang_price"
	"github.com/wzdxt/lianjia-spider/models/xiaoqu"
	"time"
)

func InspectErshoufang(content string) *ershoufang.Ershoufang {

	//matched, err := regexp.MatchString("")

	return nil
}

func InspectErshoufangFromUrl(url string) (*ershoufang.Ershoufang, *ershoufang_price.ErshoufangPrice) {
	doc := GetDocFromUrl(url)
	html, _ := doc.Html()

	r := regexp.MustCompile("房源编号：(sh\\d+)")
	pageId := r.FindStringSubmatch(html)[1]
	name := doc.Find("div.esf-top h1.main").Text()
	node := (doc.Find("div.content div.houseInfo div.area div.mainInfo"))
	fSize, _ := strconv.ParseFloat(node.Nodes[0].FirstChild.Data, 32)
	size := math.Floor(fSize * 100 + 0.5)
	xiaoquLink, exist := doc.Find("table.aroundInfo a.propertyEllipsis.ml_5").Attr("href")
	var xiaoquPageId string
	if exist {
		xiaoquPageId = regexp.MustCompile("/(\\d+)\\.html").FindStringSubmatch(xiaoquLink)[1]
	} else {
		xiaoquPageId = ""
	}
	var soldDate *time.Time = nil
	//if doc.Find("#album-box div.tag_yixiajia").Size() > 0 || doc.Find("div.pic-cj").Size() > 0 {
	//	tmp := time.Now()
	//	soldDate = &tmp
	//}
	price, _ := strconv.ParseInt(doc.Find("div.content div.houseInfo div.price div.mainInfo").Nodes[0].FirstChild.Data, 10, 64)
	unitPriceStr := doc.Find("table.aroundInfo td:contains(单价) span").Nodes[0].NextSibling.Data
	unitPrice, _ := strconv.ParseInt(regexp.MustCompile("\\d+").FindString(unitPriceStr), 10, 64)
	return ershoufang_repo.New(pageId, name, size, xiaoquPageId, soldDate), ershoufang_price_repo.New(0, 0, int(price), int(unitPrice))
}

func InspectXiaoquFromUrl(url string) *xiaoqu.Xiaoqu {
	doc := GetDocFromUrl(url)
	//html, _ := doc.Html()
	//log.Println(html)

	//log.Println(doc.Find("div.res-top.clear span.t").Text())
	name := doc.Find("div.res-top.clear span.t h1").Text()
	pageId, _ := doc.Find("div#notice_focus").Attr("propertyno")
	numberStr := doc.Find("div#res-nav li a.js_outLink").Text()
	number64, _ := strconv.ParseInt(regexp.MustCompile("（(\\d+)）").FindStringSubmatch(numberStr)[1], 10, 64)

	log.Println(pageId)
	log.Println(name)
	return xiaoqu_repo.New(pageId, name, int(number64))
}

