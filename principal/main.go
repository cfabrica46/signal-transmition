package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	var messageBinary []bool
	var exit bool

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGABRT, syscall.SIGALRM, syscall.SIGINT)

	f, err := os.OpenFile("../pid.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f.Truncate(0)

	_, err = f.Write([]byte(strconv.Itoa(os.Getpid())))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Esperando Signals")

	for !exit {

		for i := 0; i < 8; i++ {
			sig := <-sigs

			switch sig {
			case syscall.SIGABRT:

				messageBinary = append(messageBinary, false)

			case syscall.SIGALRM:

				messageBinary = append(messageBinary, true)

			default:
			}

			if sig == syscall.SIGINT {
				exit = true
				break
			}

		}

	}

	message := convertToString(messageBinary)

	fmt.Printf("Mensaje Recivido: %s\n", message)

}

func convertToString(messageBinary []bool) (message []byte) {
	var container int

	for i := range messageBinary {

		var binaryValue int

		switch messageBinary[i] {
		case false:
			binaryValue = 0
		case true:
			binaryValue = 1
		}

		move := i % 8

		a := binaryValue << (7 - move)

		container += a

		if move == 7 {
			message = append(message, byte(container))
			container = 0
		}

	}
	return
}
