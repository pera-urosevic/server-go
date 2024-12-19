package check

import (
	"server/api/micron/log"
	"strconv"
	"strings"
	"time"
)

func WeekDayTime(weekDay string, dayTime string, lastRun string) bool {
	// last run
	timeLastRun, err := time.Parse(time.RFC3339, lastRun)
	if err != nil {
		log.Log("Error parsing lastRun in DayTime", err)
		return false
	}
	// now
	timeNow := time.Now()
	// next run
	parts := strings.Split(dayTime, ":")
	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Log("Error parsing time hour in DayTime", err)
		return false
	}
	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Log("Error parsing time minute in DayTime", err)
		return false
	}
	timeNextRun := time.Date(timeLastRun.Year(), timeLastRun.Month(), timeLastRun.Day(), hour, minute, 0, 0, time.Local)
	for {
		timeNextRun = timeNextRun.Add(time.Hour * 24)
		if timeNextRun.Weekday().String() == weekDay {
			break
		}
	}
	// skip if already ran today
	if timeLastRun.After(timeNextRun) {
		return false
	}
	// is it the time to run
	return timeNextRun.Before(timeNow)
}
