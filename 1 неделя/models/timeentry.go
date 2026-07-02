package models

import "time"

type TimeEntry struct {
	Task      string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	IsRunning bool
}

func NewTimeEntry(task string) *TimeEntry {
	entry := TimeEntry{
		Task:      task,
		StartTime: time.Now(),
		IsRunning: true,
	}
	return &entry
}
