package grpcserver

import (
	"database/sql"
	"net"

	"github.com/douglasandradeee/go-expert-course/gRPC/internal/database"
	"github.com/douglasandradeee/go-expert-course/gRPC/internal/pb"
	"github.com/douglasandradeee/go-expert-course/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//chamando o categoryDB.
	categoryDb := database.NewCategory()
	// injetando o categoryDB no service e chamando o service.
	categoryService := service.NewCategoryService(*categoryDb)

	//criando servidor gRPC.
	grpcServer := grpc.NewServer()
	//Registrando o nosso service category no servidor gRPC.
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	//necessário utilizar o relfection, pois, iremos utilizar o clientGRPC evans.
	// reflection consegue ler e processar sua própria informação
	reflection.Register(grpcServer)

	//Após isso precisamos abrir uma conexão TCP para camunicar com o gRPC.
	//Porta padrão do gRPC é a :50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	//Iniciando o servidor gRPC.
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
