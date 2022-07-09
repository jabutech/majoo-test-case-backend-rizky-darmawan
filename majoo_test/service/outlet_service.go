package service

import (
	"github.com/majoo-test/models/domain"
	"github.com/majoo-test/models/web"
	"github.com/majoo-test/repository"
)

type OutletService interface {
	GetListOutletTransactions(outletID web.OutletID, pagination *web.Pagination, dateFrom string, dateTo string) ([]domain.OutletTransactions, error)
	FindByID(outletID web.OutletID) (domain.Outlet, error)
}

type outletService struct {
	repository repository.OutletRepository
}

func NewOutletService(repository repository.OutletRepository) *outletService {
	return &outletService{repository}
}

func (s *outletService) GetListOutletTransactions(outletID web.OutletID, pagination *web.Pagination, dateFrom string, dateTo string) ([]domain.OutletTransactions, error) {
	transactions, err := s.repository.FindOutletTransaction(outletID, pagination, dateFrom, dateTo)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *outletService) FindByID(outletID web.OutletID) (domain.Outlet, error) {
	outlet, err := s.repository.FindByID(outletID)
	if err != nil {
		return outlet, err
	}

	return outlet, nil
}
