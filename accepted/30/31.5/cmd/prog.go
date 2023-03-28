package main

import (
	"flag"
	"os"
	"os/signal"
	"sb30_5/api"
	"syscall"
)

func main() {

	userExitSignal := make(chan os.Signal)
	signal.Notify(userExitSignal, syscall.SIGINT)
	go waitExit(userExitSignal)

	host := flag.String("h", ":8080", "server host")
	flag.Parse()
	api.ServerUp(host)
}

func waitExit(userExitSignal chan os.Signal) {

	<-userExitSignal
}
