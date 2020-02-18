package main

//Item items
type Item struct {
	ID     int    `json:"id,omitempty" bson:"id,omitempty"`
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty"`
}

//TypeRepository s
type TypeRepository interface {
	CreateItem(newItem Item)
	UpdateItem(item Item)
	GetItems() []Item
	GetItem(int) Item
	DeleteItem(int)
}
