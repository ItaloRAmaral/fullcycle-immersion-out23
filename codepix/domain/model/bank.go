package model

import (
	"time"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name     string     `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:BankID" valid:"-"` // um banco pode ter várias contas, por isso o slice (array) de Accounts
}

/* 
	Criando um método para o Bank, que recebe um ponteiro para Bank e retorna um erro
	Para criar um método em Go, basta criar uma função que recebe um ponteiro para a struct

	---

	govalidator é uma biblioteca que valida os campos de uma struct
	_ significa que não será utilizado o retorno da função, apenas o erro


*/
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

func NewBank (code string, name string) (*Bank, error) {
	bank := Bank {
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	
	err := bank.isValid()

	if err != nil {
		return nil, err
	}

	return &bank, nil // & retorna o endereço de memória da variável Bank
}