package handlers

import (
	"assignment-2/database"
	"assignment-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "Ini home")
}

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	// Bind data dari permintaan ke variabel newOrder
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil item_code dari permintaan client
	itemCode := newOrder.Items[0].ItemCode

	var count int64
	database.DB.Model(&models.Item{}).Where("item_code = ?", itemCode).Count(&count)

	if count > 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "item_code sudah ada di database"})
		return
	}
	// Simpan order ke dalam database
	if err := database.DB.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, newOrder)

}
