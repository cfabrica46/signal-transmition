package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	//1001000 1101111 1101100 1100001
	pOrigin := []byte("Hola")
	fmt.Printf("%b\n", pOrigin)

	//b := make([]byte, 8)

	o, err := strconv.ParseInt("1001000", 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", []byte{byte(o)})
}
