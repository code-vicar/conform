package conform

import "testing"

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
