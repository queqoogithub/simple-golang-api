package services

import (
	"simple-crud-api/models"
)

type mockItemService struct {
	items []models.Item
}

// func NewMockItemService() ItemService {
// 	return &mockItemService{
// 		items: []models.Item{
// 			{ID: "1", Name: "Mock Item One", Price: "$10"},
// 			{ID: "2", Name: "Mock Item Two", Price: "$20"},
// 		},
// 	}
// }

func NewMockItemService() ItemService {
	return &mockItemService{
		items: []models.Item{},
	}
}

func (s *mockItemService) GetItems() []models.Item {
	return s.items
}

func (s *mockItemService) GetItem(id string) *models.Item {
	for _, item := range s.items {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func (s *mockItemService) CreateItem(item models.Item) models.Item {
	s.items = append(s.items, item)
	return item
}

func (s *mockItemService) UpdateItem(id string, newItem models.Item) *models.Item {
	for index, item := range s.items {
		if item.ID == id {
			newItem.ID = id
			s.items[index] = newItem
			return &newItem
		}
	}
	return nil
}

func (s *mockItemService) DeleteItem(id string) []models.Item {
	for index, item := range s.items {
		if item.ID == id {
			s.items = append(s.items[:index], s.items[index+1:]...)
			break
		}
	}
	return s.items
}
