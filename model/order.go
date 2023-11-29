package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Orders struct {
	Id           uint      `gorm:"primaryKey" json:"Id"`
	CustomerName string    `json:"CustomerName" form:"CustomerName" swagger:"description(CustomerName)" valid:"required"`
	OrderedAt    time.Time `json:"OrderedAt" form:"OrderedAt" swagger:"description(OrderedAt)" valid:"required"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	Items        []Items   `gorm:"foreignKey:OrderId;references:Id" json:"Items" form:"Items" swagger:"description(Items)" valid:"required"`
}

var location, err = time.LoadLocation("Asia/Jakarta")

func (e *Orders) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(e)
	if err != nil {
		fmt.Println("Error loading Jakarta time zone:", err)
		return
	}
	e.CreatedAt = time.Now().In(location)
	if errCreate != nil {
		fmt.Println(errCreate)
		err = errCreate
		return err
	}

	err = nil
	return
}

func (e *Orders) BeforeUpdate(tx *gorm.DB) (err error) {
	if err != nil {
		fmt.Println("Error loading Jakarta time zone:", err)
		return
	}
	e.UpdatedAt = time.Now().In(location)
	return nil
}
