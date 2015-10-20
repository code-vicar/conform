package conform

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
