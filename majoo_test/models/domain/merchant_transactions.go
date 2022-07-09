package domain

import "time"

type MerchantTransactions struct {
	Omzet      float64
	Date       time.Time
	MerchantID int64
	Merchant   Merchant
}
