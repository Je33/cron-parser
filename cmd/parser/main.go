package main

import (
	"cron-parser/internal/service/schedule"
	"cron-parser/internal/service/unit"
	"cron-parser/internal/service/unit/command"
	"cron-parser/internal/service/unit/date"
	"cron-parser/internal/service/unit/date/period"
	"cron-parser/internal/service/unit/date/step"
	"flag"
	"fmt"
	"strings"
)

const argMessage = `please specify cron schedule, format: "[minute] [hour] [day of month] [month] [day of week] [command]", example: "*/15 1,12 1-4 * * /bin/command -with -parameters"`

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println(argMessage)
		return
	}

	periodParser := period.NewParser()
	stepParser := step.NewParser()

	dateParser := date.NewParser(periodParser, stepParser)
	commandParser := command.NewParser()

	unitParser := unit.NewParser(dateParser, commandParser)

	parser := schedule.NewParser(unitParser)

	res, err := parser.Parse(args[0])
	if err != nil {
		fmt.Println(err)
		fmt.Println(argMessage)
		return
	}

	fmt.Printf("minutes:       %s\n", strings.Join(res.Minute, " "))
	fmt.Printf("hours:         %s\n", strings.Join(res.Hour, " "))
	fmt.Printf("days of month: %s\n", strings.Join(res.DayMonth, " "))
	fmt.Printf("months:        %s\n", strings.Join(res.Month, " "))
	fmt.Printf("days of week:  %s\n", strings.Join(res.DayWeek, " "))
	fmt.Printf("command:       %s\n", res.Command)
}
