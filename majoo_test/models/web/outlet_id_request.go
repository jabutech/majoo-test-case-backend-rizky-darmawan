package web

type OutletID struct {
	ID int64 `uri:"outlet_id" binding:"required"`
}
