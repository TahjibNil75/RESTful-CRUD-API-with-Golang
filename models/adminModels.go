package models

import "time"

type Admin struct {
	Uid             uint      `json:"uid" gorm:"primaryKey:unique"`
	UserName        string    `json:"userName" gorm:"not null" validate:"required, min=3,max=50"`
	Email           string    `json:"email" gorm:"not null:unique" validate:"email, required"`
	Password        string    `json:"password" gorm:"not null" validate:"required"`
	PasswordConfirm string    `json:"passwordConfirm" bindings:"required"`
	IsAdmin         bool      `json:"isAdmin" gorm:"default:true"`
	CreatedAt       time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"not null"`

	// Token and RefreshToken are declared as pointers to allow flexibility
	// and represent the associated token and refresh token for the admin.
	// They can be nil if there are no associated tokens.
	Token        *string `json:"token"`
	RefreshToken *string `json:"refreshToken"`
}

type AdminResponse struct {
	ID        uint      `json:"id,omitempty"`
	UserName  string    `json:"userName,omitempty"`
	Email     string    `json:"email,omitempty"`
	IsAdmin   bool      `json:"admin,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
