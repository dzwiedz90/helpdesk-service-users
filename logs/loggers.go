package logs

import (
	"log"
	"os"

	"github.com/dzwiedz90/helpdesk-service-users/config"
)

var cfg *config.Config

func NewLoggers(a *config.Config) {
	cfg = a
}

func InfoLogger(message string) {
	cfg.InfoLog.Println(message)
	logInfoToFile(message)
}

func ErrorLogger(message string) {
	cfg.ErrorLog.Println(message)
	logErrorToFile(message)
}

// func ClientError(status int) {
// 	cfg.InfoLog.Println("Client error with status of", status)
// 	trace := "Client error with status of" + strconv.Itoa(status)
// 	logErrorToFile(trace)
// }

// // Przerobic jako ze serwisy nieoperuja na http i tylko zwracaja bledy wiec printowac z bledow do logow
// func ServerError(err error) {
// 	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
// 	cfg.ErrorLog.Println(trace)
// 	logErrorToFile(trace)
// }

func logInfoToFile(trace string) {
	file, err := os.OpenFile("logs/console.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Println(trace)
}

func logErrorToFile(trace string) {
	file, err := os.OpenFile("logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Println(trace)
}

// logger with data from context
