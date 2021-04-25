package domain

import "time"

type(
	Post struct {
		ID uint `json:"id" gorm:"primaryKey"`
		AuthorId uint `json:"author_id"`
		Title string `json:"title"`
		Description string `json:"description"`
		Content string `json:"content"`
		Date time.Time `json:"date"`
		Author Author `gorm:"foreignKey:AuthorId"`
	}
)