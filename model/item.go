package model

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Items struct {
	Id          uint      `gorm:"primaryKey" json:"Id"`
	OrderId     uint      `gorm:"foreignKey:OrderId;references:Id" json:"OrderId"`
	ItemCode    string    `json:"ItemCode" form:"ItemCode" swagger:"description(ItemCode)" valid:"required"`
	Description string    `json:"Description" form:"Description" swagger:"description(Description)" valid:"required"`
	Quantity    int       `json:"Quantity" form:"Quantity" swagger:"description(Quantity)" valid:"required"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"CreatedAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"UpdatedAt"`
	Order       *Orders   `gorm:"foreignKey:OrderId;references:Id"`
}

func (e *Items) BeforeCreate(tx *gorm.DB) (err error) {
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

func (e *Items) BeforeUpdate(tx *gorm.DB) (err error) {
	if err != nil {
		fmt.Println("Error loading Jakarta time zone:", err)
		return
	}
	e.UpdatedAt = time.Now().In(location)
	return nil
}
