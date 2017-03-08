package inspector

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"time"
	"strings"
	"errors"
	"net/http"
	"net/url"
)

var limit = make(chan struct{}, 10)

func init() {
	for i := 0; i < 10; i++ {
		limit <- struct{}{}
	}
}

func init() {
	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		log.Printf("%#v", req)
		log.Printf("%#v", via)
		if len(via) == 1 && strings.HasPrefix(via[0].URL.Path, "/ershoufang/") {
			return YichengjiaoError{}
		}
		return nil
	}
}

func GetDocFromUrl(link string) (*goquery.Document, error) {
	<-limit
	defer func() {
		limit <- struct{}{}
	}()
	var doc *goquery.Document
	var err error
	out:
	for {
		doc, err = goquery.NewDocument(link)
		if err != nil {
			switch err.(type) {
			case *url.Error:
				switch err.(*url.Error).Err.(type) {
				case YichengjiaoError:
					log.Println("成交!", link)
					err = err.(*url.Error).Err
					break out
				}
			}
			log.Println("http get error", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if doc.Find("p.errorMessageInfo").Size() == 0 {
			break
		}
		text := doc.Find("p.errorMessageInfo").Text()
		log.Println(text)
		log.Println("wait error page ", link)
		if strings.Contains(text, "稍安勿躁") {
			return nil, errors.New(text)
		}
		time.Sleep(10 * time.Second)
	}
	return doc, err
}

type YichengjiaoError struct{}

func (_ YichengjiaoError)Error() (string) {
	return ""
}
