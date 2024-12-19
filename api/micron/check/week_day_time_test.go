package check

import (
	"fmt"
	"testing"
	"time"
)

func TestCheckWeekDayTime(t *testing.T) {
	timeNow := time.Now()
	hourMinute := fmt.Sprintf("%d:%d", timeNow.Hour(), timeNow.Minute())
	weekDay := timeNow.Weekday().String()

	lastweek := timeNow.Add(time.Hour * 24 * -7)
	runLastWeek := WeekDayTime(weekDay, hourMinute, lastweek.Format(time.RFC3339))
	if !runLastWeek {
		t.Errorf("Check daytime lastweek failed")
	}

	soon := timeNow.Add(time.Minute * -1)
	runSoon := WeekDayTime(weekDay, hourMinute, soon.Format(time.RFC3339))
	if runSoon {
		t.Errorf("Check daytime soon failed")
	}

	tomorrow := timeNow.Add(time.Minute * 1)
	runTomorrow := WeekDayTime(weekDay, hourMinute, tomorrow.Format(time.RFC3339))
	if runTomorrow {
		t.Errorf("Check daytime tomorrow failed")
	}
}
