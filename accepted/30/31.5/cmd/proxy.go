package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	proxyAddress = "localhost:9000"
	serverCount  = 4
)

var (
	hostMask = "localhost:808X"
	turn     = 0
)

func main() {

	http.HandleFunc("/", handleProxy)
	log.Fatalln(http.ListenAndServe(proxyAddress, nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	resp := newRequest(r, body)
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Printf("SERV-%d: %s\n", turn, string(body))

	if turn++; turn >= serverCount {
		turn = 0
	}
}

func newRequest(r *http.Request, body []byte) *http.Response {

	url := r.URL.Path

	req, err := http.NewRequest(r.Method, unmaskHost()+url, bytes.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

func unmaskHost() string {

	X := strings.Index(hostMask, "X")
	return "http://" + hostMask[:X] + strconv.Itoa(turn) + hostMask[X+1:]
}
