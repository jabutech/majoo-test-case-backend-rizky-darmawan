package domain

import "time"

type Merchant struct {
	ID           int64
	UserID       int32
	MerchantName string
	CreatedAt    time.Time
	CreatedBy    int64
	UpdatedAt    time.Time
	UpdatedBy    int64
}

func (Merchant) TableName() string {
	return "Merchants"
}
