package logger

import (
	"log"
	"os"
)

type Logs struct {
	LogInf *log.Logger
	LogErr *log.Logger
	Screen *log.Logger
}

func createLogFile() *os.File {

	f, err := os.OpenFile("server.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func InitLogger() *Logs {

	logfile := createLogFile()
	logs := &Logs{
		LogInf: log.New(logfile, "INF:\t",
			log.Ldate|log.Ltime),
		LogErr: log.New(logfile, "ERROR:\t",
			log.Ldate|log.Ltime|log.Llongfile),
		Screen: log.New(os.Stdout, "", log.LstdFlags),
	}
	return logs
}
