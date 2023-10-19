package usecase

import (
	"github.com/ItaloRAmaral/fullcycle-immersion-out23/domain/model"
)

/*
	Aqui é usado o PixKeyRepositoryInterface pois ele é como se fosse uma classe abstrata, o contrato, que foi definida lá no domínio. 
	Sendo assim, quem for implementar o PixKeyRepository deve implementar os métodos que o PixKeyRepositoryInterface do jeito que quiser.
*/
type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}