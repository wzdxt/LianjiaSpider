package spider

import (
	"github.com/wzdxt/lianjia-spider/models/xiaoqu"
	"github.com/PuerkitoBio/goquery"
	"log"
	xiaoqu_repo "github.com/wzdxt/lianjia-spider/models/xiaoqu/repo"
	ershoufang_repo "github.com/wzdxt/lianjia-spider/models/ershoufang/repo"
	"github.com/go-sql-driver/mysql"
	"sync"
	"time"
	"fmt"
	"strconv"
	"github.com/wzdxt/lianjia-spider/inspector"
)

var wg = sync.WaitGroup{}

func TravelAllXiaoqu() {
	xiaoqus := xiaoqu_repo.AllWithErshoufang()
	wg.Add(len(xiaoqus))
	for _, xiaoqu := range xiaoqus {
		go travelXiaoqu(xiaoqu)
	}
	wg.Wait()
	log.Println("finish")
}

var mutex = sync.Mutex{}
func travelXiaoqu(xiaoqu *xiaoqu.Xiaoqu) {
	for i := 1; ; i++ {
		url := xiaoqu.GetTravelUrl(i)
		doc := inspector.GetDocFromUrl(url)

		if doc.Find("div.main-box div.list-wrap ul.house-lst li h2 a").Each(func(_ int, sel *goquery.Selection) {
			defer func() {
				if err := recover(); err != nil {
					if err, ok := err.(*mysql.MySQLError); ok && err.Number == 1062 {
					} else {
						panic(err)
					}
				}
			}()
			ershoufangPageId, _ := sel.Attr("key")
			ershoufangName := sel.Text()
			log.Printf("found %#v", ershoufangName)
			mutex.Lock()
			defer mutex.Unlock()
			ershoufang_repo.Create(ershoufangPageId, ershoufangName, 0, xiaoqu.PageId, nil)
		}).Size() == 0 {
			break;
		}
	}
	wg.Done()
	log.Println("done")
}

func TranvelXiaoquList() {
	c := make(chan struct{})
	out:
	for i := 0; i < 2000; i += 10 {
		select {
		case <-c:
			break out
		default:
			go travelXiaoquListRange(i + 1, i + 10, c)
		}
		time.Sleep(1 * time.Second)
	}
	log.Println("finish")
}

func travelXiaoquListRange(l, r int, breakChan chan struct{}) {
	for i := l; i <= r; i++ {
		url := fmt.Sprintf("http://sh.lianjia.com/xiaoqu/d%ds13", i)
		doc := inspector.GetDocFromUrl(url)

		if doc.Find("ul#house-lst li div.info-panel").Has("h2 a").Each(func(_ int, sel *goquery.Selection) {
			defer func() {
				if err := recover(); err != nil {
					if err, ok := err.(*mysql.MySQLError); ok && err.Number == 1062 {
					} else {
						panic(err)
					}
				}
			}()
			xiaoquPageId, _ := sel.Find("h2 a").Attr("key")
			xiaoquName := sel.Find("h2 a").Text()
			xiaoquNumber64, _ := strconv.ParseInt(sel.Find("div.square span.num").Text(), 10, 64)
			xiaoquNumber := int(xiaoquNumber64)
			log.Printf("in page %d found %s", i, xiaoquName)
			old := xiaoqu_repo.GetByPageId(xiaoquPageId)
			if (old == nil) {
				xiaoqu := xiaoqu_repo.Create(xiaoquPageId, xiaoquName, xiaoquNumber)
				log.Printf("create new xiaoqu %#v", xiaoqu)
			} else {
				old.Name, old.Number = xiaoquName, xiaoquNumber
				xiaoqu_repo.Save(old)
				log.Printf("udpated xiaoqu %#v", old)
			}
		}).Size() == 0 {
			breakChan <- struct{}{}
			break;
		}
	}
}
