package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	cmdline = flag.String("c", "", "cmdline to execute")
	file    = flag.String("f", "", "output file")
	verbose = flag.String("v", "", "verbose output")
)

func main() {
	flag.Parse()

	if *cmdline == "" {
		fmt.Printf("cmdline not defined.\n")
		flag.Usage()
		os.Exit(1)
	}

	if *file == "" {
		fmt.Printf("file not defined.\n")
		flag.Usage()
		os.Exit(1)
	}

	args := strings.Split(*cmdline, " ")

	if *verbose != "" {
		for i, arg := range args {
			fmt.Printf("args[%d]:%s\n", i, arg)
		}
	}

	cmd := exec.Command(args[0], args[1:]...)

	file, err := os.Create(*file)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		file.Close()
	}()

	cmd.Stdout = io.MultiWriter(os.Stdout, file)
	cmd.Stderr = file

	err = cmd.Run()
	if err != nil {
		os.Exit(cmd.ProcessState.ExitCode())
	}
}
