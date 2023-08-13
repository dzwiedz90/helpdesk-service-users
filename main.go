package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/dzwiedz90/go-loadenvconf/loadenvconf"
	pb "github.com/dzwiedz90/helpdesk-proto/services/users"
	"google.golang.org/grpc"

	"github.com/dzwiedz90/helpdesk-service-users/config"
	"github.com/dzwiedz90/helpdesk-service-users/driver"
	"github.com/dzwiedz90/helpdesk-service-users/logs"
	"github.com/dzwiedz90/helpdesk-service-users/service"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

func main() {
	cfg := &config.Config{}
	loadEnvConfig(cfg)

	db := openDB(cfg)
	defer db.SQL.Close()

	addr := cfg.GRPCADDRESS + ":" + cfg.GRPCPORT

	logs.InfoLogger(fmt.Sprintf("Starting application helpdesk-service-users on port %s", addr))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		message := fmt.Sprintf("Failed to listen on: %v\n", err)
		logs.ErrorLogger(message)
		log.Fatalf(message)
	}

	logs.InfoLogger(fmt.Sprintf("Listening on %s", addr))

	timeout, err := strconv.Atoi(cfg.TIMEOUT)
	if err != nil {
		message := fmt.Sprintf("Could not read timeout value form config: %v", err)
		logs.ErrorLogger(message)
		log.Fatal(message)
	}
	serverTimeout := time.Duration(timeout) * time.Second

	server := grpc.NewServer(grpc.ConnectionTimeout(serverTimeout))
	pb.RegisterUsersServiceServer(server, &service.Server{
		ServerConfig: &service.ServerConfig{
			DB: db,
		},
	})

	if err = server.Serve(lis); err != nil {
		message := fmt.Sprintf("Failed to serve %v\n", err)
		logs.ErrorLogger(message)
		log.Fatalf(message)
	}
}

func loadEnvConfig(cfg *config.Config) {
	_, err := loadenvconf.LoadEnvConfig(".env", cfg)
	if err != nil {
		message := fmt.Sprintf("Could not load config: %v", err)
		logs.ErrorLogger(message)
		log.Fatal(message)
	}

	evaluateLogFilesSize()
	logs.NewLoggers(cfg)
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	cfg.ErrorLog = errorLog

	logs.InfoLogger("Config loaded!")
}

func openDB(cfg *config.Config) *driver.DB {
	logs.InfoLogger("Connecting to database...")
	host := "host=" + cfg.DB_HOST + " "
	dbPort := "port=" + cfg.DB_PORT + " "
	dbName := "dbname=" + cfg.DB_NAME + " "
	dbUser := "user=" + cfg.DB_USER + " "
	dbPassword := "password=" + cfg.DB_PASSWORD + " "

	dsn := host + dbPort + dbName + dbUser + dbPassword
	db, err := driver.ConnectSQL(dsn)
	if err != nil {
		message := "Cannot connect to database. Dying..."
		logs.ErrorLogger(message)
		log.Fatal(message)
	}
	logs.InfoLogger("Connected to database!")

	return db
}

func evaluateLogFilesSize() error {
	maxFileSize := int64(20 * 1024 * 1024)

	consoleLog := "logs/console.log"
	errorLog := "logs/error.log"

	// Check console.log
	fileInfo, err := os.Stat(consoleLog)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Error when getting file information: %v", err))
		return err
	}
	fileSize := fileInfo.Size()

	// Check error.log
	fileInfo2, err := os.Stat(errorLog)
	if err != nil {
		logs.ErrorLogger(fmt.Sprintf("Error when getting file information: %v", err))
		return err
	}
	fileSize2 := fileInfo2.Size()

	// Check if file size greater than 20MB and prune it if so
	if fileSize > maxFileSize {
		logs.InfoLogger("console.log over 20MB - cleaning")
		file, err := os.Create(consoleLog)
		if err != nil {
			logs.ErrorLogger(fmt.Sprintf("Error during console.log cleaning: %v", err))
			panic(err)
		}
		defer file.Close()
	}
	if fileSize2 > maxFileSize {
		logs.InfoLogger("error.log over 20MB - cleaning")
		file2, err := os.Create(errorLog)
		if err != nil {
			logs.ErrorLogger(fmt.Sprintf("Error during error.log cleaning: %v", err))
			panic(err)
		}
		defer file2.Close()
	}

	return nil
}
