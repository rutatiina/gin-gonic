package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `uri:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name         string `json:"name"`
	Email        string `json:"email"`
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

// for the validation details
// https://github.com/go-playground/validator
type UserNew struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Age          uint8
	Birthday     time.Time `json:"omitempty"`
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

type UserUpdate struct {
	// ID           uint   `uri:"id" binding:"required" gorm:"primarykey"`
	Name         string    `json:"name,omitempty" binding:"omitempty,alphanum"`
	Email        string    `json:"email,omitempty" binding:"omitempty,email"`
	Age          uint8     `json:"age,omitempty"`
	Birthday     time.Time `json:"birthday,omitempty"`
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}
