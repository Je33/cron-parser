package command

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(unit string) (string, error) {
	return unit, nil
}
