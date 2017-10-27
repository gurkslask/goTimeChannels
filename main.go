package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	fmt.Println("Hello")
}

type timeChannel []timePoint

//Implement sort interface
func (tc timeChannel) Len() int           { return len(tc) }
func (tc timeChannel) Less(i, j int) bool { return tc[i].getDaySecond() < tc[j].getDaySecond() }
func (tc timeChannel) Swap(i, j int)      { tc[i], tc[j] = tc[j], tc[i] }

func (tc timeChannel) checkstate(timeNow time.Time) (bool, error) {
	//timeNow := time.Now()
	daySecondNow := daySecond(timeNow.Hour(), timeNow.Minute(), timeNow.Second())
	var thisWeekDayTimePoints timeChannel
	for _, timepoint := range tc {
		if timepoint.weekDay == timeNow.Weekday() {
			thisWeekDayTimePoints = append(thisWeekDayTimePoints, timepoint)
		}
	}
	sort.Sort(thisWeekDayTimePoints)
	for index, timePoint := range thisWeekDayTimePoints {
		if timePoint.getDaySecond() > daySecondNow {
			return thisWeekDayTimePoints[index-1].state, nil
		}
	}
	return false, nil
}

type timePoint struct {
	weekDay time.Weekday
	hour    int
	minute  int
	second  int
	state   bool
}

func newtimePoint(weekDay, hour, minute, second int, state bool) *timePoint {
	return &timePoint{
		weekDay: time.Weekday(weekDay),
		hour:    hour,
		minute:  minute,
		second:  second,
		state:   state,
	}
}

func daySecond(hour, minute, second int) int {
	return (hour*3600 + minute*60 + second)
}
func (tp timePoint) getDaySecond() int {
	return daySecond(tp.hour, tp.minute, tp.second)
}
