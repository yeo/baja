package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	askCommand := flag.NewFlagSet("ask", flag.ExitOnError)
	questionFlag := askCommand.String("question", "", "Question that you are asking for")

	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	recipientFlag := sendCommand.String("recipient", "", "Recipient of your message")
	messageFlag := sendCommand.String("message", "", "Text message")

	if len(os.Args) == 1 {
		fmt.Println("usage: siri <command> [<args>]")
		fmt.Println("The most commonly used git commands are: 	")
		fmt.Println(" ask   Ask questions")
		fmt.Println(" send  Send messageFlagges to your contacts")
		return
	}

	switch os.Args[1] {
	case "ask":
		askCommand.Parse(os.Args[2:])
	case "send":
		sendCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
