package controllers

import (
	"Assignment2/database"
	"Assignment2/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
			"message": "Failed to Create Order",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Order Created Successfully",
		"result":  order,
	})

}

func GetAllOrders(ctx *gin.Context) {
	db := database.GetDB()
	orders := []models.Order{}
	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		panic(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"result":  orders,
	})
}

func UpdateOrder(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}
	order_id := ctx.Param("orderId")
	orders := []models.Order{}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	temp, _ := strconv.Atoi(order_id)
	order.OrderID = uint(temp)

	err := db.Where("order_id = ?", order_id).First(&orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("Order with id %v not found", order_id),
		})
	} else {
		db.Model(&order).Updates(&order)
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"result":  order,
		})
	}

}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()
	orders := []models.Order{}
	orderID := ctx.Param("orderId")
	err := db.Where("order_id = ?", orderID).First(&orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Data Not Found",
			"message": fmt.Sprintf("Order with id %v not found", orderID),
		})
	} else {
		db.Delete(&orders)
		ctx.JSON(http.StatusAccepted, gin.H{
			"success": true,
			"message": fmt.Sprintf("Order with id %v has been successfully delete", orderID),
		})
	}

}
