package check

import (
	"fmt"
	"testing"
	"time"
)

func TestCheckDayTime(t *testing.T) {
	timeNow := time.Now()
	hourMinute := fmt.Sprintf("%d:%d", timeNow.Hour(), timeNow.Minute())

	yesterday := timeNow.Add(time.Hour * -25)
	runYesterday := DayTime(hourMinute, yesterday.Format(time.RFC3339))
	if !runYesterday {
		t.Errorf("Check daytime yesterday failed")
	}

	soon := timeNow.Add(time.Minute * -1)
	runSoon := DayTime(hourMinute, soon.Format(time.RFC3339))
	if runSoon {
		t.Errorf("Check daytime soon failed")
	}

	tomorrow := timeNow.Add(time.Minute * 1)
	runTomorrow := DayTime(hourMinute, tomorrow.Format(time.RFC3339))
	if runTomorrow {
		t.Errorf("Check daytime tomorrow failed")
	}
}
