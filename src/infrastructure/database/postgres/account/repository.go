package account

import (
	"fmt"

	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/account"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"

	"gorm.io/gorm"
)

type Pagination = shared.Pagination[[]*account.Account]

type AccountRepositoryDb struct {
	Ctx shared.Ctx
	Db  *gorm.DB
}

func NewAccountRepository(ctx shared.Ctx, db *gorm.DB) *AccountRepositoryDb {
	return &AccountRepositoryDb{Db: db, Ctx: ctx}
}

func (repo *AccountRepositoryDb) Create(a *account.Account) (*account.Account, error) {
	model := ToModel(a)
	err := repo.Db.Create(model).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(model), nil
}

func (repo *AccountRepositoryDb) FindByID(id string) (*account.Account, error) {
	var account Account

	err := repo.Db.First(&account, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(&account), nil
}

func (repo *AccountRepositoryDb) FindAll(pagination *Pagination) (*Pagination, error) {
	var f []*Account

	err := repo.Db.Scopes(shared.Paginate(f, pagination, repo.Db)).Find(&f).Error
	pagination.Rows = ToEntityMany(f)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

func (repo *AccountRepositoryDb) FindByOwner(id string) (*account.Account, error) {
	var account Account

	err := repo.Db.First(&account, "owner_id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(&account), nil
}
func (repo *AccountRepositoryDb) FindLastAccountNumber() (string, error) {
	var account Account

	err := repo.Db.Order("created_at DESC").First(&account).Error
	fmt.Println(err)
	if err != nil {
		return "", err
	}

	return account.AccountNumber, nil
}

func (repo *AccountRepositoryDb) Delete(id string) error {
	var account Account

	err := repo.Db.Delete(&account, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *AccountRepositoryDb) Update(a *account.Account) (*account.Account, error) {
	model := ToModel(a)
	err := repo.Db.Save(model).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(model), nil
}
