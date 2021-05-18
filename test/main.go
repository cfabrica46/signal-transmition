package main

import (
	"fmt"
)

func main() {

	//fmt.Printf("%08b\n", 1)
	//fmt.Printf("%08b\n", 2)
	//fmt.Printf("%08b\n", 4)
	//fmt.Printf("%08b\n", 8)
	//fmt.Printf("%08b\n", 16)
	//fmt.Printf("%08b\n", 32)
	//fmt.Printf("%08b\n", 64)
	//fmt.Printf("%08b\n", 128)

	var b []bool
	s := "Holi"

	for _, v := range s {

		for i := 0; i < 8; i++ {
			move := uint(7 - i)

			switch int((v >> move) & 1) {
			case 0:
				b = append(b, false)
			case 1:
				b = append(b, true)
			}
		}

	}

	fmt.Println(b)

	var container int

	bb := []byte{}

	for i := range b {

		var x int

		switch b[i] {
		case false:
			x = 0
		case true:
			x = 1
		}

		move := i % 8

		a := x << (7 - move)

		container += a

		if move == 7 {
			bb = append(bb, byte(container))
			container = 0
		}

	}
	fmt.Printf("%v\n", []byte(s))
	fmt.Printf("%v\n", bb)
	fmt.Printf("%s\n", bb)

}
