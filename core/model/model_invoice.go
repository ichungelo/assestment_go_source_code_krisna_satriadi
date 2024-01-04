package model

import "time"

type Invoice struct {
	Id         *int       `json:"id" gorm:"primaryKey;autoIncrement"`
	IsDelete   *bool      `json:"isDelete" gorm:"default:false"`
	CreatedAt  *time.Time `json:"createdAt" gorm:"autoCreateTime" sql:"type:timestamptz"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"autoUpdateTime" sql:"type:timestamptz"`
	DueDate    *time.Time `json:"dueDate"`
	DeletedAt  *time.Time `json:"deletedAt"`
	Subject    *string    `json:"subject"`
	CustomerId *int       `json:"customerId"`
	Customer   *Customer  `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Items      []Item     `gorm:"many2many:quantities;foreignKey:Id;joinForeignKey:InvoiceId;References:Id;joinReferences:ItemId" json:"items"`
}

// ! Request
type RequestCreateInvoice struct {
	Subject    string        `json:"subject" validate:"required"`
	CustomerId int           `json:"customerId" validate:"required"`
	DueDate    time.Time     `json:"dueDate" validate:"required"`
	Items      []RequestItem `json:"items" validate:"required"`
}

type RequestUpdateInvoiceById struct {
	InvoiceId  int           `json:"invoiceId" validate:"required"`
	Subject    string        `json:"subject" validate:"required"`
	CustomerId int           `json:"customerId" validate:"required"`
	Items      []RequestItem `json:"items" validate:"required"`
}

type RequestItem struct {
	ItemId   int `json:"itemId" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
}

type RequestGetInvoiceById struct {
	InvoiceId int `json:"invoiceId" validate:"required"`
}

type RequestGetListInvoice struct {
	IsDelete   *string `json:"isDelete"`
	Limit      *string `json:"limit"`
	Offset     *string `json:"offset"`
	IssueDate  *string `json:"issueDate"`
	Subject    *string `json:"subject"`
	TotalItems *string `json:"totalItems"`
	Customer   *string `json:"customer"`
	DueDate    *string `json:"dueDate"`
	InvoiceId  *string `json:"invoiceId"`
}

type RequestDeleteInvoiceById struct {
	InvoiceId int `json:"invoiceId" validate:"required"`
}

// ! Response
type ResponseGetListInvoiceResult struct {
	InvoiceId    *int       `json:"invoiceId"`
	IssueDate    *time.Time `json:"issueDate"`
	Subject      *string    `json:"subject"`
	TotalItems   *int       `json:"totalItems"`
	CustomerName *string    `json:"customerName"`
	DueDate      *time.Time `json:"dueDate"`
	Status       *string    `json:"status"`
}

type ResponseGetListInvoice struct {
	Total  int                            `json:"total"`
	Count  int                            `json:"count"`
	Start  int                            `json:"start"`
	Result []ResponseGetListInvoiceResult `json:"result"`
}

type ResponseInvoiceById struct {
	Id           *int       `json:"id"`
	IssueDate    *time.Time `json:"issueDate"`
	Subject      *string    `json:"subject"`
	TotalItems   *int       `json:"totalItems"`
	CustomerName *string    `json:"customerName"`
	DueDate      *time.Time `json:"dueDate"`
	Status       *string    `json:"status"`
	Items        []Item     `gorm:"many2many:quantities;foreignKey:Id;joinForeignKey:InvoiceId;References:Id;joinReferences:ItemId" json:"items"`
}
