package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primarykey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

func (p *Post) TableName() string {
	return "posts"
}

func (p Post) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":      p.ID,
		"title":   p.Title,
		"content": p.Content,
	}
}
