package model

import "time"

type Tag struct {
	ID        uint32    `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Order     uint32    `json:"order"`
	IsDeleted uint8     `json:"is_deleted"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Tag)Databses()string {
	return "lanblog"
}

func (t *Tag) Collection() string {
	return "tag"
}