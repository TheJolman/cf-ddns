package main

import (
	"fmt"
	"os"
)

func help() {
	fmt.Printf("%s\nUpdate dynamic public IPs\n", os.Args[0])
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		switch arg {
		case "-h", "--help":
			help()
		case "-4":
			fmt.Println("4")
		case "-6":
			fmt.Println("6")
		default:
			fmt.Fprintf(os.Stderr, "Error: unrecognized argument '%s'\n", arg)
			help()
		}
	}
}
