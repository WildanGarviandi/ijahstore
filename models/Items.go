package models

import "time"

// Items is database entites for table Items
type Items struct {
	SKU       string `gorm:"primary_key;not null"`
	ItemsName string `json:"items_name"`
	Stock     uint64 `json:"stock"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TransformedItem used to reformat data to json from table Items
type TransformedItem struct {
	SKU       string `json:"SKU"`
	ItemsName string `json:"items_name"`
	Stock     uint64 `json:"stock"`
}
