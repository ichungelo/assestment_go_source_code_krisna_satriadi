package model

type Quantity struct {
	ItemId    *int     `json:"itemId" gorm:"primaryKey"`
	Item      *Item    `json:"item" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	InvoiceId *int     `json:"invoiceId" gorm:"primaryKey"`
	Invoice   *Invoice `json:"invoice" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Count     *int     `json:"count"`
}
