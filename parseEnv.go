package conform

import "strings"

const ENVIRONMENT_KEYVALUE_SEPARATOR = "="

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
