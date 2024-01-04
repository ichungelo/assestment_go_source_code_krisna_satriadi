package model

import "time"

type ItemType struct {
	Id        *int       `json:"id" gorm:"primaryKey;autoIncrement"`
	IsDelete  *bool      `json:"isDelete" gorm:"default:false"`	
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime" sql:"type:timestamptz"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"autoUpdateTime" sql:"type:timestamptz"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      *string    `json:"name"`
	Items     []Item     `json:"items" gorm:"foreignKey:ItemTypeId;references:Id"`
}

// ! Request
type RequestCreateItemType struct {
	Name string `json:"name" validate:"required"`
}

type RequestUpdateItemTypeById struct {
	ItemTypeId int    `json:"itemTypeId" validate:"required"`
	Name       string `json:"name" validate:"required"`
}

type RequestDeleteItemTypeById struct {
	ItemTypeId int `json:"itemTypeId" validate:"required"`
}

// ! Response
type ResponseGetListItemType struct {
	Id   *int `json:"id"`
	Name *string `json:"name"`
}
