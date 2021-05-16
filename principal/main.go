package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

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
}
