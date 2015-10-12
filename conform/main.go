package main

import (
	"fmt"
	"os"
)

type conformInput struct {
	environment, arguments []string
}

func parseEnv(env []string, prefix string) map[string]string {
	var m = make(map[string]string)

	return m
}

func conform(input conformInput) string {
	return ""
}

func main() {
	input := conformInput{
		os.Environ(),
		os.Args[1:],
	}
	fmt.Println(conform(input))
}
