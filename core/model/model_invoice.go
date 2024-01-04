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
	Status     *string    `json:"status"`
	Customer   *Customer  `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Items      []Item     `gorm:"many2many:quantities;foreignKey:Id;joinForeignKey:InvoiceId;References:Id;joinReferences:ItemId" json:"items"`
}

// ! Request
type RequestCreateInvoice struct {
	Subject    string        `json:"subject" validate:"required"`
	CustomerId int           `json:"customerId" validate:"required"`
	Status     string        `json:"status"  validate:"required"`
	DueDate    time.Time     `json:"dueDate" validate:"required"`
	Items      []RequestItem `json:"items" validate:"required"`
}

type RequestUpdateInvoiceById struct {
	InvoiceId  int           `json:"invoiceId" validate:"required"`
	Subject    string        `json:"subject" validate:"required"`
	CustomerId int           `json:"customerId" validate:"required"`
	Status     string        `json:"status"  validate:"required"`
	DueDate    time.Time     `json:"dueDate" validate:"required"`
	Items      []RequestItem `json:"items" validate:"required"`
}

type RequestItem struct {
	ItemId int `json:"itemId" validate:"required"`
	Count  int `json:"count" validate:"required"`
}

type RequestGetInvoiceById struct {
	InvoiceId int `json:"invoiceId" validate:"required"`
}

type RequestGetListInvoice struct {
	IsDelete   *string `json:"isDelete" query:"isDelete"`
	Limit      *string `json:"limit" query:"limit"`
	Offset     *string `json:"offset" query:"offset"`
	IssueDate  *string `json:"issueDate" query:"issueDate"`
	Subject    *string `json:"subject" query:"subject"`
	TotalItems *string `json:"totalItems" query:"totalItems"`
	Customer   *string `json:"customer" query:"customer"`
	DueDate    *string `json:"dueDate" query:"dueDate"`
	InvoiceId  *string `json:"invoiceId" query:"invoiceId"`
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
	InvoiceId    *int                      `json:"invoiceId,omitempty"`
	IssueDate    *time.Time                `json:"issueDate,omitempty"`
	DueDate      *time.Time                `json:"dueDate,omitempty"`
	Subject      *string                   `json:"subject,omitempty"`
	CustomerName *string                   `json:"customerName,omitempty"`
	CustomerId   *int                      `json:"customerId,omitempty"`
	TotalItems   *int                      `json:"totalItems,omitempty"`
	Status       *string                   `json:"status,omitempty"`
	Items        []ResponseListItemInvoice `json:"items,omitempty" gorm:"-"`
}

type ResponseListItemInvoice struct {
	Id        *int    `json:"id"`
	Name      *string `json:"name"`
	UnitPrice *int    `json:"unitPrice"`
	Quantity  *int    `json:"quantity"`
}
