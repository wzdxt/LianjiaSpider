package repo

import (
	"github.com/wzdxt/lianjia-spider/models/chengjiao"
	"github.com/wzdxt/lianjia-spider/db"
	"time"
)

func Save(chengjiao *chengjiao.Chengjiao) *chengjiao.Chengjiao {
	res, err := db.DBInstance().Exec("insert into chengjiao (" +
		"name, page_id, pic, qu, bankuai, louceng, chaoxiang, zhuangxiu, date, unit_price, price" +
		") values(?,?,?,?,?,?,?,?,?,?,?)",
		chengjiao.Name, chengjiao.PageId, chengjiao.Pic, chengjiao.Qu, chengjiao.Bankuai,
		chengjiao.Louceng, chengjiao.Chaoxiang, chengjiao.Zhuangxiu,
		chengjiao.Date, chengjiao.UnitPrice, chengjiao.Price)
	if err != nil {
		panic(err)
	}
	chengjiao.Id, _ = res.LastInsertId()
	return chengjiao
}

func New(name, pageId, pic, qu, bankuai, louceng, chaoxiang, zhuangxiu string, date time.Time, unitPrice, price     int, room string, size float64) *chengjiao.Chengjiao {
	return &chengjiao.Chengjiao{
		Name  :name,
		PageId:  pageId,
		Pic  :pic,
		Qu  :qu,
		Bankuai  :bankuai,
		Louceng  :louceng,
		Chaoxiang:  chaoxiang,
		Zhuangxiu:  zhuangxiu,
		Date  :date,
		UnitPrice  :unitPrice,
		Price  :price,
		Room:room,
		Size:size,
	}
}

func Create(name, pageId, pic, qu, bankuai, louceng, chaoxiang, zhuangxiu string, date time.Time, unitPrice, price     int, room string, size float64) *chengjiao.Chengjiao {
	return Save(New(name, pageId, pic, qu, bankuai, louceng, chaoxiang, zhuangxiu, date, unitPrice, price, room, size))
}

