package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/dzwiedz90/go-loadenvconf/loadenvconf"
	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"google.golang.org/grpc"

	"github.com/dzwiedz90/helpdesk-service-users/config"
	"github.com/dzwiedz90/helpdesk-service-users/driver"
	"github.com/dzwiedz90/helpdesk-service-users/service"
)

func main() {
	cfg := &config.Config{}
	loadEnvConfig(cfg)
	db := openDB(cfg)
	defer db.SQL.Close()

	addr := cfg.GRPCAddress + ":" + cfg.GRPCPort

	fmt.Printf("Starting application helpdesk-service-users on port %s\n", addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	timeout, err := strconv.Atoi(cfg.Timeout)
	if err != nil {
		log.Fatal("Could not read timeout value form config", err)
	}
	serverTimeout := time.Duration(timeout) * time.Second

	server := grpc.NewServer(grpc.ConnectionTimeout(serverTimeout))
	pb.RegisterUsersServiceServer(server, &service.Server{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
}

func loadEnvConfig(cfg *config.Config) {
	_, err := loadenvconf.LoadEnvConfig(".env", cfg)
	if err != nil {
		log.Fatal("Could not load config", err)
	}
	log.Println("Config loaded!")
}

func openDB(cfg *config.Config) *driver.DB {
	log.Println("Connecting to database...")
	host := "host=" + cfg.DB_HOST + " "
	dbPort := "port=" + cfg.DB_PORT + " "
	dbName := "dbname=" + cfg.DB_NAME + " "
	dbUser := "user=" + cfg.DB_USER + " "
	dbPassword := "password=" + cfg.DB_PASSWORD + " "

	dsn := host + dbPort + dbName + dbUser + dbPassword
	db, err := driver.ConnectSQL(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database. Dying...")
	}
	log.Println("Connected to database!")

	return db
}
