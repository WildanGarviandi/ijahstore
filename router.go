package main

import (
	"github.com/IjahStore/controller"
	"github.com/gin-gonic/gin"
)

// SetRouter is function to define router in app
func SetRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		// stock items
		v1.POST("/items", controller.WriteItems)
		v1.GET("/items", controller.GetItems)
		v1.GET("/item/:sku", controller.GetItem)
		v1.PUT("/items/:sku", controller.EditItems)
		v1.DELETE("/items/:sku", controller.DeleteItems)

		// inbound items
		v1.POST("/log-inbound", controller.WriteItemsInbound)
		v1.GET("/log-inbounds", controller.GetItemsInbound)
		v1.GET("/log-inbound/:sku", controller.GetItemInbound)
		v1.PUT("/log-inbound/:sku", controller.EditItemsInbound)
		v1.DELETE("/log-inbound/:sku", controller.DeleteItemsInbound)

		// outbound items
		v1.POST("/log-outbound", controller.WriteItemsOutbound)
		v1.GET("/log-outbounds", controller.GetItemsOutbound)
		v1.GET("/log-outbound/:sku", controller.GetItemOutbound)
		v1.PUT("/log-outbound/:sku", controller.EditItemsOutbound)
		v1.DELETE("/log-outbound/:sku", controller.DeleteItemsOutbound)

		// reporting
		v1.GET("/generate-report/items-value", controller.GenerateReportValueItems)
		v1.GET("/generate-report/sales", controller.GenerateReportSales)
	}

	return r
}
