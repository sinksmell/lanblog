package model

type Pager struct {
	Offset int64
	Limit  int64
}

const(
	NotDeleted = 0
	IsDeleted = 1
)