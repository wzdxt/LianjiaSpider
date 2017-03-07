package inspector

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"time"
)

var limit = make(chan struct{}, 10)

func init() {
	for i := 0; i < 10; i++ {
		limit <- struct{}{}
	}
}

func GetDocFromUrl(url string) *goquery.Document {
	<-limit
	defer func() {
		limit <- struct{}{}
	}()
	var doc *goquery.Document
	for {
		var err error
		doc, err = goquery.NewDocument(url)
		if err != nil {
			log.Println("http get error", err)
			continue
		}
		if doc.Find("p.errorMessageInfo").Size() == 0 {
			break
		}
		log.Println(doc.Find("p.errorMessageInfo").Text())
		log.Println("wait error page ", url)
		time.Sleep(10 * time.Second)
	}
	return doc
}
