package load

import (
	"fmt"
	"log"
	"os"

	"github.com/dzwiedz90/go-loadenvconf/loadenvconf"
	"github.com/dzwiedz90/helpdesk-service-users/config"
	"github.com/dzwiedz90/helpdesk-service-users/logs"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

func LoadEnvConfig(cfg *config.Config) *logs.Logger {
	_, err := loadenvconf.LoadEnvConfig(".env", cfg)
	if err != nil {
		message := fmt.Sprintf("Could not load config: %v", err)
		fmt.Println(message)
		log.Fatal(message)
	}

	logger := logs.NewLoggers(cfg)
	evaluateLogFilesSize(&logger)
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	cfg.ErrorLog = errorLog

	logger.InfoLogger("Config loaded!")

	return &logger
}

func evaluateLogFilesSize(logger *logs.Logger) error {
	maxFileSize := int64(20 * 1024 * 1024)

	consoleLog := "logs/console.log"
	errorLog := "logs/error.log"

	// Check console.log
	fileInfo, err := os.Stat(consoleLog)
	if err != nil {
		logger.ErrorLogger(fmt.Sprintf("Error when getting file information: %v", err))
		return err
	}
	fileSize := fileInfo.Size()

	// Check error.log
	fileInfo2, err := os.Stat(errorLog)
	if err != nil {
		logger.ErrorLogger(fmt.Sprintf("Error when getting file information: %v", err))
		return err
	}
	fileSize2 := fileInfo2.Size()

	// Check if file size greater than 20MB and prune it if so
	if fileSize > maxFileSize {
		logger.InfoLogger("console.log over 20MB - cleaning")
		file, err := os.Create(consoleLog)
		if err != nil {
			logger.ErrorLogger(fmt.Sprintf("Error during console.log cleaning: %v", err))
			panic(err)
		}
		defer file.Close()
	}
	if fileSize2 > maxFileSize {
		logger.InfoLogger("error.log over 20MB - cleaning")
		file2, err := os.Create(errorLog)
		if err != nil {
			logger.ErrorLogger(fmt.Sprintf("Error during error.log cleaning: %v", err))
			panic(err)
		}
		defer file2.Close()
	}

	return nil
}
