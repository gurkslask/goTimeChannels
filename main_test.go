package main

import "testing"
import "time"

func TestDaysecond(t *testing.T) {
	got := daySecond(1, 1, 1)
	want := 3661
	if got != want {
		t.Errorf("Got: %v, wanted %v", got, want)
	}
	got = daySecond(-1, 1, 1)
	want = 0
	if got != want {
		t.Errorf("Got: %v, wanted %v", got, want)
	}
	got = daySecond(99999, 1, 1)
	want = 90000
	if got != want {
		t.Errorf("Got: %v, wanted %v", got, want)
	}
}

func TestMin(t *testing.T) {
	got := min(1, 2)
	want := 1
	if got != want {
		t.Errorf("Got: %v, wanted %v", got, want)
	}
	got = min(-1, 2)
	want = -1
	if got != want {
		t.Errorf("Got: %v, wanted %v", got, want)
	}
}

func TestMax(t *testing.T) {
	got := max(1, 2)
	want := 2
	if got != want {
		t.Errorf("Got: %v, wanted %v", got, want)
	}
	got = max(-1, 2)
	want = 2
	if got != want {
		t.Errorf("Got: %v, wanted %v", got, want)
	}
}

func TestTimechannel(t *testing.T) {
	tcs := inittimeChannels()
	tcs.newtimeChannel("test")
	tcs.Tcs["test"].Timepoints.newtimePoint(1, 2, 2, 2, false)
	tcs.Tcs["test"].Timepoints.newtimePoint(1, 1, 1, 1, true)
	tcs.Tcs["test"].Timepoints.newtimePoint(1, 3, 3, 3, true)
	tcs.Tcs["test"].Timepoints.newtimePoint(2, 0, 0, 1, true)

	_ = tcs.Tcs["test"].checkstate(time.Date(2017, 10, 23, 2, 2, 3, 0, time.UTC))
	want := false
	if tcs.Tcs["test"].output != want {
		t.Errorf("Got %v, wanted %v", tcs.Tcs["test"].output, want)
	}

	_ = tcs.Tcs["test"].checkstate(time.Date(2017, 10, 23, 1, 1, 3, 0, time.UTC))
	want = true
	if tcs.Tcs["test"].output != want {
		t.Errorf("Got %v, wanted %v", tcs.Tcs["test"].output, want)
	}

	_ = tcs.Tcs["test"].checkstate(time.Date(2017, 10, 23, 23, 1, 3, 0, time.UTC))
	want = true
	if tcs.Tcs["test"].output != want {
		t.Errorf("Got %v, wanted %v", tcs.Tcs["test"].output, want)
	}

	_ = tcs.Tcs["test"].checkstate(time.Date(2017, 10, 24, 23, 1, 3, 0, time.UTC))
	want = true
	if tcs.Tcs["test"].output != want {
		t.Errorf("Got %v, wanted %v", tcs.Tcs["test"].output, want)
	}
}
