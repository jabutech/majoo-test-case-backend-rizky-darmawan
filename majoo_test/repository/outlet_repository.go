package repository

import (
	"github.com/majoo-test/models/domain"
	"github.com/majoo-test/models/web"
	"gorm.io/gorm"
)

type OutletRepository interface {
	FindOutletTransaction(outletID web.OutletID, pagination *web.Pagination, dateFrom string, dateTo string) ([]domain.OutletTransactions, error)
	FindByID(outletID web.OutletID) (domain.Outlet, error)
}

type outletRepository struct {
	db *gorm.DB
}

func NewRepositoryOutlet(db *gorm.DB) *outletRepository {
	return &outletRepository{db}
}

func (r *outletRepository) FindOutletTransaction(outletID web.OutletID, pagination *web.Pagination, dateFrom string, dateTo string) ([]domain.OutletTransactions, error) {
	var transactions []domain.OutletTransactions

	offset := (pagination.Page - 1) * pagination.Limit

	err := r.db.Table("Transactions").Select("sum(bill_total) as omzet, date(created_at) as date, merchant_id, outlet_id").Where("outlet_id = ?", outletID.ID).Where("date(created_at) BETWEEN ? AND ?", dateFrom, dateTo).Preload("Merchant").Preload("Outlet").Group("date(created_at), merchant_id, outlet_id").Limit(pagination.Limit).Offset(offset).Order("date ASC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *outletRepository) FindByID(outletID web.OutletID) (domain.Outlet, error) {
	var outlet domain.Outlet
	err := r.db.Where("id = ?", outletID.ID).Find(&outlet).Error
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
