package services

import (
	"context"
	"encoding/json"
	"simple-crud-api/models"

	"github.com/go-redis/redis/v8"
)

type ItemService interface {
	GetItems() []models.Item
	GetItem(id string) *models.Item
	CreateItem(item models.Item) models.Item
	UpdateItem(id string, newItem models.Item) *models.Item
	DeleteItem(id string) []models.Item
}

type itemService struct {
	redisClient *redis.Client
}

func NewItemService() ItemService {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Redis database index
	})
	return &itemService{
		redisClient: rdb,
	}
}

func (s *itemService) GetItems() []models.Item {
	var items []models.Item
	ctx := context.Background()
	keys, _ := s.redisClient.Keys(ctx, "item:*").Result()
	for _, key := range keys {
		itemJson, _ := s.redisClient.Get(ctx, key).Result()
		var item models.Item
		json.Unmarshal([]byte(itemJson), &item)
		items = append(items, item)
	}
	return items
}

func (s *itemService) GetItem(id string) *models.Item {
	ctx := context.Background()
	itemJson, err := s.redisClient.Get(ctx, "item:"+id).Result()
	if err == redis.Nil {
		return nil
	}
	var item models.Item
	json.Unmarshal([]byte(itemJson), &item)
	return &item
}

func (s *itemService) CreateItem(item models.Item) models.Item {
	ctx := context.Background()
	itemJson, _ := json.Marshal(item)
	s.redisClient.Set(ctx, "item:"+item.ID, itemJson, 0)
	return item
}

func (s *itemService) UpdateItem(id string, newItem models.Item) *models.Item {
	ctx := context.Background()
	newItem.ID = id
	itemJson, _ := json.Marshal(newItem)
	s.redisClient.Set(ctx, "item:"+id, itemJson, 0)
	return &newItem
}

func (s *itemService) DeleteItem(id string) []models.Item {
	ctx := context.Background()
	s.redisClient.Del(ctx, "item:"+id)
	return s.GetItems()
}
