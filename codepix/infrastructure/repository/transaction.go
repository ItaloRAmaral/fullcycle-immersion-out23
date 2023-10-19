package repository

import (
	"fmt"
	"github.com/ItaloRAmaral/fullcycle-immersion-out23/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

/*
	(r TransactionRepositoryDb) é o receiver, ou seja, o objeto que vai receber a função. No caso, o TransactionRepositoryDb que é uma struct que contém o atributo Db do tipo *gorm.DB

	Register(transaction *model.Transaction) error é a função que será executada pelo receiver. Ela recebe um ponteiro para um objeto do tipo Transaction e retorna um erro caso ocorra algum problema na execução da função.
*/
func (t *TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := t.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

/*
	"Db.Save" é uma função do GORM que permite que você atualize um registro no banco de dados. No caso, estamos atualizando o registro da transação. 
*/
func (t *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := t.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}
/*
	Find(id string) (*model.Transaction, error) é a função que será executada pelo receiver. Ela recebe um id do tipo string e retorna um ponteiro para um objeto do tipo Transaction e um erro caso ocorra algum problema na execução da função.
	-------
	-> Preload é uma função do GORM que permite que você carregue dados relacionados a um registro. No caso, estamos carregando os dados da conta e do banco relacionados a transação.
	-> First é uma função do GORM que permite que você busque o primeiro registro que atenda a condição passada como parâmetro. No caso, estamos buscando o primeiro registro de transação que tenha o id igual ao id passado como parâmetro.
	-> &transaction é uma referencia a variavel transaction, ou seja, estamos passando o endereço de memória da variável transaction para que ela seja alterada com o resultado da busca no banco de dados.
*/
func (t *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	t.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}