package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/IjahStore/utils"
	"github.com/gin-gonic/gin"
)

// GenerateReportValueItems for excel file
func GenerateReportValueItems(c *gin.Context) {
	db, errDB := utils.ConnectDB()

	if errDB != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": errDB})
	}

	type valueItem struct {
		SKU           string  `json:"SKU"`
		ItemsName     string  `json:"items_name"`
		Stock         uint64  `json:"stock"`
		PurchaseValue float64 `json:"purchase_value"`
	}

	var reportItems []valueItem

	// query join table items & log_items_outbonds to get value item report
	db.Table("items").Select("items.sku, items.items_name, items.stock, log_items_outbonds.purchase_value").Joins("left join log_items_outbonds ON items.sku = log_items_outbonds.sku").Scan(&reportItems)

	categories := map[string]string{"A8": "SKU", "B8": "Nama Item", "C8": "Jumlah", "D8": "Rata-Rata Harga Beli", "E8": "Total"}
	var values map[string]string
	values = make(map[string]string)
	var index = 9

	for _, item := range reportItems {
		number := strconv.Itoa(index)
		stock := strconv.FormatUint(item.Stock, 10)
		purchaseValue := strconv.FormatFloat(item.PurchaseValue, 'E', -1, 64)

		values["A"+number] = item.SKU
		values["B"+number] = item.ItemsName
		values["C"+number] = stock
		values["D"+number] = purchaseValue

		index = index + 1
	}

	xlsx := excelize.NewFile()
	for k, v := range categories {
		xlsx.SetCellValue("Laporan Nilai Barang", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Laporan Nilai Barang", k, v)
	}
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./ReportValueItems.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

// GenerateReportSales for excel file
func GenerateReportSales(c *gin.Context) {
	categories := map[string]string{"A10": "ID Pesanan", "B10": "Waktu", "C10": "SKU", "D10": "Nama Barang", "E10": "Jumlah", "F10": "Harga Jual", "G10": "Total", "H10": "Harga Beli", "I10": "Laba"}
	values := map[string]string{"A10": "ID Pesanan", "B10": "Waktu", "C10": "SKU", "D10": "Nama Barang", "E10": "Jumlah", "F10": "Harga Jual", "G10": "Total", "H10": "Harga Beli", "I10": "Laba"}
	xlsx := excelize.NewFile()
	for k, v := range categories {
		xlsx.SetCellValue("Laporan Penjualan", k, v)
	}
	for k, v := range values {
		xlsx.SetCellValue("Laporan Penjualan", k, v)
	}
	// Save xlsx file by the given path.
	err := xlsx.SaveAs("./ReportSales.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
