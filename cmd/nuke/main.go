package main

import (
	"fmt"
	"os"

	"github.com/alexintosh/gocleaner/pkg/app"
)

func main() {
	if err := app.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
} 