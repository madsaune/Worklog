package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Worklog keeps track of starting, stopping and formatting our timer
type Worklog struct {
	Title     string
	Metadata  map[string]string
	StartTime time.Time
	StopTime  time.Time

	db *sql.DB
}

// NewWorklogClient returns a new Worklog client
func NewWorklogClient(args []string, db *sql.DB) *Worklog {
	w := &Worklog{}
	w.Title = args[1]
	w.db = db
	w.parseMetadata(args[1:])
	return w
}

func (w *Worklog) InitDB(path string) {
	// Create table if database does not exist
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		statement, err := w.db.Prepare(`CREATE TABLE worklog(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, title TEXT, start_time TEXT, stop_time TEXT, metadata TEXT)`) // Prepare statement. This is good to avoid SQL injections
		if err != nil {
			log.Fatalf("Could not prepare schema: %v", err)
		}

		_, err = statement.Exec()
		if err != nil {
			log.Fatalf("Could not create table 'worklog': %v", err)
		}
	}
}

// Start start the timer
func (w *Worklog) Start() {
	w.StartTime = time.Now()
}

// Stop stops the timer
func (w *Worklog) Stop() {
	w.StopTime = time.Now()
}

// GetTotalDuration return the total duration from start to stop
func (w Worklog) GetTotalDuration() time.Duration {
	return w.StopTime.Sub(w.StartTime)
}

// GetDuration returns the duration between StopTime and StartTime
func (w Worklog) GetDuration() time.Duration {
	currentTime := time.Now()
	return currentTime.Sub(w.StartTime)
}

// NewEntry creates a new entry in our worklog database
func (w Worklog) NewEntry() error {

	statement, err := w.db.Prepare(`INSERT INTO worklog(title, start_time, stop_time) VALUES (?, ?, ?)`) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		return err
	}

	timeFmt := "2006-01-02 15:04:05 MST"
	startTimeFormat := w.StartTime.Format(timeFmt)
	stopTimeFormat := w.StopTime.Format(timeFmt)

	_, err = statement.Exec(w.Title, startTimeFormat, stopTimeFormat)
	if err != nil {
		return err
	}

	return nil
}

func (w *Worklog) parseMetadata(args []string) {
	metadata := make(map[string]string)

	for i := 0; i < len(args); i++ {
		parts := strings.Split(args[i], "=")
		if len(parts) > 1 {
			metadata[parts[0]] = parts[1]
		} else {
			metadata[parts[0]] = parts[0]
		}
	}

	w.Metadata = metadata
}

func (w Worklog) String() string {
	return fmt.Sprintf(
		"%-10s: %s\n%-10s: %s\n%-10s: %s\n%-10s: %s",
		"Title", w.Title,
		"Start", FormatTime(w.StartTime),
		"Stop", FormatTime(w.StopTime),
		"Duration", FormatDuration(w.GetTotalDuration()),
	)
}

// FormatTime formats a time.Time type to hh:mm:ss
func FormatTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

// FormatDuration formats a time.Duration type to 00h, 00m
func FormatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02dh, %02dm, %02ds", h, m, s)
}
