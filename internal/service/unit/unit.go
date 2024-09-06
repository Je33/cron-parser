package unit

import (
	"cron-parser/internal/domain"
	"errors"
	"fmt"
)

var (
	ErrUnitPosition = errors.New("unknown unit position")
	ErrUnitParse    = errors.New("unit parse error")
)

type DateParser interface {
	Parse(s string, min int, max int) ([]string, error)
}

type CommandParser interface {
	Parse(s string) (string, error)
}

type Parser struct {
	dateParser    DateParser
	commandParser CommandParser
}

func NewParser(dateParser DateParser, commandParser CommandParser) *Parser {
	return &Parser{
		dateParser:    dateParser,
		commandParser: commandParser,
	}
}

func (p *Parser) Parse(unit string, position domain.SchedulePosition, schedule *domain.Schedule) error {
	var err error

	switch position {
	case domain.SchedulePositionMinute:
		schedule.Minute, err = p.dateParser.Parse(unit, 0, 59)

	case domain.SchedulePositionHour:
		schedule.Hour, err = p.dateParser.Parse(unit, 0, 23)

	case domain.SchedulePositionDayMonth:
		schedule.DayMonth, err = p.dateParser.Parse(unit, 1, 31)

	case domain.SchedulePositionMonth:
		schedule.Month, err = p.dateParser.Parse(unit, 1, 12)

	case domain.SchedulePositionDayWeek:
		schedule.DayWeek, err = p.dateParser.Parse(unit, 0, 6)

	case domain.SchedulePositionCommand:
		schedule.Command, err = p.commandParser.Parse(unit)

	default:
		return fmt.Errorf("%w: %d", ErrUnitPosition, position)
	}

	if err != nil {
		return fmt.Errorf("%w: %w", ErrUnitParse, err)
	}

	return nil
}
