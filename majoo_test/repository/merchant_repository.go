package repository

import (
	"github.com/majoo-test/models/domain"
	"github.com/majoo-test/models/web"
	"gorm.io/gorm"
)

type MerchantRepository interface {
	FindMerchantTransaction(merchantID web.MerchantID, pagination *web.Pagination, dateFrom string, dateTo string) ([]domain.MerchantTransactions, error)
	FindByID(merchantId web.MerchantID) (domain.Merchant, error)
	FindByUserID(userID int) (domain.Merchant, error)
}

type merchantRepository struct {
	db *gorm.DB
}

func NewRepositoryMerchant(db *gorm.DB) *merchantRepository {
	return &merchantRepository{db}
}

func (r *merchantRepository) FindMerchantTransaction(merchantID web.MerchantID, pagination *web.Pagination, dateFrom string, dateTo string) ([]domain.MerchantTransactions, error) {
	var transactions []domain.MerchantTransactions

	offset := (pagination.Page - 1) * pagination.Limit

	err := r.db.Table("Transactions").Select("sum(bill_total) as omzet, date(created_at) as date, merchant_id").Where("merchant_id = ?", merchantID.ID).Where("date(created_at) BETWEEN ? AND ?", dateFrom, dateTo).Preload("Merchant").Group("date(created_at), merchant_id").Limit(pagination.Limit).Offset(offset).Order("date ASC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *merchantRepository) FindByID(merchantId web.MerchantID) (domain.Merchant, error) {
	var merchant domain.Merchant

	err := r.db.Where("id = ?", merchantId.ID).Find(&merchant).Error
	if err != nil {
		return merchant, err
	}

	return merchant, nil

}
func (r *merchantRepository) FindByUserID(userID int) (domain.Merchant, error) {
	var merchant domain.Merchant

	err := r.db.Where("user_id = ?", userID).Find(&merchant).Error
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}
