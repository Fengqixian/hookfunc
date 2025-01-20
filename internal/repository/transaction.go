package repository

import (
	"context"
	"hookfunc/internal/model"
)

type TransactionRepository interface {
	ExistByTransactionId(ctx context.Context, id string) (bool, error)
	FirstByTransactionId(ctx context.Context, id string) (*model.Transaction, error)
	SaveTransaction(ctx context.Context, transaction *model.Transaction) error
}

func NewTransactionRepository(repository *Repository) TransactionRepository {
	return &transactionRepository{
		Repository: repository,
	}
}

type transactionRepository struct {
	*Repository
}

func (r *transactionRepository) ExistByTransactionId(ctx context.Context, id string) (bool, error) {
	var transaction model.Transaction
	if err := r.DB(ctx).Where("transaction_id = ?", id).First(&transaction).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *transactionRepository) SaveTransaction(ctx context.Context, transaction *model.Transaction) error {
	if err := r.DB(ctx).Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) FirstByTransactionId(ctx context.Context, id string) (*model.Transaction, error) {
	var transaction model.Transaction
	if err := r.DB(ctx).Where("transaction_id = ?", id).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}
