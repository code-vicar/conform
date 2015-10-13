package main

import (
	"testing"
)

func TestArgParsing(t *testing.T) {
	args := getMockArgs()
	flags, err := parseArgs(args)

	if err != nil {
		t.Fatal("Error while parsing command flags")
	}

	if flags.format != "ini" {
		t.Fatal("Unexpected value for format flag")
	}

	if flags.prefix != "COUCHDB_" {
		t.Fatal("Unexpected value for prefix flag")
	}
}

func TestEnvParsing(t *testing.T) {
	env := getMockEnv()
	prefix := "COUCHDB_"

	envMap := parseEnv(env, prefix)
	value, ok := envMap["HTTPD{}BIND_ADDRESS"]

	if !ok {
		t.Fatal("Did not parse env key")
	}

	if value != "0.0.0.0" {
		t.Fatal("Did not parse env value")
	}
}

func TestOutputsIniFormat(t *testing.T) {
	input, want := setup()

	output := conform(input)

	if output != want {
		t.Fatal("Output did not match wanted")
	}
}

func setup() (conformInput, string) {
	input := conformInput{
		environment: getMockEnv(),
		arguments:   getMockArgs(),
	}

	want := `[httpd]
  bind_address 0.0.0.0
  `
	return input, want
}

func getMockEnv() []string {
	env := []string{
		"COUCHDB_HTTPD{}BIND_ADDRESS=0.0.0.0",
	}
	return env
}

func getMockArgs() []string {
	args := []string{
		"-p",
		"COUCHDB_",
		"-f",
		"ini",
	}

	return args
}
