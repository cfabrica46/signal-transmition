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

	b0 := syscall.SIGABRT
	b1 := syscall.SIGALRM

	h := []syscall.Signal{b1, b0, b0, b1, b0, b0, b0}
	o := []syscall.Signal{b1, b1, b0, b1, b1, b1, b1}
	l := []syscall.Signal{b1, b1, b0, b1, b1, b0, b0}
	a := []syscall.Signal{b1, b1, b0, b0, b0, b0, b1}

	message := [][]syscall.Signal{h, o, l, a}

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

		for indx := range message[i] {

			err := process.Signal(message[i][indx])

			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(100 * time.Microsecond)

		}

	}

	process.Signal(syscall.SIGINT)

	fmt.Println("Mensaje enviado correctamente")

}
