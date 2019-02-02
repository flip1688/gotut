package main

import "time"

// Task type has ID and name
type Task struct {
	ID     int
	Name   string
	Time   time.Time
	Person Person
}

// Person type
type Person string
