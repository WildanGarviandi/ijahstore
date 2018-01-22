package controller

import (
	"fmt"
	"net/http"

	"github.com/IjahStore/models"
	"github.com/IjahStore/utils"
	"github.com/gin-gonic/gin"
)

// GetItemOutbound this if function for get single item outbound
func GetItemOutbound(c *gin.Context) {
	var logOutbond models.LogItemsOutbond
	sku := c.Param("sku")

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Where(&models.LogItemsOutbond{SKU: sku}).Find(&logOutbond)

	if logOutbond.SKU == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No single outbond item found!"})
		return
	}

	_logInbound := models.TransformedLogItemsOutbond{
		SKU:           logOutbond.SKU,
		ItemsName:     logOutbond.ItemsName,
		TotalOrder:    logOutbond.TotalOrder,
		PurchaseValue: logOutbond.PurchaseValue,
		Total:         logOutbond.Total,
		Notes:         logOutbond.Notes,
		CreatedDate:   logOutbond.CreatedAt,
		ModifiedDate:  logOutbond.UpdatedAt,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _logInbound})
}

// GetItemsOutbound this is funtion for get all items outbound
func GetItemsOutbound(c *gin.Context) {
	var itemOutbounds []models.LogItemsOutbond
	var _itemOutbounds []models.TransformedLogItemsOutbond

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Find(&itemOutbounds)
	utils.CloseDB()

	if len(itemOutbounds) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No log outbound found!"})
		return
	}

	for _, outbound := range itemOutbounds {
		_itemOutbounds = append(_itemOutbounds, models.TransformedLogItemsOutbond{
			SKU:           outbound.SKU,
			ItemsName:     outbound.ItemsName,
			TotalOrder:    outbound.TotalOrder,
			PurchaseValue: outbound.PurchaseValue,
			Total:         outbound.Total,
			Notes:         outbound.Notes,
			CreatedDate:   outbound.CreatedAt,
			ModifiedDate:  outbound.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _itemOutbounds})
}

// WriteItemsOutbound this is function to create one record item outbound
func WriteItemsOutbound(c *gin.Context) {
	var logOutbond models.LogItemsOutbond
	c.BindJSON(&logOutbond)

	if !isLogOutbondItemsCompleted(logOutbond, c) {
		return
	}

	db, err := utils.ConnectDB()

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
	}

	if err := db.Create(&logOutbond).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "No log outbond created!"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Log outbond added", "SKU": logOutbond.SKU})
	}
}

// isLogItemsCompleted to check param
func isLogOutbondItemsCompleted(logOutbound models.LogItemsOutbond, c *gin.Context) bool {
	if logOutbound.SKU == "" {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "SKU not found"})
		return false
	} else if logOutbound.ItemsName == "" {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "item_name not found"})
		return false
	}

	return true
}

// EditItemsOutbound this is function to edit item outbound with specific sku
func EditItemsOutbound(c *gin.Context) {
	var logOutbond models.LogItemsOutbond
	sku := c.Param("sku")

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Where(&models.LogItemsOutbond{SKU: sku}).Find(&logOutbond)

	if logOutbond.SKU == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No single outbond item found!"})
		return
	}

	db.Model(&logOutbond).Update("items_name", c.PostForm("items_name"))
	db.Model(&logOutbond).Update("purchase_value", c.PostForm("purchase_value"))
	db.Model(&logOutbond).Update("notes", c.PostForm("notes"))

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Log Outbond " + logOutbond.SKU + " updated successfully!"})
}

// DeleteItemsOutbound tihs is function to delete item outbound with specific sku
func DeleteItemsOutbound(c *gin.Context) {
	var logOutbond models.LogItemsOutbond
	sku := c.Param("sku")

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Where(&models.LogItemsOutbond{SKU: sku}).Find(&logOutbond)

	if logOutbond.SKU == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No single outbond item found!"})
		return
	}

	db.Delete(&logOutbond)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Log outbond deleted successfully!"})
}
