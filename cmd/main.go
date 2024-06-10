package main

import (
	"flag"
	"fmt"
	"os"
)

// I don't have any time for formatting or normal code. Go fuck yourself

func main() {
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	serverPort := serverCmd.Int("port", 8080, "Port for the server to listen on")

	agentCmd := flag.NewFlagSet("agent", flag.ExitOnError)
	agentHost := agentCmd.String("host", "localhost", "Host for the agent to connect to")

	if len(os.Args) < 2 {
		fmt.Println("expected 'server' or 'agent' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		serverCmd.Parse(os.Args[2:])
	case "agent":
		agentCmd.Parse(os.Args[2:])
	default:
		fmt.Println("expected 'server' or 'agent' subcommands")
		os.Exit(1)
	}

	if serverCmd.Parsed() {
		fmt.Printf("Running server on port %d\n", *serverPort)
	}

	if agentCmd.Parsed() {
		fmt.Printf("Running agent, connecting to host %s\n", *agentHost)
	}
}
