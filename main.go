package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"os/user"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	worklogPath  = "/.worklog"
	databasePath = worklogPath + "/worklog.db"
)

func main() {

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	userHomeDir := currentUser.HomeDir

	// Create .worklog folder under user HOME if not exist
	_, err = os.Stat(userHomeDir + worklogPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(userHomeDir+worklogPath, 0700)
		if err != nil {
			log.Fatalln("Could not create .worklog folder in users HomeDir")
		}
	}

	sqlDatabase, err := sql.Open("sqlite3", userHomeDir+databasePath)
	if err != nil {
		log.Fatal("Could not open database")
	}

	w := NewWorklogClient(os.Args, sqlDatabase)
	w.InitDB(userHomeDir + databasePath)
	w.Start()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		w.Stop()
		err := w.NewEntry()
		if err != nil {
			log.Fatalln("Could not create a new entry in the database")
		}
		fmt.Printf("\n\n%s", w)

		os.Exit(0)
	}()

	fmt.Println("Started tracking...")
	fmt.Println()
	fmt.Println("Press CTRL+C to stop.")
	for {
		time.Sleep(10 * time.Minute) // or runtime.Gosched() or similar per @misterbee
		fmt.Printf("Still tracking... [%s]\n", FormatDuration(w.GetDuration()))
	}
}
