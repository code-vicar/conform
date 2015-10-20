package conform

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Input struct {
	Arguments, Environment []string
}

func Run(input Input) (string, error) {
	flags, err := parseArgs(input.Arguments)

	if err != nil {
		return "", err
	}

	envMap := parseEnv(input.Environment, flags.prefix)

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
