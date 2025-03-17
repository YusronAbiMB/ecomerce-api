package repository

import (
	"context"

	"github.com/YusronAbi/ecomerce-api/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) models.TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) GetAllTransaction(ctx context.Context, userId int64) ([]*models.Transaction, error) {
	var transaction []*models.Transaction
	err := r.db.Model(&models.Transaction{}).Where("user_id = ?", userId).Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil

}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	if err := r.db.Create(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *TransactionRepository) GetTransactionByID(ctx context.Context, userId int64, id int64) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	if res := r.db.Model(transaction).Where("id = ? AND user_id = ?", id, userId).First(transaction); res.Error != nil {
		return nil, res.Error
	}

	return transaction, nil
}
func (r *TransactionRepository) GetTransactionReport(ctx context.Context) (*models.TransactionReport, error) {
	report := &models.TransactionReport{}
	resUnpaid := r.db.Model(&models.Transaction{}).
		Where("payment = ?", false).
		Select("SUM(sub_total)").Scan(&report.Unpaid)

	if resUnpaid.Error != nil {
		return nil, resUnpaid.Error
	}
	resPaid := r.db.Model(&models.Transaction{}).
		Where("payment = ?", true).
		Select("SUM(sub_total)").Scan(&report.Paid)

	if resPaid.Error != nil {
		return nil, resPaid.Error
	}

	resTotalAmount := r.db.Model(&models.Transaction{}).
		Select("SUM(sub_total)").Scan(&report.Amount)

	if resTotalAmount.Error != nil {
		return nil, resTotalAmount.Error
	}
	var transaction []*models.Transaction
	err := r.db.Model(&models.Transaction{}).Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	report.Transaction = transaction
	return report, nil
}

func (r *TransactionRepository) UpdateTransactionByID(ctx context.Context, userId int64, id int64, data map[string]interface{}) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	res := r.db.Model(&transaction).Where("id = ? AND user_id = ?", id, userId).Updates(data)

	if res.Error != nil {
		return nil, res.Error
	}

	return transaction, nil
}
func (r *TransactionRepository) UpdateTransactionPayment(ctx context.Context, userId int64, id int64) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	res := r.db.Model(&transaction).Where("id = ? AND user_id = ?", id, userId).Find(&transaction)

	if res.Error != nil {
		return nil, res.Error
	}
	resUpdate := r.db.Model(&transaction).Where("id = ? AND user_id = ?", id, userId).Update("payment", true)
	if resUpdate.Error != nil {
		return nil, resUpdate.Error
	}

	return transaction, nil
}

func (r *TransactionRepository) DeleteTransactionByID(ctx context.Context, userId int64, id int64) error {
	transaction := &models.Transaction{}
	res := r.db.Model(&transaction).Where("id = ?", id).Delete(transaction)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
