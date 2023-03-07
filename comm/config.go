package comm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBX *gorm.DB

type People struct {
	ID    uint32 `gorm:"privateKey,column:id" json:"id,omitempty"`
	Name  string `gorm:"column:name" json:"name,omitempty"`
	Age   uint8  `gorm:"column:age" json:"age,omitempty"`
	Email string `gorm:"column:email" json:"email,omitempty"`
}

func GetPeopleTable() string {
	return "people"
}

type PersonQuery struct {
	AgeStart uint8 `sqlK:"age >"`
	AgeEnd   uint  `sqlK:"age <"`
}
type PersonReq struct {
	PersonQuery
	OrderInfo
	PageInfo
}

func InitDB() {
	dbx, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:12345@tcp(localhost:3306)/test?charset=utf8mb4",
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DBX = dbx
}

func (p *PersonReq) FindPeople() (pl []People) {
	sqlMap := GetSqlBuildMapFromStruct(p)
	cond, val, _ := WhereBuild(sqlMap)
	p.GetPageSql(p.GetOrderSql(DBX.Table(GetPeopleTable()).Where(cond, val...))).Find(&pl)
	return
}
