package models

import (
	"time"
)

// LogItemsOutbond is entities for table log_items_outbond
type LogItemsOutbond struct {
	SKU           string  `gorm:"primary_key;not null"`
	ItemsName     string  `json:"items_name"`
	TotalOrder    int     `json:"total_order"`
	PurchaseValue float64 `json:"purchase_value"`
	Total         float64 `json:"total"`
	Notes         string  `json:"notes"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TransformedLogItemsOutbond is formated data from table
type TransformedLogItemsOutbond struct {
	SKU           string    `json:"SKU"`
	ItemsName     string    `json:"items_name"`
	TotalOrder    int       `json:"total_order"`
	PurchaseValue float64   `json:"purchase_value"`
	Total         float64   `json:"total"`
	Notes         string    `json:"notes"`
	CreatedDate   time.Time `json:"created_at"`
	ModifiedDate  time.Time `json:"updated_at"`
}
