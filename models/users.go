package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"not null;default:null"`
	Email     string `json:"email" gorm:"not null;default:null; unique"`
	Password  string `json:"password" gorm:"not null;default:null"`
	SubDomain string `json:"subDomain" gorm:"not null;default:null; unique"`
}
