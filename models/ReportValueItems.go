package models

// ReportValueItems is formated data for reporting
type ReportValueItems struct {
	SKU          string  `json:"SKU"`
	ItemsName    string  `json:"items_name"`
	Total        int     `json:"total"`
	AverageValue float64 `json:"average_value"`
	TotalValue   float64 `json:"total_value"`
}
