package main

import (
	"log"
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
	electronrun()

}
