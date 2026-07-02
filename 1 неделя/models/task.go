package models

import "fmt"

type Task struct {
	Title string
	Hours float64
	Done  bool
}

func (s Task) Summary() string {
	s.Done = true
	return fmt.Sprintf("%s: %.1f часов", s.Title, s.Hours)
}
func (c *Task) Complete() bool {
	c.Done = true

	return true
}
func Reset(t *Task) {
	t.Done = false
	return
}
