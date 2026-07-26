package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:50;not null"`
	Password string `gorm:"type:text;not null"`
}
