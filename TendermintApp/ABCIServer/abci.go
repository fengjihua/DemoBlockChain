package main

import (
	"DemoBlockChain/TendermintApp/ABCIServer/example/counter"
	"DemoBlockChain/TendermintApp/ABCIServer/example/dummy"
	"DemoBlockChain/lib"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	abcicli "github.com/tendermint/abci/client"
	"github.com/tendermint/abci/server"
	"github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/log"
)

// client is a global variable so it can be reused by the console
var (
	client abcicli.Client
	logger log.Logger
)

// flags
var (
	// global
	flagAddress  = "tcp://0.0.0.0:46658" // Address of application socket
	flagAbci     = "socket"              // Either socket or grpc
	flagVerbose  = false                 // for the println output
	flagLogLevel = "debug"               // for the logger

	// query
	flagPath   string
	flagHeight int
	flagProve  bool

	// counter
	flagAddrC  string
	flagSerial bool

	// dummy
	flagAddrD   string
	flagPersist string
)

func Execute() error {
	err := preRun()
	lib.HandleError(err)

	go func() {
		// err := runCounter()
		err := runDummy()
		lib.HandleError(err)
	}()

	runConsole()

	return nil
}

func preRun() error {
	if logger == nil {
		allowLevel, err := log.AllowLevel(flagLogLevel)
		if err != nil {
			return err
		}

		f, err := os.Create("logs/abci.log")
		if err != nil {
			fmt.Println("ABCI log init error:", err)
		}
		multiWriter := io.MultiWriter(f, os.Stdout)

		// logger = log.NewFilter(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), allowLevel)
		logger = log.NewFilter(log.NewTMLogger(log.NewSyncWriter(multiWriter)), allowLevel)
	}
	return nil
}

func runAccountBook() error {
	return nil
}

func runCounter() error {

	app := counter.NewCounterApplication(flagSerial)

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	// Start the listener
	srv, err := server.NewServer(flagAddress, flagAbci, app)
	if err != nil {
		return err
	}
	srv.SetLogger(logger.With("module", "abci-server"))
	if err := srv.Start(); err != nil {
		return err
	}

	// Wait forever
	cmn.TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})
	return nil
}

func runDummy() error {
	// logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	// Create the application - in memory or persisted to disk
	var app types.Application
	if flagPersist == "" {
		app = dummy.NewDummyApplication()
	} else {
		app = dummy.NewPersistentDummyApplication(flagPersist)
		app.(*dummy.PersistentDummyApplication).SetLogger(logger.With("module", "dummy"))
	}

	// Start the listener
	srv, err := server.NewServer(flagAddress, flagAbci, app)
	if err != nil {
		return err
	}
	srv.SetLogger(logger.With("module", "abci-server"))
	if err := srv.Start(); err != nil {
		return err
	}

	// Wait forever
	cmn.TrapSignal(func() {
		// Cleanup
		srv.Stop()
	})
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

		fmt.Println("ABCI Server,", line)
	}
}
