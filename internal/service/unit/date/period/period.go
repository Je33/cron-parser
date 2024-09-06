package period

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrDateRange = errors.New("invalid range")
	ErrDateEnd   = errors.New("invalid end date")
	ErrDateStart = errors.New("invalid start date")
)

// Range of dates
var datePeriod = regexp.MustCompile(`-`)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(unit string, min int, max int) ([]string, error) {
	var result []string

	parts := datePeriod.Split(unit, 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("%w: %s", ErrDateRange, unit)
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDateStart, parts[0])
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDateEnd, parts[1])
	}

	if start < min || start > max {
		return nil, fmt.Errorf("%w: %d", ErrDateStart, start)
	}

	if end < min || end > max {
		return nil, fmt.Errorf("%w: %d", ErrDateEnd, end)
	}

	for i := start; i <= end; i++ {
		result = append(result, fmt.Sprintf("%d", i))
	}

	return result, nil
}

func (p *Parser) IsPeriod(unit string) bool {
	return datePeriod.MatchString(unit)
}
