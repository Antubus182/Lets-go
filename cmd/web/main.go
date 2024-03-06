package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
	dsn := flag.String("dsn", "admin:Onveilig41@tcp(192.168.2.150:3306)/letsgo?parseTime=true", "MySQL data source name")
	//password in open project, bad idea ;-)
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

	db, err := app.openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	} else {
		logger.Debug("Database Connection established")
	}

	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	logger.Info("Starting server", "port", *addr)
	//ListenAndServe takes the port and the mux
	err = http.ListenAndServe(*addr, app.muxroutes())
	logger.Error(err.Error())
	os.Exit(1)
}

func (app *application) openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	app.logger.Debug("Database Pinged")

	return db, nil
}
