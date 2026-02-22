package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	res := 0
	for _, v := range birdsPerDay {
		res = res + v
	}
	return res
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	slice := birdsPerDay[(week-1)*7 : (week * 7)]
	res := TotalBirdCount(slice)
	return res
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	present := true
	for i, v := range birdsPerDay {
		switch {
		case present:
			birdsPerDay[i] = v + 1
			present = false
		default:
			present = true
		}
	}
	return birdsPerDay
}
