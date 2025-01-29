package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	ev "github.com/tamercuba/golisp/evaluator"
	lx "github.com/tamercuba/golisp/lexer"
	pr "github.com/tamercuba/golisp/parser"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	fmt.Println("Welcome to the GoLisp REPL!")
	fmt.Println("Type 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)
	evaluator := ev.NewEvaluator()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	// Goroutine para lidar com sinais
	go func() {
		for {
			<-signalChan
			clearScreen()
			fmt.Print("> ")
		}
	}()

	for {
		fmt.Print("> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		l := lx.NewLexer(input)
		p, err := pr.ParseProgram(l)

		if err != nil {
			fmt.Println(err)
		}

		result := evaluator.EvalProgram(p)
		fmt.Printf("%s\n", result.Inspect())
	}
}
