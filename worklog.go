package main

import (
	"fmt"
	"strings"
	"time"
)

// Worklog keeps track of starting, stopping and formatting our timer
type Worklog struct {
	Title     string
	Metadata  map[string]string
	StartTime time.Time
	StopTime  time.Time
}

// NewWorklogClient returns a new Worklog client
func NewWorklogClient(args []string) *Worklog {
	w := &Worklog{}
	w.Title = args[1]
	w.parseMetadata(args[1:])
	return w
}

// Start start the timer
func (w *Worklog) Start() {
	w.StartTime = time.Now()
}

// Stop stops the timer
func (w *Worklog) Stop() {
	w.StopTime = time.Now()
}

// GetDuration returns the duration between StopTime and StartTime
func (w Worklog) GetDuration() time.Duration {
	return w.StopTime.Sub(w.StartTime)
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
		"Duration", FormatDuration(w.GetDuration()),
	)
}

// FormatTime formats a time.Time type to hh:mm:ss
func FormatTime(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

// FormatDuration formats a time.Duration type to 00h, 00m
func FormatDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02dh, %02dm", h, m)
}
