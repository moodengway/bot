package model

type Match struct {
	ID    uint    `gorm:"column:id;primaryKey;autoIncrement"`
	Host  string  `gorm:"column:host"`
	Guest *string `gorm:"column:guest"`
}
