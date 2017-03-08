package chengjiao

import "time"

type Chengjiao struct {
	Id        int64
	Name      string
	PageId    string
	Pic       string
	Qu        string
	Bankuai   string
	Louceng   string
	Chaoxiang string
	Zhuangxiu string
	Date      time.Time
	UnitPrice int
	Price     int
	Room      string
	Size      float64
}

