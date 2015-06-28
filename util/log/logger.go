package logger

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
)

type hyenaLogger struct {
	logPath string
}

var instance *hyenaLogger

// GetInstance return singlton instance
func GetInstance() *hyenaLogger {
	if instance == nil {
		hyenaPath := os.Getenv("HYENA_DIR_PATH")
		if hyenaPath == "" {
			usr, err := user.Current()
			if err != nil {
				log.Fatal(err)
			}
			hyenaPath = path.Join(usr.HomeDir, ".config/hyena")
		}
		logPath := path.Join(hyenaPath, "hyena.log")
		instance = &hyenaLogger{logPath: logPath}
	}
	return instance
}

func (mylogger *hyenaLogger) Println(s string) {
	logf, err := os.OpenFile(mylogger.logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error opening file: %v", err)
	}
	defer logf.Close()

	log.SetOutput(logf)
	log.Println(s)
}

func (mylogger *hyenaLogger) Fatalln(e error) {
	logf, err := os.OpenFile(mylogger.logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error opening file: %v", err)
	}
	defer logf.Close()

	log.SetOutput(logf)
	log.Fatalln(e)
}

func (mylogger *hyenaLogger) Printf(format string, e error) {
	logf, err := os.OpenFile(mylogger.logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error opening file: %v", err)
	}
	defer logf.Close()

	log.SetOutput(logf)
	log.Printf(format, e)
}
