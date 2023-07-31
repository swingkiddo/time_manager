package main

import (
	"errors"
	"fmt"
	"time"
)

type Scheduler struct {
	schedule Schedule
}

func (s *Scheduler) AddEvent(event *Event) error {
	return nil
}

func (s *Scheduler) AddDay(year int, month time.Month, day int) (*Day, error) {
	d := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	if e := s.schedule.GetDayByDate(d); e != nil {
		return nil, errors.New(d.String() + " already exists")
	}
	newDay := Day{Date: d}
	s.schedule.Days = append(s.schedule.Days, newDay)
	return &newDay, nil
}

type Schedule struct {
	Days []Day
}

func (s *Schedule) GetDayByDate(date time.Time) *Day {
	for _, day := range s.Days {
		if day.Date == date {
			return &day
		}
	}
	return nil
}

type Day struct {
	Date time.Time
	Events []Event
}

func (d *Day) AddEvent(name string, c Category, start time.Time, end time.Time, s Status) error {
	duration := end.Sub(start)
	e := Event{Name: name, Category: c, Duration: &duration, StartTime: start, EndTime: end, Status: s}
	d.Events = append(d.Events, e)
	return nil
}

type Category = uint64
const (
	Fortnite Category = iota
	Coding
	Sports
	Streaming
	Music
	Work
) 

type Status = uint64
const (
	Done Status = iota
	NotDone
)

type Event struct {
	Name string
	Category Category
	Duration *time.Duration
	StartTime time.Time
	EndTime time.Time
	Status Status
}

func main() {
	schedule := Schedule{Days: make([]Day, 0)}
	s := Scheduler{schedule: schedule}
	newDay, err := s.AddDay(2023, 8, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	var start time.Time
	var end time.Time
	start = start.Add(1 * time.Hour).Add(30 * time.Minute)
	end = end.Add(3 * time.Hour)

	newDay.AddEvent("Creative practice", 0, start, end, 1)
	fmt.Println(newDay.Events)
	
}

