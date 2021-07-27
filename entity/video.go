package entity

type Video struct {
	BaseModel   BaseModel `gorm:"embedded"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Url         string    `gorm:"type:varchar(255)" json:"url"`
}
