package models

type Currency struct {
	ID    int     `gorm:"column:vid;size:10"`
	Code  string  `gorm:"column:vcode;size:3"`
	Name  string  `gorm:"column:vname;size:20"`
	Price float64 `gorm:"column:vprice;scale:2;precision:10"`
}

func (c Currency) TableName() string {
	return "valuta"
}
