package model

type Quantity struct {
	ItemId    *int     `json:"itemId"`
	Item      *Item    `json:"item" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	InvoiceId *int     `json:"invoiceId"`
	Invoice   *Invoice `json:"invoice" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Count     *int     `json:"count"`
}

//! Request
type RequestDeleteQuantityById struct {
	ItemId    int `json:"itemId" validate:"required"`
	InvoiceId int `json:"invoiceId" validate:"required"`
}
