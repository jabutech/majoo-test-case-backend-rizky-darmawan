package domain

import "time"

type OutletTransactions struct {
	Omzet      float64
	Date       time.Time
	MerchantID int64
	OutletID   int64
	Merchant   Merchant
	Outlet     Outlet
}
