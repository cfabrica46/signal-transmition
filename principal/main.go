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

	var bb [][]byte

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

loop:
	for {
		var b []byte

		for i := 0; i < 7; i++ {

			sig := <-sigs

			switch sig {
			case syscall.SIGABRT:

				b = append(b, []byte("0")...)

			case syscall.SIGALRM:

				b = append(b, []byte("1")...)

			case syscall.SIGINT:

				break loop
			}

		}

		bb = append(bb, b)
	}

	for _, v := range bb {

		letter, err := strconv.ParseInt(string(v), 2, 64)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", []byte{byte(letter)})

	}

	fmt.Println()
}
