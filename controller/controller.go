package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"orders/config"
	"orders/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetAllOrder godoc
// @Summary Get data order
// @Description Get data all order
// @Tags Order
// @Accept json
// @Produce json
// @Param model.Orders body model.Orders true "Get data all order"
// @Success 200 {object} model.Orders
// @Router /orders [get]
func GetAllOrder(c echo.Context) error {
	db := config.GetDB()

	var order []model.Orders
	db.Preload("Items").Find(&order)

	response := model.DataResponse{
		Data: order,
	}

	fmt.Println("get all data order")

	return c.JSON(http.StatusOK, response)
}

// CreateOrder godoc
// @Summary Create data order
// @Description Create data order
// @Tags Order
// @Accept json
// @Produce json
// @Param model.Orders body model.Orders true "Create data order"
// @Success 200 {object} model.Orders
// @Router /orders [post]
func CreateOrder(ctx echo.Context) error {
	db := config.GetDB()
	order := model.Orders{}
	if err := ctx.Bind(&order); err != nil {
		return err
	}
	err := db.Create(&order)
	if err.Error != nil {
		return ctx.JSON(http.StatusBadRequest, "order tidak berhasil input")
	}
	fmt.Println("order berhasil input")
	respon := model.PutPostResponse{
		Data:    order,
		Message: "Create Data Success",
		Success: true,
	}
	return ctx.JSON(http.StatusOK, respon)
}

// UpdateOrder godoc
// @Summary Update data order
// @Description Update data order
// @Tags Order
// @Accept json
// @Produce json
// @Param model.Orders body model.Orders true "Update data order"
// @Success 200 {object} model.Orders
// @Router /orders/:orderId [put]
func UpdateOrder(ctx echo.Context) error {
	db := config.GetDB()
	id, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	var order model.Orders
	if err := db.First(&order, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}

	var updatedOrder model.Orders
	if err := ctx.Bind(&updatedOrder); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	db.Model(&order.Items).Where("order_id = ?", id).Delete(&order.Items)
	updatedOrder.Id = uint(id)

	if err := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updatedOrder).Error; err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	respon := model.PutPostResponse{
		Data:    updatedOrder,
		Message: "Update Data Success",
		Success: true,
	}
	fmt.Println(updatedOrder)
	return ctx.JSON(http.StatusOK, respon)

}

// DeleteOrder godoc
// @Summary Delete data order
// @Description Delete data order
// @Tags Order
// @Accept json
// @Produce json
// @Param model.Orders body model.Orders true "Delete data order"
// @Success 200 {object} model.Orders
// @Router /orders/:orderId [delete]
func DeleteOrder(ctx echo.Context) error {
	db := config.GetDB()
	order := model.Orders{}
	respon := model.DeleteResponse{
		Status:  http.StatusOK,
		Message: "Delete Data Success",
		Success: true,
	}
	if err := ctx.Bind(&order); err != nil {
		return err
	}
	id := ctx.Param("orderId")
	if err := db.First(&order, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}
	if err := db.Where("order_id = ?", id).Delete(&model.Items{}).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	db.Model(&order.Items).Where("order_id = ?", id).Delete(&order.Items)
	if err := db.Delete(&model.Orders{}, id).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	db.Model(&order).Where("id = ?", id).Delete(&order)
	fmt.Println("order berhasil hapus")
	return ctx.JSON(http.StatusOK, respon)
}
