package usecase

import (
	"errors"
	"log"

	"github.com/MelkdeSousa/codepix/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepositoryInterface
	PixRepository         model.PixKeyRepositoryInterface
}

func (transaction *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {
	account, err := transaction.PixRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := transaction.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	newTransaction, err := model.NewTransaction(account, amount, pixKey, description)

	if err != nil {
		return nil, err
	}

	transaction.TransactionRepository.Save(newTransaction)

	if newTransaction.ID == "" {
		return newTransaction, err
	}

	return nil, errors.New("unable to process this transaction")
}

func (transaction *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {
	foundTransaction, err := transaction.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	foundTransaction.Status = model.TransactionConfirmed
	err = transaction.TransactionRepository.Save(foundTransaction)

	if err != nil {
		return nil, err
	}

	return foundTransaction, nil
}

func (transaction *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {
	foundTransaction, err := transaction.TransactionRepository.Find(transactionId)

	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	foundTransaction.Status = model.TransactionCompleted
	err = transaction.TransactionRepository.Save(foundTransaction)

	if err != nil {
		return nil, err
	}

	return foundTransaction, nil
}

func (transaction *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {
	foundTransaction, err := transaction.TransactionRepository.Find(transactionId)

	if err != nil {
		return nil, err
	}

	foundTransaction.Status = model.TransactionError
	foundTransaction.CancelDescription = reason

	err = transaction.TransactionRepository.Save(foundTransaction)

	if err != nil {
		return nil, err
	}

	return foundTransaction, nil
}
