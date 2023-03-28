package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	progCopies = 4
	filename   = `.\cmd\prog.go`
	flag       = "-h"
)

func main() {

	userExitSignal := make(chan os.Signal)
	signal.Notify(userExitSignal,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGTERM)
	timeToFinish := make(chan bool)
	go waitExit(userExitSignal, timeToFinish)

	shut := make(chan bool)
	go prog(timeToFinish, shut)
	if <-timeToFinish {
		for i := 0; i < progCopies; i++ {
			shut <- true
		}
	}

	exitDuration(3400)
}

func waitExit(userExitSignal chan os.Signal, timeToFinish chan bool) {

	<-userExitSignal
	timeToFinish <- true
}

func prog(timeToFinish chan bool, shut chan bool) {

	wg := sync.WaitGroup{}
	wg.Add(progCopies)
	for number := 0; number < progCopies; number++ {
		go runProg(&wg, number, shut)
	}
	wg.Wait()

	if len(timeToFinish) == 0 {
		timeToFinish <- false
	}
}

func runProg(wg *sync.WaitGroup, number int, shut chan bool) {

	port := fmt.Sprintf(":808%d", number)
	cmd := exec.Command("go", "run", filename, flag, port)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	go closeProg(port, shut)
	if err := cmd.Run(); err != nil {
		if fmt.Sprint(err) != "exit status 1" {
			log.Println("runProg:", err)
		}
	}

	wg.Done()
}

func closeProg(port string, shut chan bool) {

	<-shut
	url := "localhost" + port + "/exit"
	cmd := exec.Command("curl", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println("closeProg:", err)
	}
}

func exitDuration(dur time.Duration) {

	time.Sleep(dur * 10 / 34 * time.Millisecond)
	fmt.Print("exit")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			time.Sleep(dur / 22 * time.Millisecond)
			fmt.Print(".")
		}
		time.Sleep(dur / 22 * time.Millisecond)
		fmt.Print("\b\b\b   \b\b\b")
	}
	fmt.Print("\r\t\b")
}
