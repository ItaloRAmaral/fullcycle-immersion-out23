package grpc

import (
	"context"
	"github.com/ItaloRAmaral/fullcycle-immersion-out23/application/usecase"
	"github.com/ItaloRAmaral/fullcycle-immersion-out23/application/grpc/pb"
)

type PixGrpcService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

/*
	Usamos a interface que o gRPC gerou para nós no arquivo "pixkey_grpc.pb.go" para implementar o método RegisterPixKey.
	"pb" é o pacote que o gRPC gerou para nós, e PixKeyRegistration é o tipo que definimos no arquivo "pixkey.proto".

	"in" é o parametro que o gRPC vai receber que é do tipo PixKeyRegistration.

	É como se fosse func(in: PixKeyRegistration) {} no typescript

	Depois de ter criado esses métodos, precisamos gerar um novo serviço no arquivo server.go para que o gRPC possa reconhecer esses métodos. Diretamente ou através de factories.

*/
func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.PixUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error: err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id: key.ID,
		Status: "created",
	}, nil
}

func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.PixUseCase.FindKey(in.Key, in.Kind)
	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:        pixKey.ID,
		Kind:      pixKey.Kind,
		Key:       pixKey.Key,
		Account:   &pb.Account{
			AccountId:     pixKey.AccountID,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankID,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}

func NewPixGrpcService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: usecase,
	}
}