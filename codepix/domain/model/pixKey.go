package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

/* 
	PixKeyRepositoryInterface é uma interface que contém os métodos que o PixKeyRepository deve implementar.
	É um "contrato" para quem for implementar o agregado (pixKey, account, bank) no banco de dados saiba quais metodos deve implementar/utilizar.
	É uma forma de garantir que o PixKeyRepository implemente os métodos que o PixKeyRepositoryInterface contém.
*/
type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" gorm:"type:varchar(20)" valid:"notnull"` // tipo de chave pix (CPF ou email)
	Key       string   `json:"key" gorm:"type:varchar(255)" valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
}

/*
	Valida os campos da struct PixKey
	Feito de uma forma "rudmentar"
	Indepentente se é rudmentar ou não, validações, estado e etc devem ser feitos no domínio (no caso, neste model)
*/
func (p *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(p)

	// Essas validações, são regras de negocio
	if p.Kind != "email" && p.Kind != "cpf" { 
		return errors.New("invalid type of key")
	}

	if p.Status != "active" && p.Status != "inactive" {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:      kind,
		Key:       key,
		Account:   account,
		AccountID: account.ID,
		Status:    "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	
	err := pixKey.isValid()

	if err != nil {
		return nil, err
	}
	return &pixKey, nil
}