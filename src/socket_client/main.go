package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port \n", os.Args[0])
		os.Exit(1)
	}

	count := 0
	for {
		count += 1
		fmt.Println("try ", count)
		go doConnect()
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func doConnect() {
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	//reponse := make([]byte, 128)
	//_, err = conn.Read(reponse)
	checkError(err)
	fmt.Println(string(result))
	//fmt.Println(string(reponse))
	os.Exit(0)
}
