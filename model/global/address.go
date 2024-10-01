package modelglobal

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	BusinessName string  `json:"busines_name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	City         string  `json:"city" binding:"required"`
	PostalCode   uint64  `json:"postal_code" binding:"required"`
	Province     string  `json:"province" binding:"required"`
	Phone        string  `json:"phone" binding:"required"`
	Name         string  `json:"name" binding:"required"`
	Subdistrict  string  `json:"subdistrict" binding:"required"`
	UserId       uint64  `json:"user_id" binding:"required"`
	Latitude     float64 `gorm:"type:decimal(9,6);not null" json:"latitude" binding:"required"`
	Longitude    float64 `gorm:"type:decimal(9,6);not null" json:"longitude" binding:"required"`
	Status       *string `gorm:"default:null" json:"status"`
}
