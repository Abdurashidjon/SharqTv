package main

import (
	"fmt"
	"net"

	"bitbucket.org/udevs/sharqtv_go_user_service/config"
	"bitbucket.org/udevs/sharqtv_go_user_service/genproto/user_service"
	"bitbucket.org/udevs/sharqtv_go_user_service/pkg/logger"
	"bitbucket.org/udevs/sharqtv_go_user_service/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "sharqtv_go_user_service")
	defer logger.Cleanup(log)

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		"disable",
	)
	// conStr = `host=localhost port=5432 user=postgres password=20072003 dbname=user_service sslmode=disable`
	db, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		log.Error("error while connecting database", logger.Error(err))
		return
	}

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

	rolesService := service.NewRoleService(db, log)

	s := grpc.NewServer()
	reflection.Register(s)

	user_service.RegisterRolesServiceServer(s, rolesService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
	}
}
