package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type conformInput struct {
	arguments, environment []string
}

func main() {
	input := conformInput{
		arguments:   os.Args[1:],
		environment: os.Environ(),
	}
	output, err := conform(input)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(output)
	}
}

func conform(input conformInput) (string, error) {
	flags, err := parseArgs(input.arguments)

	if err != nil {
		log.Fatal(err.Error())
	}

	envMap := parseEnv(input.environment, flags.prefix)

	switch {
	case flags.format == "json":
		return createJson(envMap)
	default:
		return createIni(envMap)
	}
}

func createIni(m map[string]string) (string, error) {
	ini := ""

	for key, value := range m {
		ini += fmt.Sprintf("%v=%v\n", strings.ToLower(key), value)
	}

	return ini, nil
}

func createJson(m map[string]string) (string, error) {
	lowerCaseM := make(map[string]string)

	for key, value := range m {
		lowerCaseM[strings.ToLower(key)] = value
	}
	b, err := json.MarshalIndent(lowerCaseM, "", "  ")

	if err != nil {
		return "", err
	}

	return string(b), nil
}
