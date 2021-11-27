package transaction

import (
	"errors"
	repoModels "github.com/stevenfrst/crowdfunding-api/drivers/repository"
	"github.com/stevenfrst/crowdfunding-api/usecase/transaction"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new Transaction interface
func NewTransactionRepository(gormDb *gorm.DB) transaction.TransactionRepoInterface {
	return &TransactionRepository{
		db: gormDb,
	}
}

// GetByID get a transaction via id
func (t TransactionRepository) GetByID(ID int) (transaction.Domain,error) {
	var transaction repoModels.Transaction
	err := t.db.Where("id = ?",ID).Find(&transaction).Error
	//log.Println("REPO",transaction)
	if err != nil || transaction.ID == 0{
		return transaction.ToDomain(),errors.New("Error When Processing TransactionDB/ not found")
	}
	return transaction.ToDomain(),nil
}

// UpdateTransaction methods to update transaction and return to use case domain
func (t TransactionRepository) UpdateTransaction(transaction *transaction.Domain) (*transaction.Domain,error) {
	err := t.db.Save(repoModels.FromDomainTransaction(transaction)).Error

	if err != nil {
		return transaction, err
	}
	return transaction, nil

}

// GetLastTransactionID methods to get last transaction
func (t TransactionRepository) GetLastTransactionID() (int, error) {
	var transaction repoModels.Transaction
	err := t.db.Last(&transaction).Error
	if err != nil {
		return 0,err
	}

	return int(transaction.ID),nil
}

// CreateTransaction methods to create transaction
func (t TransactionRepository) CreateTransaction(transaksi *transaction.Domain) (transaction.Domain,error) {
	//log.Println("Creating transaction",transaksi)
	transactionDomain := repoModels.FromDomainTransaction(transaksi)
	result := t.db.Create(&transactionDomain)

	if result.Error != nil {
		return *transaksi,result.Error
	}
	return *transaksi,nil
}



