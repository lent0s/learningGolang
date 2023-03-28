package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

var (
	clients     int = 4
	minReqTime      = 50
	maxReqTime      = 100
	maxCountReq     = 150
	methods         = []string{"PUT", "DELETE", "POST"}
	urls            = []string{"user", "make_friends", "create"}
	ID              = 1
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(clients)

	for port := 0; port < clients; port++ {
		go client(&wg, port)
	}

	wg.Wait()
}

func client(wg *sync.WaitGroup, num int) {

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < maxCountReq; i++ {
		time.Sleep(time.Duration(rand.Intn(maxReqTime-minReqTime)+
			minReqTime) * time.Millisecond)
		//url := fmt.Sprintf("http://localhost:808%d/", num)
		url := fmt.Sprintf("http://localhost:9000/")
		method := methods[rand.Intn(3)]
		body := checkMethod(method, &url, num)
		if body == "" {
			continue
		}

		cmd := exec.Command("curl", "-X", method, url, "-d", body)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		fmt.Printf("\n%-6s %-35s %-36s : ", method, url, body)

		if err := cmd.Run(); err != nil {
			log.Printf("3: %v", err)
		}
	}
	wg.Done()
}

func checkMethod(method string, url *string, num int) (body string) {

	switch method {
	case methods[0]:
		*url += strconv.Itoa(rand.Intn(ID))
		body = fmt.Sprintf(`{"new age": %d}`, rand.Intn(100))
	case methods[1]:
		if rand.Intn(4) != 3 {
			return
		}
		*url += urls[0]
		body = fmt.Sprintf(`{"target_id": %d}`, rand.Intn(ID))
	case methods[2]:
		number := rand.Intn(2) + 1
		*url += urls[number]
		if number == 1 {
			body = fmt.Sprintf(`{"source_id": %d, "target_id": %d}`,
				rand.Intn(ID), rand.Intn(ID))
		} else {
			body = fmt.Sprintf(`{"name": "Serv%d-%d", "age": %d}`,
				num, rand.Intn(500), rand.Intn(100))
			ID++
		}
	}
	return
}
