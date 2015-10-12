package main

import (
	"testing"
)

func TestEnvParsing(t *testing.T) {
	env := getMockEnvCouchDB()
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
	input, want := initIni()

	output := conform(input)

	if output != want {
		t.Fatal("Output did not match wanted")
	}
}

func initIni() (conformInput, string) {
	input := conformInput{
		getMockEnvCouchDB(),
		getMockArgsIni(),
	}

	want := `[httpd]
  bind_address 0.0.0.0
  `
	return input, want
}

func getMockEnvCouchDB() []string {
	env := []string{
		"COUCHDB_HTTPD{}BIND_ADDRESS=0.0.0.0",
	}
	return env
}

func getMockArgsIni() []string {
	args := []string{
		"-p",
		"COUCHDB_",
		"-f",
		"ini",
	}

	return args
}
