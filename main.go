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

const (
	DEFAULT_COLOR string = "\033[0m"
	RETURN_COLOR  string = "\033[90m"
	ERROR_COLOR   string = "\033[31m"
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
			fmt.Println(ERROR_COLOR, err, DEFAULT_COLOR)
		} else if r, err := evaluator.EvalProgram(p); err != nil {
			fmt.Println(ERROR_COLOR, err, DEFAULT_COLOR)
		} else {
			fmt.Println(RETURN_COLOR, fmt.Sprintf("; %s", r.Inspect()), DEFAULT_COLOR)
		}

	}
}
