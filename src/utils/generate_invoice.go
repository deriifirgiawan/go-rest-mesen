package utils

import (
	"fmt"
	"time"
)

func GenerateInvoiceNumber() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()/1000000)
}