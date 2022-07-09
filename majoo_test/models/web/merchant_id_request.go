package web

type MerchantID struct {
	ID int64 `uri:"merchant_id" binding:"required"`
}
