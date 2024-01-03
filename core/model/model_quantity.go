package model

type Quantity struct {
	ItemId    *int `json:"itemId" gorm:"primaryKey"`
	InvoiceId *int `json:"invoiceId" gorm:"primaryKey"`
	Count     *int `json:"count"`
}
