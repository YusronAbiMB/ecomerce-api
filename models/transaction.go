package models

import "context"

type TransactionReport struct {
	Unpaid      int64       `json:"unpaid"`
	Paid        float64     `json:"paid"`
	Amount      float64     `json:"amount"`
	Transaction interface{} `json:"transaction"`
}

type Transaction struct {
	ID        int64   `json:"id"`
	UserID    int64   `json:"user_id"`
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	SubTotal  float64 `json:"sub_total"`
	Payment   bool    `json:"payment" gorm:"default:false"`
}

type TransactionRepository interface {
	GetAllTransaction(ctx context.Context, userId int64) ([]*Transaction, error)
	GetTransactionByID(ctx context.Context, userId int64, id int64) (*Transaction, error)
	GetTransactionReport(ctx context.Context) (*TransactionReport, error)
	CreateTransaction(ctx context.Context, category *Transaction) (*Transaction, error)
	UpdateTransactionByID(ctx context.Context, userId int64, id int64, data map[string]interface{}) (*Transaction, error)
	UpdateTransactionPayment(ctx context.Context, userId int64, id int64) (*Transaction, error)
	DeleteTransactionByID(ctx context.Context, userId int64, id int64) error
}
