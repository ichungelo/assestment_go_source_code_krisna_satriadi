package utilerrors

import "net/http"

type ErrorApi struct {
	Code           Code   `json:"code,omitempty"`
	Err            error  `json:"error,omitempty"`
	Message        string `json:"message,omitempty"`
	HttpStatusCode int    `json:"httpStatusCode,omitempty"`
}

type Code string
type ErrorCode struct {
	Code
	Err error
}

const (
	Success                Code = "000"
	ErrValidate            Code = "001"
	ErrParseData           Code = "002"
	ErrFailedUnmarshalJson Code = "003"

	ErrFailedCreateCustomer Code = "004"
	ErrFailedGetCustomer    Code = "005"
	ErrFailedUpdateCustomer Code = "006"
	ErrFailedDeleteCustomer Code = "007"

	ErrFailedCreateInvoice Code = "008"
	ErrFailedGetInvoice    Code = "009"
	ErrFailedUpdateInvoice Code = "010"
	ErrFailedDeleteInvoice Code = "011"

	ErrFailedCreateItemType Code = "012"
	ErrFailedGetItemType    Code = "013"
	ErrFailedUpdateItemType Code = "014"
	ErrFailedDeleteItemType Code = "015"

	ErrFailedCreateItem Code = "016"
	ErrFailedGetItem    Code = "017"
	ErrFailedUpdateItem Code = "018"
	ErrFailedDeleteItem Code = "019"

	ErrFailedDeleteQuantity Code = "020"
	ErrParseDate            Code = "021"

	ErrNotFound            Code = "099"
	ErrInternalServerError Code = "100"
	ErrGeneral             Code = "999"
)

func GetErrorData(errCode ErrorCode) ErrorApi {
	errData := ErrorDataMap[errCode.Code]
	errData.Err = errCode.Err
	return errData
}

var ErrorDataMap = map[Code]ErrorApi{
	Success: {
		Code:           Success,
		Message:        "Success",
		HttpStatusCode: http.StatusOK,
	},
	ErrValidate: {
		Code:           ErrValidate,
		Message:        "Failed to validate",
		HttpStatusCode: http.StatusBadRequest,
	},
	ErrParseData: {
		Code:           ErrParseData,
		Message:        "Failed to parse data",
		HttpStatusCode: http.StatusBadRequest,
	},
	ErrFailedUnmarshalJson: {
		Code:           ErrFailedUnmarshalJson,
		Message:        "Failed to unmarshal json",
		HttpStatusCode: http.StatusUnauthorized,
	},

	ErrFailedCreateCustomer: {
		Code:           ErrFailedCreateCustomer,
		Message:        "Failed to create customer",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedGetCustomer: {
		Code:           ErrFailedGetCustomer,
		Message:        "Failed to get customer",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedUpdateCustomer: {
		Code:           ErrFailedUpdateCustomer,
		Message:        "Failed to update customer",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedDeleteCustomer: {
		Code:           ErrFailedDeleteCustomer,
		Message:        "Failed to delete customer",
		HttpStatusCode: http.StatusInternalServerError,
	},

	ErrFailedCreateInvoice: {
		Code:           ErrFailedCreateInvoice,
		Message:        "Failed to create invoice",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedGetInvoice: {
		Code:           ErrFailedGetInvoice,
		Message:        "Failed to get invoice",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedUpdateInvoice: {
		Code:           ErrFailedUpdateInvoice,
		Message:        "Failed to update invoice",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedDeleteInvoice: {
		Code:           ErrFailedDeleteInvoice,
		Message:        "Failed to delete invoice",
		HttpStatusCode: http.StatusInternalServerError,
	},

	ErrFailedCreateItemType: {
		Code:           ErrFailedCreateItemType,
		Message:        "Failed to create item type",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedGetItemType: {
		Code:           ErrFailedGetItemType,
		Message:        "Failed to get item type",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedUpdateItemType: {
		Code:           ErrFailedUpdateItemType,
		Message:        "Failed to update item type",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedDeleteItemType: {
		Code:           ErrFailedDeleteItemType,
		Message:        "Failed to delete item type",
		HttpStatusCode: http.StatusInternalServerError,
	},

	ErrFailedCreateItem: {
		Code:           ErrFailedCreateItem,
		Message:        "Failed to create item",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedGetItem: {
		Code:           ErrFailedGetItem,
		Message:        "Failed to get item",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedUpdateItem: {
		Code:           ErrFailedUpdateItem,
		Message:        "Failed to update item",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrFailedDeleteItem: {
		Code:           ErrFailedDeleteItem,
		Message:        "Failed to delete item",
		HttpStatusCode: http.StatusInternalServerError,
	},

	ErrFailedDeleteQuantity: {
		Code:           ErrFailedDeleteQuantity,
		Message:        "Failed to delete quantity",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrParseDate: {
		Code:           ErrParseDate,
		Message:        "Failed to parse date",
		HttpStatusCode: http.StatusInternalServerError,
	},

	ErrNotFound: {
		Code:           ErrNotFound,
		Message:        "Page not found",
		HttpStatusCode: http.StatusNotFound,
	},
	ErrInternalServerError: {
		Code:           ErrInternalServerError,
		Message:        "internal server error",
		HttpStatusCode: http.StatusInternalServerError,
	},
	ErrGeneral: {
		Code:           ErrGeneral,
		Message:        "error happen",
		HttpStatusCode: http.StatusInternalServerError,
	},
}
