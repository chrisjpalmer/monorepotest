package pkg1

import (
	"fmt"
	"os"

	"github.com/google/logger"
)

const logPath = "/test.log"

func Pkg1Function() {
	fmt.Println("Hello Pkg1!")
	fmt.Println("Going to log something")
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	defer logger.Init("LoggerExample", true, true, lf).Close()

	logger.Info("I'm about to do something!")
}
