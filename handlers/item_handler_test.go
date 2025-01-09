package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simple-crud-api/models"
	"simple-crud-api/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemHandler(t *testing.T) {
	mockService := services.NewMockItemService()
	handler := NewItemHandler(mockService)

	// 🍔 Uncomment the following lines to test with real service (i.e. Integration Test)
	// service := services.NewItemService()
	// handler := NewItemHandler(service)

	t.Run("HandleItems POST", func(t *testing.T) {
		// Arrange (for all tests)
		item1 := models.Item{ID: "1", Name: "Item One", Price: "$10"}
		item2 := models.Item{ID: "2", Name: "Item Two", Price: "$20"}

		itemJson1, _ := json.Marshal(item1)
		itemJson2, _ := json.Marshal(item2)

		req1, err := http.NewRequest("POST", "/items", bytes.NewBuffer(itemJson1))
		assert.NoError(t, err)
		rr1 := httptest.NewRecorder()

		req2, err := http.NewRequest("POST", "/items", bytes.NewBuffer(itemJson2))
		assert.NoError(t, err)
		rr2 := httptest.NewRecorder()

		// Act
		handler.HandleItems(rr1, req1)
		handler.HandleItems(rr2, req2)

		// Assert
		assert.Equal(t, http.StatusOK, rr1.Code)
		assert.Equal(t, http.StatusOK, rr2.Code)

		// Verify the response body
		var createdItem1, createdItem2 models.Item
		err = json.NewDecoder(rr1.Body).Decode(&createdItem1)
		assert.NoError(t, err)
		assert.Equal(t, item1, createdItem1)

		err = json.NewDecoder(rr2.Body).Decode(&createdItem2)
		assert.NoError(t, err)
		assert.Equal(t, item2, createdItem2)
	})

	t.Run("HandleItems GET", func(t *testing.T) {
		// Arrange
		req, err := http.NewRequest("GET", "/items", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		// Act
		handler.HandleItems(rr, req)

		// Assert
		assert.Equal(t, http.StatusOK, rr.Code)

		// Verify the response body
		var items []models.Item
		err = json.NewDecoder(rr.Body).Decode(&items)
		assert.NoError(t, err)
		assert.Len(t, items, 2) // 🫙 Bittle test to check if 2 items are returned !
	})

	t.Run("HandleItem GET", func(t *testing.T) {
		// Arrange
		req, err := http.NewRequest("GET", "/items/1", nil)
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		// Act
		handler.HandleItem(rr, req)

		// Assert
		assert.Equal(t, http.StatusOK, rr.Code)

		// Verify the response body
		var item models.Item
		err = json.NewDecoder(rr.Body).Decode(&item)
		assert.NoError(t, err)
		assert.Equal(t, "1", item.ID)
	})

	t.Run("HandleItem PUT", func(t *testing.T) {
		// Arrange
		updatedItem := models.Item{Name: "Updated Item One", Price: "$15"}
		itemJson, _ := json.Marshal(updatedItem)
		req, err := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(itemJson))
		assert.NoError(t, err)
		rr := httptest.NewRecorder()

		// Act
		handler.HandleItem(rr, req)

		// Assert
		assert.Equal(t, http.StatusOK, rr.Code)

		// Verify the response body
		var item models.Item
		err = json.NewDecoder(rr.Body).Decode(&item)
		assert.NoError(t, err)
		assert.Equal(t, "1", item.ID)
		assert.Equal(t, updatedItem.Name, item.Name)
		assert.Equal(t, updatedItem.Price, item.Price)
	})

	t.Run("HandleItem DELETE", func(t *testing.T) {
		// Arrange
		req1, err := http.NewRequest("DELETE", "/items/1", nil)
		assert.NoError(t, err)
		rr1 := httptest.NewRecorder()

		req2, err := http.NewRequest("DELETE", "/items/2", nil)
		assert.NoError(t, err)
		rr2 := httptest.NewRecorder()

		// Act
		handler.HandleItem(rr1, req1)
		handler.HandleItem(rr2, req2)

		// Assert
		assert.Equal(t, http.StatusOK, rr1.Code)
		assert.Equal(t, http.StatusOK, rr2.Code)

		// Verify item 1 is deleted
		reqGet1, err := http.NewRequest("GET", "/items/1", nil)
		assert.NoError(t, err)
		rrGet1 := httptest.NewRecorder()
		handler.HandleItem(rrGet1, reqGet1)
		assert.Equal(t, http.StatusNotFound, rrGet1.Code)

		// Verify item 2 is deleted
		reqGet2, err := http.NewRequest("GET", "/items/2", nil)
		assert.NoError(t, err)
		rrGet2 := httptest.NewRecorder()
		handler.HandleItem(rrGet2, reqGet2)
		assert.Equal(t, http.StatusNotFound, rrGet2.Code)
	})
}
