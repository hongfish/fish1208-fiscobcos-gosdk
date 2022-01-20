package entity

type TxBlockTran struct {
	Id           int    `json:"id" xorm:"not null pk bigint"`
	BusinessType string `json:"businessType" xorm:"varchar(200)"`
	BusinessId   string `json:"businessId" xorm:"varchar(1000)"`
	Params       string `json:"params" xorm:"varchar(3000)"`
	Contract     string `json:"contract" xorm:"varchar(100)"`
	Langue       string `json:"langue" xorm:"varchar(100)"`
	Hash         string `json:"hash" xorm:"varchar(256)"`
	Status       int    `json:"status" xorm:"tinyint(1)"`
}
