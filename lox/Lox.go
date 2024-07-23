package lox

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	var content, err = os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %s\n", err)
		os.Exit(1)
	}
	run(string(content))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading a line")
		}
		run(line)
	}
}
