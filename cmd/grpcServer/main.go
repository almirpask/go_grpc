package main

import (
	"database/sql"
	"net"

	"github.com/almirpask/go_grpc/internal/database"
	"github.com/almirpask/go_grpc/internal/pb"
	"github.com/almirpask/go_grpc/internal/services"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := services.NewCategoryService(*categoryDB)
	grpcServer := grpc.NewServer()

	pb.RegisterCateogryServiceServer(grpcServer, categoryService)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
