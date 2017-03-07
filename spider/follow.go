package spider

import (
	"fmt"
	"github.com/wzdxt/lianjia-spider/inspector"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"regexp"
)

func FollowChengjiao() {
	for i := 1; i < 2; i++ {
		if cnt := travelChengjiao(i); cnt == 0 {
			return
		}
	}
}

func travelChengjiao(i int) int {
	url := fmt.Sprintf("http://sh.lianjia.com/chengjiao/d%d", i)
	doc := inspector.GetDocFromUrl(url)
	return doc.Find("ul.clinch-list li").Each(func(_ int, sel *goquery.Selection) {
		//name := sel.Find("div.info-panel h2").Text()
		//pic, _ :=sel.Find("div.pic-panel > div.pic-panel a img").Attr("src")
		//qu :=sel.Find("div.info-panel div.col-1 div.other div.con a").Nodes[0].Data
		//bankuai :=sel.Find("div.info-panel div.col-1 div.other div.con a").Nodes[1].Data
		str := sel.Find("div.info-panel div.col-1 div.other div.con").Text()
		parseMix(str)
		//louceng
		//chaoxiang
		//zhuangxiu
		//date
		//unitPrice
		//price
	}).Size()
}

func parseMix(text string) (string, string, string) {
	text = strings.Replace(text, "\t", "", -1);
	text = strings.Replace(text, "\n", "", -1);
	text = strings.Replace(text, " ", "", -1);
	strs := regexp.MustCompile("\\|([^\\|]*)").FindAllStringSubmatch(text, 3)
	log.Println(strs)
	return "","",""
}

