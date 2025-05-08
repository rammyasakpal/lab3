TCP Calculator Server in Go

This project implements a concurrent TCP-based calculator server in Go. The server accepts client connections, processes arithmetic commands over the network, and returns results in real time using the CRLF protocol format.

How It Works

Each client session supports basic arithmetic operations:
- ADD n – Add n to accumulator
- SUB n – Subtract n from accumulator
- MUL n – Multiply accumulator by n
- SET n – Set accumulator to n
- Sending a blank line (\r\n) returns the current value and resets the accumulator.

Each session is handled in its own goroutine, supporting concurrent clients.

Build & Run
bash: go build -o bin/calculator src/server.go
