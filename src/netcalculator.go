package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {

	port := flag.Int("port", 3333, "Port to accept connections on")
	host := flag.String("host", "127.0.0.1", "Host to bind to")
	flag.Parse()

	address := *host + ":" + strconv.Itoa(*port)

	log.Printf("Server will accept connections on %s...", address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("[Error] Could not start server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[Error] Accepting connection failed:", err)
			continue
		}
		go calculatorSession(conn)
	}
}

func calculatorSession(conn net.Conn) {
	defer conn.Close()
	acc := int64(0)
	reader := bufio.NewReader(conn)

	for {
		line, readErr := readCRLF(reader)
		if readErr != nil {
			return
		}

		if line == "" {
			fmt.Fprintf(conn, "%d\r\n", acc)
			acc = 0
			continue
		}

		cmd, num, parseOk := parseInput(line)
		if !parseOk {
			continue
		}

		switch cmd {
		case "ADD":
			acc += num
		case "SUB":
			acc -= num
		case "MUL":
			acc *= num
		case "SET":
			acc = num
		}
	}
}

func readCRLF(reader *bufio.Reader) (string, error) {
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	data = strings.TrimRight(data, "\r\n")
	return data, nil
}

func parseInput(input string) (string, int64, bool) {
	tokens := strings.Fields(input)
	if len(tokens) != 2 {
		return "", 0, false
	}
	cmd := strings.ToUpper(tokens[0])
	val, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return "", 0, false
	}
	return cmd, val, true
}
