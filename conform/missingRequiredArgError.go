package main

import "fmt"

type MissingRequiredArgError struct {
	arg string
}

func (err MissingRequiredArgError) Error() string {
	return fmt.Sprintf("Missing required arg, %v", err.arg)
}
