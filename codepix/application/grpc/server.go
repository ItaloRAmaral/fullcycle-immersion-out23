package grpc

import (
	"fmt"
	"github.com/ItaloRAmaral/fullcycle-immersion-out23/application/grpc/pb"
	"github.com/ItaloRAmaral/fullcycle-immersion-out23/application/usecase"
	"github.com/ItaloRAmaral/fullcycle-immersion-out23/infrastructure/repository"
	"github.com/jinzhu/gorm"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/reflection"
 	"log"
 	"net"
)
/*
	evans -r repl -> cliente interativo do grpc para testar os métodos no bash do docker
*/
func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer) // usado para debugar o gRPC

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)

	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)

	err = grpcServer.Serve(listener)
	
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}