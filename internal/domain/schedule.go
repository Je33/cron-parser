package domain

type Schedule struct {
	Minute   []string
	Hour     []string
	DayMonth []string
	Month    []string
	DayWeek  []string
	Command  string
}

type SchedulePosition int

const (
	SchedulePositionMinute SchedulePosition = iota
	SchedulePositionHour
	SchedulePositionDayMonth
	SchedulePositionMonth
	SchedulePositionDayWeek
	SchedulePositionCommand
)
