package date

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var ErrDateParse = errors.New("date parse error")

// Delimiter between dates
var dateDelimiter = regexp.MustCompile(`,`)

type PeriodParser interface {
	Parse(unit string, min int, max int) ([]string, error)
	IsPeriod(unit string) bool
}

type StepParser interface {
	Parse(unit string, min int, max int) ([]string, error)
	IsStep(unit string) bool
}

type Parser struct {
	periodParser PeriodParser
	stepParser   StepParser
}

func NewParser(periodParser PeriodParser, stepParser StepParser) *Parser {
	return &Parser{
		periodParser: periodParser,
		stepParser:   stepParser,
	}
}

func (p *Parser) Parse(unit string, min int, max int) ([]string, error) {
	var result []string

	if unit == "*" {
		for i := min; i <= max; i++ {
			result = append(result, fmt.Sprintf("%d", i))
		}
		return result, nil
	}

	parts := dateDelimiter.Split(unit, -1)
	for _, part := range parts {
		switch {
		case p.periodParser.IsPeriod(part):
			dates, err := p.periodParser.Parse(part, min, max)
			if err != nil {
				return nil, fmt.Errorf("%w: %w", ErrDateParse, err)
			}
			result = append(result, dates...)

		case p.stepParser.IsStep(part):
			dates, err := p.stepParser.Parse(part, min, max)
			if err != nil {
				return nil, fmt.Errorf("%w: %w", ErrDateParse, err)
			}
			result = append(result, dates...)

		default:
			date, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("%w: %s", ErrDateParse, err)
			}
			result = append(result, fmt.Sprintf("%d", date))
		}

	}

	return result, nil
}
