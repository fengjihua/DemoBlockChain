package main

import (
	"DemoBlockChain/lib"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Execute() error {
	lib.Log.Notice("Starting ABCIClient Tendermint")

	f, err := os.Create("logs/tendermint.log")
	if err != nil {
		fmt.Println("Tendermint log init error:", err)
	}
	multiWriter := io.MultiWriter(f, os.Stdout)

	go func() {
		cmd := exec.Command("bash", "-c", "sh run-tm.sh")
		cmd.Stdout = multiWriter
		cmd.Start()
	}()

	runConsole()

	return nil
}

func runConsole() error {
	for {
		fmt.Printf("> ")
		bufReader := bufio.NewReader(os.Stdin)
		line, more, err := bufReader.ReadLine()
		if more {
			return errors.New("Input is too long")
		} else if err != nil {
			return err
		}

		fmt.Println("ABCI Client,", line)
	}
}
