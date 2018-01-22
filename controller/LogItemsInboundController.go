package controller

import (
	"net/http"

	"github.com/IjahStore/models"
	"github.com/IjahStore/utils"
	"github.com/gin-gonic/gin"
)

// GetItemInbound is a function to get all data items inbound
func GetItemInbound(c *gin.Context) {
	var logInbound models.LogItemsInbound
	sku := c.Param("sku")

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Where(&models.LogItemsInbound{SKU: sku}).Find(&logInbound)

	if logInbound.SKU == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No single inbound item found!"})
		return
	}

	_logInbound := models.TransformedLogItemsInbound{
		SKU:           logInbound.SKU,
		ItemsName:     logInbound.ItemsName,
		NoInvoice:     logInbound.NoInvoice,
		TotalOrder:    logInbound.TotalOrder,
		TotalAccepted: logInbound.TotalAccepted,
		PurchaseValue: logInbound.PurchaseValue,
		Total:         logInbound.Total,
		Notes:         logInbound.Notes,
		CreatedDate:   logInbound.CreatedAt,
		ModifiedDate:  logInbound.UpdatedAt,
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _logInbound})
}

// GetItemsInbound is a function to get single item
func GetItemsInbound(c *gin.Context) {
	var logInbounds []models.LogItemsInbound
	var _logInbounds []models.TransformedLogItemsInbound

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Find(&logInbounds)
	utils.CloseDB()

	if len(logInbounds) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No log found!"})
		return
	}

	for _, logInbounds := range logInbounds {
		_logInbounds = append(_logInbounds, models.TransformedLogItemsInbound{
			SKU:           logInbounds.SKU,
			ItemsName:     logInbounds.ItemsName,
			NoInvoice:     logInbounds.NoInvoice,
			TotalOrder:    logInbounds.TotalOrder,
			TotalAccepted: logInbounds.TotalAccepted,
			PurchaseValue: logInbounds.PurchaseValue,
			Total:         logInbounds.Total,
			Notes:         logInbounds.Notes,
			CreatedDate:   logInbounds.CreatedAt,
			ModifiedDate:  logInbounds.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _logInbounds})
}

// WriteItemsInbound is a function to record transaction inbound
func WriteItemsInbound(c *gin.Context) {
	var itemInbounds models.LogItemsInbound
	c.BindJSON(&itemInbounds)

	if !isLogInbondItemsCompleted(itemInbounds, c) {
		return
	}

	db, err := utils.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err})
	}

	if err := db.Create(&itemInbounds).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Inbound item not logged!"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Items logged inbound", "SKU": itemInbounds.SKU})
	}
}

// isLogItemsCompleted to check param
func isLogInbondItemsCompleted(logInbound models.LogItemsInbound, c *gin.Context) bool {
	if logInbound.SKU == "" {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "SKU not found"})
		return false
	} else if logInbound.ItemsName == "" {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "item_name not found"})
		return false
	} else if logInbound.NoInvoice == "" {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "no_invoice not found"})
		return false
	}

	return true
}

// EditItemsInbound is a function to edit record inbound with specific SKU
func EditItemsInbound(c *gin.Context) {
	var logInbound models.LogItemsInbound
	sku := c.Param("sku")

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Where(&models.LogItemsInbound{SKU: sku}).Find(&logInbound)

	if logInbound.SKU == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No single inbound item found!"})
		return
	}

	db.Model(&logInbound).Update("items_name", c.PostForm("items_name"))
	db.Model(&logInbound).Update("total_order", c.PostForm("total_order"))
	db.Model(&logInbound).Update("total_accepted", c.PostForm("total_accepted"))
	db.Model(&logInbound).Update("purchase_value", c.PostForm("purchase_value"))
	db.Model(&logInbound).Update("notes", c.PostForm("notes"))

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Log Inbound " + logInbound.SKU + " updated successfully!"})
}

// DeleteItemsInbound is a function to delete record inbound with specific SKU
func DeleteItemsInbound(c *gin.Context) {
	var logInbound models.LogItemsInbound
	sku := c.Param("sku")

	db, err := utils.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "Message": err})
		return
	}

	db.Where(&models.LogItemsInbound{SKU: sku}).Find(&logInbound)

	if logInbound.SKU == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No single inbound item found!"})
		return
	}

	db.Delete(&logInbound)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Log inbound deleted successfully!"})
}
