package entity

type BlockLog struct {
	Id            int    `json:"id" xorm:"not null pk bigint"`
	ModuleName    string `json:"moduleName" xorm:"varchar(200)"`
	OperationDesc string `json:"operationDesc" xorm:"varchar(2000)"`
	MethodName    string `json:"methodName" xorm:"varchar(200)"`
	InputDetail   string `json:"inputDetail" xorm:"varchar(2000)"`
	OutputDetail  string `json:"outputDetail" xorm:"text"`
	Status        int    `json:"status" xorm:"tinyint(1)"`
}
