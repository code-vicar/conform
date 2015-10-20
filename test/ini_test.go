package conform_test

import (
	"testing"

	conform "github.com/code-vicar/conform"
)

func TestSingleValueIni(t *testing.T) {
	env := mockSingleValue()

	input := conform.Input{
		Arguments: []string{
			"-p",
			"PREFIX_",
			"-f",
			"ini",
		},
		Environment: env.envVars,
	}

	output, err := conform.Run(input)

	if err != nil {
		t.Fatal(err.Error())
	}

	if output != env.iniOutput {
		t.Fatalf("Expected %v, instead got %v", env.iniOutput, output)
	}
}
