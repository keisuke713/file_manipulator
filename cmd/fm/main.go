package main

import (
	"fm"
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK              = 0
	ExitCodeParserFlagError = 1
)

func main() {
	if err := run(os.Stdout, os.Stderr, os.Args); err != nil {
		switch err := err.(type) {
		default:
			fmt.Fprintf(os.Stderr, "%v \n", err)
		}
		os.Exit(ExitCodeParserFlagError)
	}
}

func run(stdout, stderr io.Writer, args []string) error {
	if len(args) < 2 {
		cmd := fm.CmdMap[fm.HELP]
		if err := cmd.Run([]string{}); err != nil {
			return fmt.Errorf("help command failed: %w", err)
		}
		return nil
	}

	sub := args[1]
	if cmd, ok := fm.CmdMap[fm.CmdName(sub)]; ok {
		if err := cmd.Run(args[2:]); err != nil {
			switch err := err.(type) {
			default:
				return fmt.Errorf("%q command failed: %q", sub, err)
			}
		}
	} else {
		return fmt.Errorf("unknown command %q", sub)
	}

	return nil
}
