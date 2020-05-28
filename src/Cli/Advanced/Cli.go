package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// SubCommands.
	// Eg: git commit -m "message" | git push
	helpCommand := flag.NewFlagSet("help", flag.ExitOnError)
	helloWorldCommand := flag.NewFlagSet("helloWorld", flag.ExitOnError)

	printCommand := flag.NewFlagSet("print", flag.ExitOnError)
	printPointer := printCommand.String("message", "", "message <Message to print> (Required)")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("Required to use a sub command.")
		os.Exit(1)
	}

	parseFlagsToSubCommands(helpCommand, helloWorldCommand, printCommand)

	// help
	helpCommandHandler(helpCommand, helloWorldCommand, printCommand)

	// helloWorld
	helloWorldCommandHandler(helloWorldCommand)

	// print
	printCommandHandler(printCommand, printPointer)

}

func printCommandHandler(printCommand *flag.FlagSet, printPointer *string) {
	if printCommand.Parsed() {
		if *printPointer == "" {
			printCommand.PrintDefaults()
			os.Exit(1)
		}

		fmt.Println(*printPointer)
	}
}

func helloWorldCommandHandler(helloWorldCommand *flag.FlagSet) {
	if helloWorldCommand.Parsed() {
		fmt.Println("Hello World")
	}
}

func helpCommandHandler(helpCommand *flag.FlagSet, helloWorldCommand *flag.FlagSet, printCommand *flag.FlagSet) {
	if helpCommand.Parsed() {
		fmt.Println("Hello World:")
		helloWorldCommand.PrintDefaults()
		fmt.Println("Print:")
		printCommand.PrintDefaults()
	}
}

func parseFlagsToSubCommands(helpCommand *flag.FlagSet, helloWorldCommand *flag.FlagSet, printCommand *flag.FlagSet) {
	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case "help":
		helpCommand.Parse(os.Args[2:])
	case "helloWorld":
		helloWorldCommand.Parse(os.Args[2:])
	case "print":
		printCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults() // ToDo find better defaults.
		os.Exit(1)
	}
}

