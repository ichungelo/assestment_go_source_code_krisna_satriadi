package model

import "time"

type Item struct {
	Id         *int       `json:"id" gorm:"primaryKey;autoIncrement"`
	IsDelete   *bool      `json:"isDelete" gorm:"default:false"`
	CreatedAt  *time.Time `json:"createdAt" gorm:"autoCreateTime" sql:"type:timestamptz"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"autoUpdateTime" sql:"type:timestamptz"`
	DeletedAt  *time.Time `json:"deletedAt"`
	Name       *string    `json:"name"`
	UnitPrice  *int       `json:"unitPrice"`
	ItemTypeId *int       `json:"itemTypeId"`
	ItemType   *ItemType  `json:"itemType" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Invoices   []Invoice  `gorm:"many2many:quantities;foreignKey:Id;joinForeignKey:ItemId;References:Id;joinReferences:InvoiceId" json:"invoices"`
}

// ! Request
type RequestCreateItem struct {
	Name       string `json:"name" validate:"required"`
	ItemTypeId int    `json:"itemTypeId" validate:"required"`
	UnitPrice  int    `json:"unitPrice" validate:"required"`
}

type RequestUpdateItemById struct {
	ItemId     int    `json:"itemId" validate:"required"`
	Name       string `json:"name" validate:"required"`
	ItemTypeId int    `json:"itemTypeId" validate:"required"`
	UnitPrice  int    `json:"unitPrice" validate:"required"`
}

type RequestDeleteItemById struct {
	ItemId int `json:"itemId" validate:"required"`
}

// ! Response
type ResponseGetListItem struct {
	Id     *int    `json:"id"`
	Name   *string `json:"name"`
	Type   *string `json:"type" gorm:"column:ItemType__name"`
}
