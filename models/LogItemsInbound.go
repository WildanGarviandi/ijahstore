package models

import (
	"time"
)

// LogItemsInbound is entities form table log_items_inbound on database
type LogItemsInbound struct {
	SKU           string  `gorm:"primary_key;not null"`
	ItemsName     string  `json:"items_name"`
	NoInvoice     string  `gorm:"not null;unique"`
	TotalOrder    int     `json:"total_order"`
	TotalAccepted int     `json:"total_accepted"`
	PurchaseValue float64 `json:"purchase_value"`
	Total         float64 `json:"total"`
	Notes         string  `json:"notes"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TransformedLogItemsInbound is formated for data from table
type TransformedLogItemsInbound struct {
	SKU           string    `json:"SKU"`
	ItemsName     string    `json:"items_name"`
	NoInvoice     string    `json:"no_inv"`
	TotalOrder    int       `json:"total_order"`
	TotalAccepted int       `json:"total_accepted"`
	PurchaseValue float64   `json:"purchase_value"`
	Total         float64   `json:"total"`
	Notes         string    `json:"notes"`
	CreatedDate   time.Time `json:"created_at"`
	ModifiedDate  time.Time `json:"updated_at"`
}
