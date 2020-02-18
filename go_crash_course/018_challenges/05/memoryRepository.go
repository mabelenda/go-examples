package main

//MemoryRepository memory repository
type MemoryRepository struct {
}

var items = []Item{}

//CreateItem createItem
func (MemoryRepository) CreateItem(newItem Item) {
	items = append(items, newItem)
}

//UpdateItem updateItem
func (MemoryRepository) UpdateItem(updatedItem Item) {
	for i, item := range items {
		if item.ID == updatedItem.ID {
			item.Title = updatedItem.Title
			item.IsDone = updatedItem.IsDone
			items = append(items[:i], item)
		}
	}
}

//GetItems getItems
func (MemoryRepository) GetItems() (items []Item) {
	return items
}

//GetItem getItem
func (MemoryRepository) GetItem(id int) (item Item) {
	var result Item

	for _, item := range items {
		if item.ID == id {
			result = item
			break
		}
	}
	return result
}

//DeleteItem deleteItem
func (MemoryRepository) DeleteItem(id int) {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
}
