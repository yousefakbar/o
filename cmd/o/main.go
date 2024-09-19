package main

import (
	"fmt"
	"os"
	"github.com/yousefakbar/o/internal/cli"
)

func main() {
	// Start `o` -- The obsidian CLI wrapper
	if err := cli.Run(os.Args); err != nil {
	    // Handle the error, e.g., log it, print a message, exit the program
	    fmt.Fprintln(os.Stderr, "Error:", err)
	    os.Exit(1)
	}
}
