package models

type User struct {
	Username 	string 		`gorm:"type:varchar(255);primaryKey"`
	Password 	string 		`gorm:"type:varchar(255)"`
	Email 		string 		`gorm:"type:varchar(255)"`
	Name 		string 		`gorm:"type:varchar(255)"`
	Experience 	[]string  	`gorm:"type:varchar(255)[]"`
	Year		[]string	`gorm:"type:varchar(255)[]"`
}
