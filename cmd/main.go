package main

import (
	"fmt"
	"net"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/config"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/genproto/product_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/service"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	c := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)

	//**Db connection
	db, err := sqlx.Connect("postgres", c)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
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
