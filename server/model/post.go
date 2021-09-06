package model

import (
	"time"
)

type Post struct {
	ID         uint64    `bson:"id"`
	Subject    string    `bson:"subject"`
	Content    string    `bson:"content"`
	Conver     string    `bson:"conver"`
	CategoryID uint32    `bson:"category_id"`
	TagID      []uint32  `bson:"tag_id"`
	IsDeleted  uint8     `bson:"is_deleted"`
	CreatedAt  time.Time `bson:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at"`
}

func (p *Post) Database() string {
	return "lanblog"
}

func (p *Post) Collection() string {
	return "post"
}
