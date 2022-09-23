package customer

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/domain/entities/customer"
	"github.com/AndrewVazzoler/dock-api-rest/src/shared"

	"gorm.io/gorm"
)

type Pagination = shared.Pagination[[]*customer.Customer]

type CustomerRepositoryDb struct {
	Ctx shared.Ctx
	Db  *gorm.DB
}

func NewCustomerRepository(ctx shared.Ctx, db *gorm.DB) *CustomerRepositoryDb {
	return &CustomerRepositoryDb{Db: db, Ctx: ctx}
}

func (repo *CustomerRepositoryDb) Create(a *customer.Customer) (*customer.Customer, error) {
	model := ToModel(a)
	err := repo.Db.Create(model).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(model), nil
}

func (repo *CustomerRepositoryDb) FindByID(id string) (*customer.Customer, error) {
	var customer Customer

	err := repo.Db.First(&customer, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(&customer), nil
}

func (repo *CustomerRepositoryDb) FindAll(pagination *Pagination) (*Pagination, error) {
	var f []*Customer

	err := repo.Db.Scopes(shared.Paginate(f, pagination, repo.Db)).Find(&f).Error
	pagination.Rows = ToEntityMany(f)
	if err != nil {
		return nil, err
	}

	return pagination, nil
}

func (repo *CustomerRepositoryDb) FindByDocument(document string) (*customer.Customer, error) {
	var customer Customer

	err := repo.Db.First(&customer, "document = ?", document).Error

	if err != nil {
		return nil, err
	}

	return ToEntity(&customer), nil
}

func (repo *CustomerRepositoryDb) Delete(id string) error {
	var customer Customer

	err := repo.Db.Delete(&customer, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
