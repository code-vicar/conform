package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

const (
	ENVIRONMENT_KEYVALUE_SEPARATOR = "="

	ARG_FORMAT_USAGE      = "Format of output file"
	ARG_FORMAT_DEFAULT    = ""
	ARG_FORMAT_LONG_NAME  = "format"
	ARG_FORMAT_SHORT_NAME = "f"

	ARG_OUTPUT_USAGE      = "Path to output file"
	ARG_OUTPUT_DEFAULT    = ""
	ARG_OUTPUT_LONG_NAME  = "output"
	ARG_OUTPUT_SHORT_NAME = "o"

	ARG_PREFIX_USAGE      = "Environment variable prefix"
	ARG_PREFIX_DEFAULT    = "ini"
	ARG_PREFIX_LONG_NAME  = "prefix"
	ARG_PREFIX_SHORT_NAME = "p"
)

type conformInput struct {
	arguments, environment []string
}

type commandFlags struct {
	format, output, prefix string
}

func main() {
	input := conformInput{
		arguments:   os.Args[1:],
		environment: os.Environ(),
	}
	log.Println(conform(input))
}

func conform(input conformInput) string {
	flags, err := parseArgs(input.arguments)

	if err != nil {
		log.Fatal(err.Error())
	}

	m := parseEnv(input.environment, flags.prefix)

	log.Println(m)
	return ""
}

func parseArgs(arguments []string) (commandFlags, error) {
	flags := commandFlags{}

	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	flagSet.StringVar(&flags.format, ARG_FORMAT_LONG_NAME, ARG_FORMAT_DEFAULT, ARG_FORMAT_USAGE)
	flagSet.StringVar(&flags.format, ARG_FORMAT_SHORT_NAME, ARG_FORMAT_DEFAULT, ARG_FORMAT_USAGE)
	flagSet.StringVar(&flags.output, ARG_OUTPUT_LONG_NAME, ARG_OUTPUT_DEFAULT, ARG_OUTPUT_USAGE)
	flagSet.StringVar(&flags.output, ARG_OUTPUT_SHORT_NAME, ARG_OUTPUT_DEFAULT, ARG_OUTPUT_USAGE)
	flagSet.StringVar(&flags.prefix, ARG_PREFIX_LONG_NAME, ARG_PREFIX_DEFAULT, ARG_PREFIX_USAGE)
	flagSet.StringVar(&flags.prefix, ARG_PREFIX_SHORT_NAME, ARG_PREFIX_DEFAULT, ARG_PREFIX_USAGE)

	err := flagSet.Parse(arguments)

	return flags, err
}

func parseEnv(environ []string, prefix string) map[string]string {
	var m = make(map[string]string)

	for _, envVar := range environ {
		if strings.HasPrefix(envVar, prefix) {
			key, value := getKeypairWithoutPrefix(envVar, prefix)
			m[key] = value
		}
	}

	return m
}

func getKeypairWithoutPrefix(envVar string, prefix string) (string, string) {
	key, value := splitEnvVarIntoKeypair(envVar)
	key = strings.TrimPrefix(key, prefix)

	return key, value
}

func splitEnvVarIntoKeypair(envVar string) (string, string) {
	kvPair := strings.Split(envVar, ENVIRONMENT_KEYVALUE_SEPARATOR)
	key := kvPair[0]
	value := kvPair[1]
	return key, value
}
