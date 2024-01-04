package model

import "time"

type Customer struct {
	Id        *int       `json:"id" gorm:"primaryKey;autoIncrement"`
	IsDelete  *bool      `json:"isDelete" gorm:"default:false"`
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime" sql:"type:timestamptz"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"autoUpdateTime" sql:"type:timestamptz"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      *string    `json:"name"`
	Address   *string    `json:"address"`
	Invoices  []Invoice  `json:"invoices" gorm:"foreignKey:CustomerId;references:Id"`
}

// ! Request
type RequestCreateCustomer struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type RequestUpdateCustomerById struct {
	CustomerId int `json:"customerId" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Address    string `json:"address" validate:"required"`
}

type RequestDeleteCustomerById struct {
	CustomerId int `json:"customerId" validate:"required"`
}
