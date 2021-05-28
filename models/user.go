package models

type User struct {
	ID       int    `gorm:"column:pid;size:10"`
	Username string `gorm:"column:puser;size:50"`
	Password string `gorm:"column:ppass;size:255"`
	Status   int    `gorm:"column:pstatus;size:2"`
}

func (u User) TableName() string {
	return "pengguna"
}
