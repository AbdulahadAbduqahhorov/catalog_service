package main

import (
	"fmt"
	"net"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/config"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/genproto/product_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/service"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
)

func main() {

	cfg := config.Load()
	connP := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)

	db, err := sqlx.Connect("postgres", connP)
	if err != nil {
		panic(err)
	}
	productService := service.NewProductService(db)

	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.Error("error while listening: %v", err)
		return
	}
	service := grpc.NewServer()
	product_service.RegisterProductServiceServer(service, productService)
	if err := service.Serve(lis); err != nil {
		log.Error("error while listening: %v", err)
	}

}
