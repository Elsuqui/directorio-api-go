package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id        uint64         `gorm:"primary_key:auto_increment" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
