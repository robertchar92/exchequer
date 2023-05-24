package models

import "time"

const CashflowTypeDebit = "debit"
const CashflowTypeCredit = "credit"

type Cashflow struct {
	ID                string    `json:"id" groups:"user,admin"`
	BookID            string    `json:"book_id" groups:"user,admin"`
	Description       string    `json:"description" groups:"user,admin"`
	Type              string    `json:"type" groups:"user,admin"`
	Amount            uint64    `json:"amount" groups:"user,admin"`
	Balance           uint64    `json:"balance" groups:"user,admin"`
	RelatedCashflowID string    `json:"releated_cashflow_id" groups:"user,admin"`
	TransactionAt     time.Time `json:"transaction_at" groups:"user,admin"`
	CreatedAt         time.Time `json:"created_at" groups:"user,admin"`
	UpdatedAt         time.Time `json:"-"`
}
