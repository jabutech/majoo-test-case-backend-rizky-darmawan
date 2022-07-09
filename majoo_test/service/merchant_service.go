package service

import (
	"github.com/majoo-test/models/domain"
	"github.com/majoo-test/models/web"
	"github.com/majoo-test/repository"
)

type MerchantService interface {
	GetListMerchantTransactions(merchantID web.MerchantID, pagination *web.Pagination, dateFrom string, dateTo string) ([]domain.MerchantTransactions, error)
	FindByID(merchantID web.MerchantID) (domain.Merchant, error)
	FindByUserID(userID int) (domain.Merchant, error)
}

type merchantService struct {
	repository repository.MerchantRepository
}

func NewMerchantService(repository repository.MerchantRepository) *merchantService {
	return &merchantService{repository}
}

func (s *merchantService) GetListMerchantTransactions(merchantID web.MerchantID, pagination *web.Pagination, dateFrom string, dateTo string) ([]domain.MerchantTransactions, error) {
	transactions, err := s.repository.FindMerchantTransaction(merchantID, pagination, dateFrom, dateTo)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *merchantService) FindByID(merchantID web.MerchantID) (domain.Merchant, error) {
	merchant, err := s.repository.FindByID(merchantID)
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (s *merchantService) FindByUserID(userID int) (domain.Merchant, error) {
	merchant, err := s.repository.FindByUserID(userID)
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}
