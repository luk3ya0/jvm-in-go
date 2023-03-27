package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOptions   string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage

	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")

	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")

	flag.StringVar(&cmd.cpOptions, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOptions, "cp", "", "classpath")

	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
