package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"syscall"
	"time"
)

func main() {

	// syscall.SIGABRT = 0
	// syscall.SIGALRM = 1
	// syscall.SIGINT = EOF

	//1001000 1101111 1101100 1100001

	var b []byte
	var message []syscall.Signal

	b0 := syscall.SIGABRT
	b1 := syscall.SIGALRM

	originalWord := "Hola"

	for i := range []byte(originalWord) {
		letterBinary := fmt.Sprintf("%b", originalWord[i])
		b = append(b, []byte(letterBinary)...)
	}

	for i := range b {

		switch string(b[i]) {
		case "0":
			message = append(message, b0)
		case "1":
			message = append(message, b1)
		}

	}

	pidString, err := ioutil.ReadFile("../pid.txt")

	if err != nil {
		log.Fatal(err)
	}

	pid, err := strconv.Atoi(string(pidString))

	if err != nil {
		log.Fatal(err)
	}

	process, err := os.FindProcess(pid)

	if err != nil {
		log.Fatal(err)
	}

	for i := range message {

		err := process.Signal(message[i])

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(500 * time.Microsecond)

	}

	process.Signal(syscall.SIGINT)

	fmt.Println("Mensaje enviado correctamente")

}
