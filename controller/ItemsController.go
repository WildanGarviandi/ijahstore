package controller

import (
	"fmt"
	"net/http"

	"github.com/IjahStore/models"
	"github.com/IjahStore/utils"
	"github.com/gin-gonic/gin"
)

// GetItem function is to get single item from items
func GetItem(c *gin.Context) {
	var item models.Items
	sku := c.Param("sku")

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Where(&models.Items{SKU: sku}).Find(&item)

	if item.SKU == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No single item found!"})
		return
	}

	_item := models.TransformedItem{SKU: item.SKU, ItemsName: item.ItemsName, Stock: item.Stock}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _item})
}

// GetItems is function to get All items in array
func GetItems(c *gin.Context) {
	var items []models.Items
	var _items []models.TransformedItem

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Find(&items)
	utils.CloseDB()

	if len(items) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Items found!"})
		return
	}

	for _, item := range items {
		_items = append(_items, models.TransformedItem{SKU: item.SKU, ItemsName: item.ItemsName, Stock: item.Stock})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _items})

}

// WriteItems is function to add record items to database
func WriteItems(c *gin.Context) {
	var items models.Items
	c.BindJSON(&items)

	if !isItemsCompleted(items, c) {
		return
	}

	db, err := utils.ConnectDB()

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
	}

	if err := db.Create(&items).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "No item created!"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Items added", "SKU": items.SKU})
	}
}

// isItemsCompleted to check param
func isItemsCompleted(items models.Items, c *gin.Context) bool {
	if items.SKU == "" {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "SKU not found"})
		return false
	} else if items.ItemsName == "" {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "item_name not found"})
		return false
	} else if items.Stock == 0 {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "stock not found"})
		return false
	}

	return true
}

// DeleteItems is function to delete one items on DB
func DeleteItems(c *gin.Context) {
	var item models.Items
	sku := c.Param("sku")

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Where(&models.Items{SKU: sku}).Find(&item)

	if item.SKU == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No single item found!"})
		return
	}

	db.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Item deleted successfully!"})
}

// EditItems is function to edit Items based on SKU
func EditItems(c *gin.Context) {
	var item models.Items
	sku := c.Param("sku")

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Where(&models.Items{SKU: sku}).Find(&item)

	if item.SKU == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No single item found!"})
		return
	}

	db.Model(&item).Update("items_name", c.PostForm("items_name"))

	db.Model(&item).Update("stock", c.PostForm("stock"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Item " + item.SKU + " updated successfully!"})
}
