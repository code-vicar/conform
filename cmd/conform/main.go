package main

import (
	"log"
	"os"

	conform "github.com/code-vicar/conform"
)

func main() {
	input := conform.Input{
		Arguments:   os.Args[1:],
		Environment: os.Environ(),
	}
	output, err := conform.Run(input)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(output)
	}
}
