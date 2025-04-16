package main

import (
	"flag"
	"log"
	"strconv"
)

func main() {
	// Command-line flags for host and port
	port := flag.Int("port", 3333, "Port to accept connections on")
	host := flag.String("host", "127.0.0.1", "Host to bind to")
	flag.Parse()

	address := *host + ":" + strconv.Itoa(*port)

	log.Printf("Server will accept connections on %s...", address)
}
