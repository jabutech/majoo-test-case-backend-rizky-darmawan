package domain

import "time"

type Outlet struct {
	ID         int64
	MerchantID int64
	OutletName string
	CreatedAt  time.Time
	CreatedBy  int64
	UpdatedAt  time.Time
	UpdatedBy  int64
}

func (Outlet) TableName() string {
	return "Outlets"
}
