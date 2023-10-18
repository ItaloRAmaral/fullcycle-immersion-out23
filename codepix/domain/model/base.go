package model

import (
	"time"
	"github.com/asaskevich/govalidator"
)

/*
	Toda vez que este pacote for importado (carregar), a função init() será executada.
	Esta função será executada antes de qualquer outra função deste pacote.
	Esta função irá validar os campos obrigatórios das structs do pacote.
*/
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
} 


/* 
	Base é uma struct que contém os campos comuns entre as entidades
	É como se fosse uma classe abstrata com atributos comuns entre as entidades do sistema.
	E esta classe seria "extendida" pelas entidades do sistema que teriam esses atributos em comum.
	Em Go, não existe herança, mas é possível criar uma struct que contém os campos comuns entre as entidades.
	----

	"valid:"-" significa que não será validado, pois é um campo que pode ser preenchido manualmente ou automaticamente pelo banco de dados.
 */

type Base struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"` // valid:"-" significa que não será validado, pois o campo é preenchido automaticamente pelo banco de dados, ou seja, não é um campo que o usuário preenche
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}