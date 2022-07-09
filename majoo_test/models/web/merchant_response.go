package web

import (
	"time"

	"github.com/majoo-test/models/domain"
)

type MerchantTransactionFormatter struct {
	MerchantName string    `json:"merchant_name"`
	Omzet        float64   `json:"omzet"`
	Date         time.Time `json:"date"`
}

func FormatMerchantTransaction(merchantTrasaction domain.MerchantTransactions) MerchantTransactionFormatter {
	format := MerchantTransactionFormatter{
		MerchantName: merchantTrasaction.Merchant.MerchantName,
		Omzet:        merchantTrasaction.Omzet,
		Date:         merchantTrasaction.Date,
	}

	return format
}

func FormatMerchantTransactions(merchantTrasactions []domain.MerchantTransactions) []MerchantTransactionFormatter {
	merchantTransactionsFormatter := []MerchantTransactionFormatter{}

	for _, merchantTransaction := range merchantTrasactions {
		merchantTransactionFormatter := FormatMerchantTransaction(merchantTransaction)
		merchantTransactionsFormatter = append(merchantTransactionsFormatter, merchantTransactionFormatter)
	}

	return merchantTransactionsFormatter
}
