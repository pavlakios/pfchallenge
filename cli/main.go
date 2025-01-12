package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// ErrHelp provides context that help was given.
var ErrHelp = errors.New("provided help")

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmsgprefix | log.Lshortfile)
}

func main() {
	if err := processCommands(os.Args); err != nil {
		fmt.Println("msg", err)
		os.Exit(1)
	}
}

// processCommands handles the execution of the commands specified on the command line.
func processCommands(args []string) error {
	switch args[1] {
	case "sessions":
		Sessions()
	default:
		fmt.Println("sessions -numb: open sessions simultaneously")
		fmt.Println("provide a command to get more help.")
		return ErrHelp
	}

	return nil
}
