package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"github.com/dzwiedz90/helpdesk-service-users/config"
	"github.com/dzwiedz90/helpdesk-service-users/config/load"
	"github.com/dzwiedz90/helpdesk-service-users/driver"
	"github.com/dzwiedz90/helpdesk-service-users/logs"
	"github.com/dzwiedz90/helpdesk-service-users/pkg/core"
	"github.com/dzwiedz90/helpdesk-service-users/service"
	"github.com/dzwiedz90/helpdesk-service-users/service/serviceconfig"
)

func main() {
	cfg := &config.Config{}
	logger := load.LoadEnvConfig(cfg)

	db := openDB(cfg, logger)
	defer db.SQL.Close()

	addr := cfg.GRPCADDRESS + ":" + cfg.GRPCPORT

	logger.InfoLogger(fmt.Sprintf("Starting application helpdesk-service-users on port %s", addr))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		message := fmt.Sprintf("Failed to listen on: %v\n", err)
		logger.ErrorLogger(message)
		log.Fatalf(message)
	}

	logger.InfoLogger(fmt.Sprintf("Listening on %s", addr))

	timeout, err := strconv.Atoi(cfg.TIMEOUT)
	if err != nil {
		message := fmt.Sprintf("Could not read timeout value form config: %v", err)
		logger.ErrorLogger(message)
		log.Fatal(message)
	}
	serverTimeout := time.Duration(timeout) * time.Second

	coreI := core.NewCore()

	server := grpc.NewServer(grpc.ConnectionTimeout(serverTimeout))
	pb.RegisterUsersServiceServer(server, &service.Server{
		Core: &coreI,
		ServerConfig: &serviceconfig.ServerConfig{
			DB:     db,
			Logger: logger,
		},
	})

	if err = server.Serve(lis); err != nil {
		message := fmt.Sprintf("Failed to serve %v\n", err)
		logger.ErrorLogger(message)
		log.Fatalf(message)
	}
}

func openDB(cfg *config.Config, logger *logs.Logger) *driver.DB {
	logger.InfoLogger("Connecting to database...")
	host := "host=" + cfg.DB_HOST + " "
	dbPort := "port=" + cfg.DB_PORT + " "
	dbName := "dbname=" + cfg.DB_NAME + " "
	dbUser := "user=" + cfg.DB_USER + " "
	dbPassword := "password=" + cfg.DB_PASSWORD + " "

	dsn := host + dbPort + dbName + dbUser + dbPassword
	db, err := driver.ConnectSQL(dsn)
	if err != nil {
		message := "Cannot connect to database. Dying..."
		logger.ErrorLogger(message)
		log.Fatal(message)
	}
	logger.InfoLogger("Connected to database!")

	return db
}
