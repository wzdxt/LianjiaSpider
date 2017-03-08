package main

import (
	"github.com/wzdxt/lianjia-spider/models/ershoufang/repo"
	"github.com/wzdxt/lianjia-spider/monitor/queue/ershoufang"
)

func main() {
	//house, price := inspector.InspectErshoufangFromUrl("http://sh.lianjia.com/ershoufang/sh1059988.html")
	//log.Println(house)
	//log.Println(price)
	//y, m, d := time.Now().Date()
	//log.Printf("%d-%02d-%02d", y, int(m), d)
	//_, err := db.Instance().Exec("insert into ershoufang(sold_date) values(?)", time.Now())
	//panic(err)
	//rows, _ := db.Instance().Query("select now() from ershoufang order by id desc limit 1")
	//rows.Next()
	//var d time.Time
	//rows.Scan(&d)
	//log.Println(d)
	//repo.Create("", "", 0, "", nil)

	//c := make(chan bool)
	//go func() {
	//	i := 0
	//	for {
	//		i += 1
	//		select {
	//		case <-c:
	//			log.Println("stop ", i)
	//			break
	//		default:
	//			log.Println("continue ", i)
	//		}
	//		time.Sleep(1 * time.Millisecond)
	//	}
	//}()
	//time.Sleep(10 * time.Millisecond)
	//c <- true

	//number64 := regexp.MustCompile("（(\\d+)）").FindStringSubmatch("在售二手房（174）")
	//log.Println(number64)

	//t,e  := time.Parse("2006-01-02MST", "2016-12-04CST")
	//log.Printf("%s", t)
	//log.Printf("%#v", e)

	//
	//client := http.DefaultClient
	//client.Jar, _ = cookiejar.New(nil)
	//client.Get("http://sh.lianjia.com")
	////httpurl, _ := url.ParseRequestURI("http://sh.lianjia.com")
	////log.Printf("%#v", httpurl)
	//t := time.Now().UnixNano() / 1000 / 1000
	//resp, _ := client.Get("https://passport.lianjia.com/cas/prelogin/loginTicket?v=" + string(t))
	//defer resp.Body.Close()
	//bs, _ := ioutil.ReadAll(resp.Body)
	//var m = make(map[string]string)
	//json.Unmarshal(bs, &m)
	////log.Printf("%#v", m)
	//resp, _ = client.PostForm("https://passport.lianjia.com/cas/login", url.Values{
	//	"username" : {"18501622774"},
	//	"password" : {"woshi123654"},
	//	"verifycode" : {""},
	//	"service":{"http://sh.lianjia.com"},
	//	"isajax":{"true"},
	//	"lt":{m["data"]},
	//})
	//defer resp.Body.Close()
	////log.Printf("%#v", client.Jar.Cookies(httpurl))
	//resp, _ = client.Get("http://user.sh.lianjia.com/index")
	//defer resp.Body.Close()
	//bs, _ = ioutil.ReadAll(resp.Body)
	////log.Printf("%#v", resp.Header)
	////log.Printf("%s", string(bs))
	////println(err)

	//doc, err := inspector.GetDocFromUrl("http://sh.lianjia.com/ershoufang/sh4431991.html")
	//log.Printf("%#v", doc)
	//log.Printf("%#v", err)

	h := repo.Get(33860)
	ershoufang.Add(h)
	<- ershoufang.Process()
}

