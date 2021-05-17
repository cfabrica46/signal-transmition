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

	var b []byte
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

		for i := 0; i < 7; i++ {

			sig := <-sigs

			switch sig {
			case syscall.SIGABRT:

				b = append(b, '0')

			case syscall.SIGALRM:

				b = append(b, '1')

			default:
			}

			if sig == syscall.SIGINT {
				exit = true
				break
			}

		}

	}

	nLetters := len(b) / 7

	for i := 0; i < nLetters; i++ {

		letterBinary := fmt.Sprintf("%s", b[i*7:7*(i+1)])

		letter, err := strconv.ParseInt(letterBinary, 2, 64)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", []byte{byte(letter)})

	}

	fmt.Println()
}
