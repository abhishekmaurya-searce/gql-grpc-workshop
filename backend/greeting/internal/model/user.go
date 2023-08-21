package model

type User struct {
	Id    string `gorm:"type:uuid;primaryKey"`
	First string `gorm:"column:Firstname"`
	Last  string `gorm:"column:Lastname"`
}
