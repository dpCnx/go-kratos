package model

type D struct {
	Id   int    `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
