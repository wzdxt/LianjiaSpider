package main

import (
	"log"
	"time"
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

	t,e  := time.Parse("2006-01-02MST", "2016-12-04CST")
	log.Printf("%s", t)
	log.Printf("%#v", e)
}

