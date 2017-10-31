package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

const settingsname = "settings.json"

func main() {
	var tc timeChannel
	tc.Name = "test"
	tc.Timepoints.newtimePoint(1, 2, 2, 2, false)
	err := tc.toJSON()
	if err != nil {
		log.Fatalf("Error at json %v", err)
	}
	err = fromJSON()
}

func fromJSON() error {
	b, err := ioutil.ReadFile(settingsname)
	if err != nil {
		return err
	}
	fmt.Println(b)
	var tc timeChannel
	err = json.Unmarshal(b, &tc)
	fmt.Println(tc)
	if err != nil {
		return err
	}
	return nil
}

type timeChannels []timeChannel

func (tcs *timeChannels) newtimeChannel(weekDay, hour, minute, second int, state bool, name string) {
	*tcs = append(*tcs, timeChannel{
		Name: name,
	})
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
	Name       string
	Timepoints timePoints
	Output     bool
}

func (tc timeChannel) toJSON() error {
	fmt.Println(tc)
	b, err := json.Marshal(tc)
	fmt.Println(b)

	if err != nil {
		return err
	}
	f, err := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0777)
	//f, err := os.Open(fmt.Sprintf("%v", tc.name))
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

func (tp timePoints) checkstate(timeNow time.Time) (bool, error) {
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
}

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
