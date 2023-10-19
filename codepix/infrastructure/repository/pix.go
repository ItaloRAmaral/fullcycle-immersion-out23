package repository

import (
	"fmt"
	"github.com/ItaloRAmaral/fullcycle-immersion-out23/domain/model"
	"github.com/jinzhu/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

/*
	Pelo fato dos models se auto validarem não é necessario fazer a validação aqui. Basta apenas chamar o método Create do GORM que ele irá persistir o objeto no banco de dados.
	
	Validação é uma regra de negocio, e regras de negocio devem ser implementadas no model.
*/

/*
	(r PixKeyRepositoryDb) é o receiver, ou seja, o objeto que vai receber a função. No caso, o PixKeyRepositoryDb que é uma struct que contém o atributo Db do tipo *gorm.DB

	AddBank(bank *model.Bank) error é a função que será executada pelo receiver. Ela recebe um ponteiro para um objeto do tipo Bank e retorna um erro caso ocorra algum problema na execução da função.
*/
func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

/*
	RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) é a função que será executada pelo receiver. Ela recebe um ponteiro para um objeto do tipo PixKey e retorna um ponteiro para um objeto do tipo PixKey e um erro caso ocorra algum problema na execução da função.
*/
func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

/*
	FindKeyByKind(key string, kind string) (*model.PixKey, error) é a função que será executada pelo receiver. Ela recebe uma string key e uma string kind e retorna um ponteiro para um objeto do tipo PixKey e um erro caso ocorra algum problema na execução da função.
	-------
	-> Preload é uma função do GORM que permite que você carregue dados relacionados a um registro. No caso, estamos carregando os dados da conta e do banco relacionados a chave pix.
	-> First é uma função do GORM que permite que você busque o primeiro registro que atenda a condição passada como parâmetro. No caso, estamos buscando a primeira chave pix que tenha o tipo e a chave passados como parâmetro.
	-> &pixKey é uma referencia a variavel pixKey, ou seja, estamos passando o endereço de memória da variável pixKey para que ela seja alterada com o resultado da busca no banco de dados.
*/
func (r PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKey, nil
}

/*

*/
func (r PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}

	return &account, nil
}

/*

*/
func (r PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	r.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}

	return &bank, nil
}