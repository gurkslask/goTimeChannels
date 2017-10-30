package main

func daySecond(hour, minute, second int) int {
	return max(0, min(90000, (hour*3600+minute*60+second)))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
