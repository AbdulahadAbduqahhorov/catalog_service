package main

import (
	"fmt"
	"net"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/config"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/genproto/category_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/genproto/product_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/pkg/logger"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger("catalog_service", cfg.Environment)
	defer logger.Cleanup(log)
	conn := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)

	//**Db connection
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Error("database connection error: %v", logger.Error(err))
		return
	}
	productService := service.NewProductService(log, db)
	categoryService := service.NewCategoryService(db)
	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}
	service := grpc.NewServer()
	product_service.RegisterProductServiceServer(service, productService)
	category_service.RegisterCategoryServiceServer(service, categoryService)

	log.Info("main: server running",
		logger.String("port", cfg.GrpcPort))

	if err := service.Serve(lis); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

}
