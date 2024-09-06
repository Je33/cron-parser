package schedule

import (
	"cron-parser/internal/domain"
	"errors"
	"fmt"
	"regexp"
)

// Delimiter between cron string parts
var scheduleDelimiter = regexp.MustCompile(`\s+`)

// Limit of cron string parts
const scheduleDelimiterLimit = 6

var (
	ErrScheduleFormat = errors.New("invalid schedule format")
)

type UnitParser interface {
	Parse(unit string, position domain.SchedulePosition, schedule *domain.Schedule) error
}

type Parser struct {
	unitParser UnitParser
}

func NewParser(unitParser UnitParser) *Parser {
	return &Parser{
		unitParser: unitParser,
	}
}

func (p *Parser) Parse(schedule string) (*domain.Schedule, error) {

	parts := scheduleDelimiter.Split(schedule, scheduleDelimiterLimit)

	if len(parts) != scheduleDelimiterLimit {
		return nil, fmt.Errorf("%w: %s", ErrScheduleFormat, schedule)
	}

	result := &domain.Schedule{}

	for i, part := range parts {
		err := p.unitParser.Parse(part, domain.SchedulePosition(i), result)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrScheduleFormat, err)
		}
	}

	return result, nil
}
