package model

type Memo struct {
	Content    string
	DisplayId  uint
	IsWishlist bool
	IsTodo     bool
	Status     int
}
