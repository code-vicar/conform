package conform_test

import (
	"testing"

	conform "github.com/code-vicar/conform"
)

func TestSingleValueJson(t *testing.T) {
	env := mockSingleValue()

	input := conform.Input{
		Arguments: []string{
			"-p",
			"PREFIX_",
			"-f",
			"json",
		},
		Environment: env.envVars,
	}

	output, err := conform.Run(input)

	if err != nil {
		t.Fatal(err.Error())
	}

	if output != env.jsonOutput {
		t.Fatalf("Expected %v, instead got %v", env.jsonOutput, output)
	}
}
