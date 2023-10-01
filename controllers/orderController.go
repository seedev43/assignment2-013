package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "Ini home")
}

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	// Bind data dari permintaan ke variabel newOrder
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	

		
	// Simpan order ke dalam database
	if err := database.DB.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": newOrder,
	})

}

func GetOrders(ctx *gin.Context) {
	var getOrders []models.Order

	if err := database.DB.Preload("Items").Find(&getOrders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": getOrders,
	})
}

func GetOrderById(ctx *gin.Context) {
	var orderItem models.Order
	orderIDStr := ctx.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ambil data dari database dengan id dari request param client
	if err := database.DB.Preload("Items").First(&orderItem, "id = ?", orderID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orderItem,
	})
}

func UpdateOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("id")
	orderID, _ := strconv.Atoi(orderIDStr)

	// Ambil data order yang akan diperbarui
	var order models.Order
	if err := database.DB.Preload("Items").First(&order, "id = ?", orderID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Bind data dari permintaan ke variabel order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbarui item-item terkait dalam order
	for i := range order.Items {
		itemID := order.Items[i].ID // ID atau primary key item

		// Perbarui item di database berdasarkan ID
		if err := database.DB.Model(&models.Item{}).Where("id = ?", itemID).Updates(&order.Items[i]).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item", "details": err.Error()})
			return
		}
	}

	// Simpan pembaruan pada data order
	if err := database.DB.Save(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func DeleteOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hapus order berdasarkan ID
	if err := database.DB.Delete(&models.Order{}, orderID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
