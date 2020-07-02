package main

import (
	"flag"
	"fmt"
	"github.com/alexshvid/go-socks5"
	"log"
	"os"
)

var (
	portNum = flag.Int("port", 8000, "Listen Port for Socks5 Proxy")
)

func main() {

	flag.PrintDefaults()

	listenAddress := fmt.Sprintf("0.0.0.0:%d", *portNum)
	fmt.Printf("Start socks proxy at %s\n", listenAddress)

	fileLog, err := os.OpenFile("socks5.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer fileLog.Close()

	conf := &socks5.Config{
		Logger:      log.New(fileLog, "socks5:", log.LstdFlags),
	}

	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 1080
	if err := server.ListenAndServe("tcp", listenAddress); err != nil {
		panic(err)
	}
}
