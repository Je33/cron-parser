package step

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrDateParse = errors.New("invalid date step")
	ErrDateStart = errors.New("invalid start date")
	ErrDateStep  = errors.New("invalid step period")
)

// Step of dates
var dateStep = regexp.MustCompile(`/`)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(unit string, min int, max int) ([]string, error) {
	var result []string
	var err error

	parts := dateStep.Split(unit, 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("%w: %s", ErrDateParse, unit)
	}

	start := min
	if parts[0] != "*" {
		start, err = strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrDateStart, parts[0])
		}
	}

	step, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDateStep, parts[1])
	}

	if start < min || start > max {
		return nil, fmt.Errorf("%w: %d", ErrDateStart, start)
	}

	if step < min || step > max {
		return nil, fmt.Errorf("%w: %d", ErrDateStep, step)
	}

	for i := start; i <= max; i += step {
		result = append(result, fmt.Sprintf("%d", i))
	}

	return result, nil
}

func (p *Parser) IsStep(unit string) bool {
	return dateStep.MatchString(unit)
}
