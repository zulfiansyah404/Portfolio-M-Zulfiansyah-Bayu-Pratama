package models

import (
	"time"
)

type Project struct {
	ID 			int 		`gorm:"primaryKey"`
	Author	 	string 		`gorm:"type:varchar(255);references:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	NameProject string 		`gorm:"type:varchar(255)"`
	StartDate 	time.Time 	`gorm:"type:date"`
	EndDate 	time.Time 	`gorm:"type:date"`
	Duration 	string 		`gorm:"type:varchar(255)"`
	Description string 		`gorm:"type:text"`
	NodeJs     	bool 		`gorm:"type:boolean"`
	ReactJs    	bool 		`gorm:"type:boolean"`
	Golang     	bool 		`gorm:"type:boolean"`
	Java 		bool 		`gorm:"type:boolean"`
	Image 		string 		`gorm:"type:varchar(255)"`
}

