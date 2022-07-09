package web

import (
	"time"

	"github.com/majoo-test/models/domain"
)

type OutletTransactionFormatter struct {
	MerchantName string    `json:"merchant_name"`
	OutletName   string    `json:"outlet_name"`
	Omzet        float64   `json:"omzet"`
	Date         time.Time `json:"date"`
}

func FormatOutletTransaction(outletTrasaction domain.OutletTransactions) OutletTransactionFormatter {
	format := OutletTransactionFormatter{
		MerchantName: outletTrasaction.Merchant.MerchantName,
		OutletName:   outletTrasaction.Outlet.OutletName,
		Omzet:        outletTrasaction.Omzet,
		Date:         outletTrasaction.Date,
	}

	return format
}

func FormatOutletTransactions(outletTrasactions []domain.OutletTransactions) []OutletTransactionFormatter {
	outletTransactionsFormatter := []OutletTransactionFormatter{}

	for _, outletTansaction := range outletTrasactions {
		outletTransactionFormatter := FormatOutletTransaction(outletTansaction)
		outletTransactionsFormatter = append(outletTransactionsFormatter, outletTransactionFormatter)
	}

	return outletTransactionsFormatter
}
