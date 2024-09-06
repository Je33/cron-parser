package schedule_test

import (
	"cron-parser/internal/service/schedule"
	"cron-parser/internal/service/unit"
	"cron-parser/internal/service/unit/command"
	"cron-parser/internal/service/unit/date"
	"cron-parser/internal/service/unit/date/period"
	"cron-parser/internal/service/unit/date/step"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {

	periodParser := period.NewParser()
	stepParser := step.NewParser()

	dateParser := date.NewParser(periodParser, stepParser)
	commandParser := command.NewParser()

	unitParser := unit.NewParser(dateParser, commandParser)

	parser := schedule.NewParser(unitParser)

	t.Run("ParseSuccess", func(t *testing.T) {

		cron := "*/15 1-12,15,20-23 1-4 * * /bin/command -with -parameters"

		res, err := parser.Parse(cron)
		assert.NoError(t, err)

		assert.Equal(t, []string{"0", "15", "30", "45"}, res.Minute)
		assert.Equal(t, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "15", "20", "21", "22", "23"}, res.Hour)
		assert.Equal(t, []string{"1", "2", "3", "4"}, res.DayMonth)
		assert.Equal(t, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}, res.Month)
		assert.Equal(t, []string{"0", "1", "2", "3", "4", "5", "6"}, res.DayWeek)
		assert.Equal(t, "/bin/command -with -parameters", res.Command)

		t.Logf("%+v", res)
	})

	t.Run("ParsePeriodError", func(t *testing.T) {

		cron := "*/15 1-12,15,20-25 * * * /bin/command -with -parameters"

		_, err := parser.Parse(cron)
		assert.ErrorIs(t, err, period.ErrDateEnd)
	})

	t.Run("ParseStepError", func(t *testing.T) {

		cron := "*/61 1-12,15,20-21 * * * /bin/command -with -parameters"

		_, err := parser.Parse(cron)
		assert.ErrorIs(t, err, step.ErrDateStep)
	})

	t.Run("ParseScheduleError", func(t *testing.T) {

		cron := "as fa 1r1 asd r1f asd qwe"

		_, err := parser.Parse(cron)
		assert.ErrorIs(t, err, schedule.ErrScheduleFormat)
	})
}
