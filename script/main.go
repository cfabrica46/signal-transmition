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

	var message []syscall.Signal

	s0 := syscall.SIGABRT
	s1 := syscall.SIGALRM

	originalWord := "Hola que tal? Uwu"

	wordBinary := convertToBinary(originalWord)

	for i := range wordBinary {
		switch wordBinary[i] {
		case 0:
			message = append(message, s0)
		case 1:
			message = append(message, s1)
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

func convertToBinary(originalWord string) (b []int) {

	for _, v := range originalWord {

		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			b = append(b, int((v>>move)&1))
		}

	}
	return
}
