package main

import "testing"

func TestArgParsing(t *testing.T) {
	args := []string{
		"-p",
		"PREFIX_",
		"-f",
		"ini",
	}
	flags, err := parseArgs(args)

	if err != nil {
		t.Fatalf("Error while parsing command flags. %v", err.Error())
	}

	if flags.format != "ini" {
		t.Fatalf("Unexpected value for format flag. Expected %v, got %v", "ini", flags.format)
	}

	if flags.prefix != "PREFIX_" {
		t.Fatalf("Unexpected value for prefix flag. Expected %v, got %v", "PREFIX_", flags.prefix)
	}
}

func TestArgMissingPrefix(t *testing.T) {
	args := []string{}
	_, err := parseArgs(args)

	expect := MissingRequiredArgError{
		arg: "prefix",
	}

	if err != expect || err.Error() != expect.Error() {
		t.Fatalf("Did not error on missing requirement.")
	}
}

func TestEnvParsing(t *testing.T) {
	prefix := "PREFIX_"
	env := []string{
		prefix + "HTTPD__BIND_ADDRESS=0.0.0.0",
	}

	envMap := parseEnv(env, prefix)
	value, ok := envMap["HTTPD__BIND_ADDRESS"]

	if !ok {
		t.Fatal("Did not parse env key")
	}

	if value != "0.0.0.0" {
		t.Fatal("Did not parse env value")
	}
}

func TestSingleValueIni(t *testing.T) {
	env := mockSingleValue()

	input := conformInput{
		arguments: []string{
			"-p",
			"PREFIX_",
			"-f",
			"ini",
		},
		environment: env.envVars,
	}

	output, err := conform(input)

	if err != nil {
		t.Fatal(err.Error())
	}

	if output != env.iniOutput {
		t.Fatalf("Expected %v, instead got %v", env.iniOutput, output)
	}
}

func TestSingleValueJson(t *testing.T) {
	env := mockSingleValue()

	input := conformInput{
		arguments: []string{
			"-p",
			"PREFIX_",
			"-f",
			"json",
		},
		environment: env.envVars,
	}

	output, err := conform(input)

	if err != nil {
		t.Fatal(err.Error())
	}

	if output != env.jsonOutput {
		t.Fatalf("Expected %v, instead got %v", env.jsonOutput, output)
	}
}

type mockEnv struct {
	envVars    []string
	name       string
	iniOutput  string
	jsonOutput string
}

func mockSingleValue() mockEnv {
	envVars := []string{
		"PREFIX_IMA_VALUE=so value",
	}

	iniOutput :=
		"ima_value=so value\n"

	jsonOutput := `{
  "ima_value": "so value"
}`

	return mockEnv{
		name:       "Single value",
		envVars:    envVars,
		iniOutput:  iniOutput,
		jsonOutput: jsonOutput,
	}
}

func mockSingleObject() mockEnv {
	envVars := []string{
		"PREFIX_HTTPD__BIND_ADDRESS=0.0.0.0",
	}

	iniOutput :=
		`[httpd]
    bind_address=0.0.0.0`

	jsonOutput :=
		`{
      "httpd": {
        "bind_address": "0.0.0.0"
      }
    }`

	return mockEnv{
		name:       "Single object",
		envVars:    envVars,
		iniOutput:  iniOutput,
		jsonOutput: jsonOutput,
	}
}

func mockSingleArray() mockEnv {
	envVars := []string{
		"PREFIX_ARR___=somevalue",
		"PREFIX_ARR___=anothervalue",
	}

	iniOutput :=
		`arr[]=somevalue
    arr[]=anothervalue`

	jsonOutput :=
		`{
      "arr": [
        "somevalue",
        "anothervalue"
      ]
    }`

	return mockEnv{
		name:       "Single array",
		envVars:    envVars,
		iniOutput:  iniOutput,
		jsonOutput: jsonOutput,
	}
}

func mockNestedObject() mockEnv {
	envVars := []string{
		"PREFIX_HTTPD__SECOND_LEVEL__NESTED_THING=foo",
		"PREFIX_HTTPD__SECOND_LEVEL__NESTED_OTHER_THING=bar",
	}

	/*
	   ini doesn't really do nesting...
	*/
	iniOutput :=
		`[httpd.second_level]
    nested_thing=foo
    nested_other_thing=bar`

	jsonOutput :=
		`{
      "httpd": {
        "second_level": {
          "nested_thing": "foo",
          "nested_other_thing": "bar"
        }
      }
    }`

	return mockEnv{
		name:       "Nested object",
		envVars:    envVars,
		iniOutput:  iniOutput,
		jsonOutput: jsonOutput,
	}
}

func mockNestedArray() mockEnv {
	envVars := []string{
		"PREFIX_HTTPD__THE_ARR___=somevalue",
		"PREFIX_HTTPD__THE_ARR___=anothervalue",
	}

	iniOutput :=
		`[httpd]
    the_arr[]=somevalue
    the_arr[]=anothervalue`

	jsonOutput :=
		`{
      "httpd": {
        "the_arr": [
          "somevalue",
          "anothervalue"
        ]
      }
    }`

	return mockEnv{
		name:       "Nested object",
		envVars:    envVars,
		iniOutput:  iniOutput,
		jsonOutput: jsonOutput,
	}
}
