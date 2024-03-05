package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	fmt.Println("Hello Snippet")

	addr := flag.String("addr", ":4000", "HTTP network port")
	debugEnable := flag.Bool("debug", false, "Set the log level to debug")
	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()

	logLevel := new(slog.LevelVar)
	if !*debugEnable {
		logLevel.Set(slog.LevelInfo)
	} else {
		logLevel.Set(slog.LevelDebug)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	app := &application{
		logger: logger,
	}

	logger.Info("Starting server", "port", *addr)
	//ListenAndServe takes the port and the mux
	err := http.ListenAndServe(*addr, app.muxroutes())
	logger.Error(err.Error())
	os.Exit(1)
}
