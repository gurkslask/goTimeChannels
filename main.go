package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

const settingsname = "settings.json"

func main() {
	tcs := inittimeChannels()
	tcs.newtimeChannel("test")
	tcs.Tcs["test"].Timepoints.newtimePoint(1, 2, 2, 2, false)
	tcs.Tcs["test"].Timepoints.newtimePoint(1, 3, 2, 2, false)
	err := tcs.toJSON()
	if err != nil {
		log.Fatalf("%v", err)
	}

}

type timeChannels struct {
	Tcs map[string]*timeChannel
}

func inittimeChannels() *timeChannels {
	return &timeChannels{
		Tcs: make(map[string]*timeChannel),
	}
}

func (tcs *timeChannels) newtimeChannel(name string) {
	tcs.Tcs[name] = newtimeChannel(name)
}

func (tcs *timeChannels) fromJSON() error {
	b, err := ioutil.ReadFile("settings.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &tcs)
	if err != nil {
		return err
	}
	return nil
}
func (tcs timeChannels) toJSON() error {
	b, err := json.Marshal(tcs)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(settingsname, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}

type timeChannel struct {
	Name       string     `json:"name"`
	Timepoints timePoints `json:"timepoints"`
	output     bool
}

func newtimeChannel(name string) *timeChannel {
	return &timeChannel{Name: name}
}

func (tc *timeChannel) checkstate(timeNow time.Time) error {
	daySecondNow := daySecond(timeNow.Hour(), timeNow.Minute(), timeNow.Second())
	var thisWeekDayTimePoints timePoints
	for _, timepoint := range tc.Timepoints {
		if timepoint.WeekDay == timeNow.Weekday() {
			thisWeekDayTimePoints = append(thisWeekDayTimePoints, timepoint)
		}
	}
	sort.Sort(thisWeekDayTimePoints)
	for index, timePoint := range thisWeekDayTimePoints {
		//  fmt.Printf("%v > %v \n", timePoint.getDaySecond(), daySecondNow)
		if timePoint.getDaySecond() > daySecondNow {
			tc.output = thisWeekDayTimePoints[index-1].State
			return nil
		}
	}
	if thisWeekDayTimePoints[len(thisWeekDayTimePoints)-1].getDaySecond() < daySecondNow {
		tc.output = thisWeekDayTimePoints[len(thisWeekDayTimePoints)-1].State
		return nil
	}
	return nil
}

type timePoints []timePoint

//Implement sort interface
func (tp timePoints) Len() int           { return len(tp) }
func (tp timePoints) Less(i, j int) bool { return tp[i].getDaySecond() < tp[j].getDaySecond() }
func (tp timePoints) Swap(i, j int)      { tp[i], tp[j] = tp[j], tp[i] }

func (tp *timePoints) newtimePoint(weekDay, hour, minute, second int, state bool) {
	*tp = append(*tp, timePoint{
		WeekDay: time.Weekday(weekDay),
		Hour:    hour,
		Minute:  minute,
		Second:  second,
		State:   state,
	})
}

/*func (tp timePoints) checkstate(timeNow time.Time) (bool, error) {
	//timeNow := time.Now()
	daySecondNow := daySecond(timeNow.Hour(), timeNow.Minute(), timeNow.Second())
	var thisWeekDayTimePoints timePoints
	for _, timepoint := range tp {
		if timepoint.WeekDay == timeNow.Weekday() {
			thisWeekDayTimePoints = append(thisWeekDayTimePoints, timepoint)
		}
	}
	sort.Sort(thisWeekDayTimePoints)
	for index, timePoint := range thisWeekDayTimePoints {
		//  fmt.Printf("%v > %v \n", timePoint.getDaySecond(), daySecondNow)
		if timePoint.getDaySecond() > daySecondNow {
			return thisWeekDayTimePoints[index-1].State, nil
		}
	}
	if thisWeekDayTimePoints[len(thisWeekDayTimePoints)-1].getDaySecond() < daySecondNow {
		return thisWeekDayTimePoints[len(thisWeekDayTimePoints)-1].State, nil
	}
	return false, nil
}*/

type timePoint struct {
	//WeekDay time.Weekday "json:'weekday'"
	WeekDay time.Weekday `json:"weekday"`
	Hour    int          `json:"hour"`
	Minute  int          `json:"minute"`
	Second  int          `json:"second"`
	State   bool         `json:"state"`
}

func newtimePoint(weekDay, hour, minute, second int, state bool) timePoint {
	return timePoint{
		WeekDay: time.Weekday(weekDay),
		Hour:    hour,
		Minute:  minute,
		Second:  second,
		State:   state,
	}
}

func (tp timePoint) getDaySecond() int {
	return daySecond(tp.Hour, tp.Minute, tp.Second)
}
