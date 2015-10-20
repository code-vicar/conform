package conform_test

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
