package model

import "time"

type Category struct {
	ID        uint32    `bson:"id"`
	Name      string    `bson:"name"`
	Desc      string    `bson:"desc"`
	Order     uint32    `bson:"order"`
	IsDeleted uint8     `bson:"is_deleted"`
	UpdatedAt time.Time `bson:"updated_at"`
	CreatedAt time.Time `bson:"created_at"`
}

func (t *Category)Database() string {
	return "lanblog"
}

func (t *Category)Collection() string {
	return "category"
}
