package transaction

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/transactions"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"

	"gorm.io/gorm"
)

type Pagination = shared.Pagination[[]*transactions.Transaction]

type TransactionRepositoryDb struct {
	Ctx shared.Ctx
	Db  *gorm.DB
}

func NewTransactionRepository(ctx shared.Ctx, db *gorm.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{Db: db, Ctx: ctx}
}

func (repo *TransactionRepositoryDb) Create(a *transactions.Transaction) (*transactions.Transaction, error) {
	model := ToModel(a)
	err := repo.Db.Create(model).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(model), nil
}

func (repo *TransactionRepositoryDb) FindByID(id string) (*transactions.Transaction, error) {
	var transactions Transaction

	err := repo.Db.First(&transactions, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(&transactions), nil
}

func (repo *TransactionRepositoryDb) FindAll(pagination *Pagination) (*Pagination, error) {
	var f []*Transaction

	err := repo.Db.Scopes(shared.Paginate(f, pagination, repo.Db)).Find(&f).Error
	pagination.Rows = ToEntityMany(f)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

func (repo *TransactionRepositoryDb) Delete(id string) error {
	var transactions Transaction

	err := repo.Db.Delete(&transactions, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *TransactionRepositoryDb) Update(a *transactions.Transaction) (*transactions.Transaction, error) {
	model := ToModel(a)
	err := repo.Db.Save(model).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(model), nil
}
