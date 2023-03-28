package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatalf("error connecting to server: %v\n", err)
	}
	fmt.Println("serv is running")
	con, err := lis.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		line, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message: %v\n", string(line))
		upperLine := strings.ToUpper(string(line))
		if _, err := con.Write([]byte(upperLine)); err != nil {
			log.Fatalln(err)
		}
	}

}
