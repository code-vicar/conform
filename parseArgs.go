package conform

import (
	"flag"
	"os"
)

const (
	ARG_FORMAT_USAGE      = "Format of output file"
	ARG_FORMAT_DEFAULT    = "ini"
	ARG_FORMAT_LONG_NAME  = "format"
	ARG_FORMAT_SHORT_NAME = "f"

	ARG_OUTPUT_USAGE      = "Path to output file"
	ARG_OUTPUT_DEFAULT    = ""
	ARG_OUTPUT_LONG_NAME  = "output"
	ARG_OUTPUT_SHORT_NAME = "o"

	ARG_PREFIX_USAGE      = "Environment variable prefix"
	ARG_PREFIX_DEFAULT    = ""
	ARG_PREFIX_LONG_NAME  = "prefix"
	ARG_PREFIX_SHORT_NAME = "p"
)

type commandFlags struct {
	format, output, prefix string
}

func parseArgs(arguments []string) (commandFlags, error) {
	flags := commandFlags{}

	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	flagSet.StringVar(&flags.format, ARG_FORMAT_LONG_NAME, ARG_FORMAT_DEFAULT, ARG_FORMAT_USAGE)
	flagSet.StringVar(&flags.format, ARG_FORMAT_SHORT_NAME, ARG_FORMAT_DEFAULT, ARG_FORMAT_USAGE)
	flagSet.StringVar(&flags.output, ARG_OUTPUT_LONG_NAME, ARG_OUTPUT_DEFAULT, ARG_OUTPUT_USAGE)
	flagSet.StringVar(&flags.output, ARG_OUTPUT_SHORT_NAME, ARG_OUTPUT_DEFAULT, ARG_OUTPUT_USAGE)
	flagSet.StringVar(&flags.prefix, ARG_PREFIX_LONG_NAME, ARG_PREFIX_DEFAULT, ARG_PREFIX_USAGE)
	flagSet.StringVar(&flags.prefix, ARG_PREFIX_SHORT_NAME, ARG_PREFIX_DEFAULT, ARG_PREFIX_USAGE)

	err := flagSet.Parse(arguments)

	if err == nil && len(flags.prefix) == 0 {
		err = MissingRequiredArgError{"prefix"}
	}

	return flags, err
}
