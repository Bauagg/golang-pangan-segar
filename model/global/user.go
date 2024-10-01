package modelglobal

import "gorm.io/gorm"

type Role string

var (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type Users struct {
	gorm.Model
	Name    string  `json:"name" binding:"required"`
	Phone   string  `json:"phone" binding:"required"`
	Role    Role    `json:"role" binding:"required" gorm:"type:enum('user', 'admin');default:'user'"`
	Active  bool    `json:"active" gorm:"default:false"`
	Pin     *string `json:"pin" gorm:"default:null"`
	Profile *string `json:"profile" gorm:"default:null"`
	Coin    uint64  `json:"coin" gorm:"default:0"`
	Email   *string `json:"email" gorm:"default:null"`
}

type InputRegister struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Role  Role   `json:"role" binding:"required" gorm:"type:enum('user', 'admin');default:'user'"`
}

type InputPinUser struct {
	Pin    *string `json:"pin" binding:"required"`
	UserId uint64  `json:"user_id" form:"user_id" binding:"required"`
}

type InputLogi struct {
	Phone string `json:"phone" binding:"required"`
}

type InputUpdateProfile struct {
	Name  string  `json:"name" form:"name"`
	Phone string  `json:"phone" form:"phone"`
	Email *string `json:"email" form:"email"`
}
