package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func InvoiceStatusValidator(s string) (*string, error) {
	var (
		defName *string
		isValid = false
	)

	listStatus := []string{
		"PAID", "UNPAID",
	}

	formattedStr := strings.ToUpper(clearString(s))

	for _, v := range listStatus {
		if formattedStr == v {

			defName = &v
		}
	}

	if !isValid {
		newErr := fmt.Errorf("status not match")
		return nil, newErr
	}

	return defName, nil
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
