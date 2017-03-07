package spider

import (
	"fmt"
	"github.com/wzdxt/lianjia-spider/inspector"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"regexp"
	"strconv"
	"time"
	"github.com/wzdxt/lianjia-spider/models/chengjiao/repo"
	"github.com/go-sql-driver/mysql"
)

var chengjiaoConflict = 0

func FollowChengjiao() {
	for i := 1; ; i++ {
		if cnt := travelChengjiao(i); cnt == 0 {
			return
		}
		//if chengjiaoConflict > 50 {
		//	return
		//}
		time.Sleep(1 * time.Second)
	}
}

func travelChengjiao(i int) int {
	url := fmt.Sprintf("http://sh.lianjia.com/chengjiao/d%d", i)
	doc := inspector.GetDocFromUrl(url)
	return doc.Find("ul.clinch-list li").Each(func(_ int, sel *goquery.Selection) {
		defer func() {
			if err := recover(); err != nil {
				if err, ok := err.(*mysql.MySQLError); ok && err.Number == 1062 {
					chengjiaoConflict++
				} else {
					panic(err)
				}
			}
		}()
		name := strings.Trim(sel.Find("div.info-panel h2").Text(), " \t\n")
		pageId, _ := sel.Find("div.info-panel h2 a").Attr("key")
		pic, _ := sel.Find("div.pic-panel > div.pic-panel a img").Attr("src")
		qu := sel.Find("div.info-panel div.col-1 div.other div.con a").Eq(0).Text()
		bankuai := sel.Find("div.info-panel div.col-1 div.other div.con a").Eq(1).Text()
		str := sel.Find("div.info-panel div.col-1 div.other div.con").Text()
		louceng, chaoxiang, zhuangxiu := parseMix(str)
		dateStr := sel.Find("div.info-panel div.col-2 div.dealType div.fl div.div-cun").Eq(0).Text()
		date, _ := time.Parse("2006-01-02 MST", dateStr + " CST")
		upStr := sel.Find("div.info-panel div.col-2 div.dealType div.fl div.div-cun").Eq(1).Text()
		unitPrice := fetchNumber(upStr)
		pStr := sel.Find("div.info-panel div.col-2 div.dealType div.fr div.div-cun").Eq(0).Text()
		price := fetchNumber(pStr)
		log.Println(name, pageId, pic, qu, bankuai, louceng, chaoxiang, zhuangxiu, date, unitPrice, price)
		repo.Create(name, pageId, pic, qu, bankuai, louceng, chaoxiang, zhuangxiu, date, unitPrice, price)
	}).Size()
}

func parseMix(text string) (string, string, string) {
	text = strings.Replace(text, "\t", "", -1);
	text = strings.Replace(text, "\n", "", -1);
	text = strings.Replace(text, " ", "", -1);
	strs := regexp.MustCompile("\\|([^\\|]*)").FindAllStringSubmatch(text, -1)
	louceng := ""
	chaoxiang := ""
	zhuangxiu := ""
	i := 0
	if len(strs) > i && strings.HasSuffix(strs[i][1], "层") {
		louceng = strs[i][1]
		i++
	}
	if len(strs) > i && strings.HasPrefix(strs[i][1], "朝") {
		chaoxiang = strs[i][1]
		i++
	}
	if len(strs) > i && strings.HasSuffix(strs[i][1], "装") {
		zhuangxiu = strs[i][1]
		i++
	}
	return louceng, chaoxiang, zhuangxiu
}

func fetchNumber(text string) int {
	str := regexp.MustCompile("\\d+").FindString(text)
	num, _ := strconv.ParseInt(str, 10, 64);
	return int(num)
}

