package handlers

import (
	"encoding/json"
	"net/http"
	"simple-crud-api/models"
	"simple-crud-api/services"
	"strings"
)

type ItemHandler struct {
	service services.ItemService
}

func NewItemHandler(service services.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) HandleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getItems(w)
	case "POST":
		h.createItem(w, r)
	}
}

func (h *ItemHandler) HandleItem(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getItem(w, r)
	case "PUT":
		h.updateItem(w, r)
	case "DELETE":
		h.deleteItem(w, r)
	}
}

func (h *ItemHandler) getItems(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.GetItems())
}

func (h *ItemHandler) getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/items/")
	item := h.service.GetItem(id)
	if item != nil {
		json.NewEncoder(w).Encode(item)
	} else {
		http.Error(w, "Item not found", http.StatusNotFound)
	}
}

func (h *ItemHandler) createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	createdItem := h.service.CreateItem(item)
	json.NewEncoder(w).Encode(createdItem)
}

func (h *ItemHandler) updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/items/")
	var newItem models.Item
	_ = json.NewDecoder(r.Body).Decode(&newItem)
	updatedItem := h.service.UpdateItem(id, newItem)
	if updatedItem != nil {
		json.NewEncoder(w).Encode(updatedItem)
	} else {
		http.Error(w, "Item not found", http.StatusNotFound)
	}
}

func (h *ItemHandler) deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/items/")
	items := h.service.DeleteItem(id)
	json.NewEncoder(w).Encode(items)
}
