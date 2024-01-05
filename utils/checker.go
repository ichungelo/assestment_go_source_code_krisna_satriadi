package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func InvoiceStatusValidator(s string) (*string, error) {

	mapStatus := map[string]string{
		"PAID":   "PAID",
		"UNPAID": "UNPAID",
	}

	formattedStr := strings.ToUpper(clearString(s))

	v, ok := mapStatus[formattedStr]
	if !ok {
		newErr := fmt.Errorf("status not match")
		return nil, newErr
	}

	return &v, nil
}

func InvoiceItemActionValidator(s string) (*string, error) {
	mapAction := map[string]string{
		"ADD":    "ADD",
		"EDIT":   "EDIT",
		"DELETE": "DELETE",
	}

	formattedStr := strings.ToUpper(clearString(s))
	v, ok := mapAction[formattedStr]
	if !ok {
		newErr := fmt.Errorf("status not match")
		return nil, newErr
	}

	return &v, nil
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
