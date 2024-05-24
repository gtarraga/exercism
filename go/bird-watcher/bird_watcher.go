package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var TotalBirdCount int

	for i := 0; i < len(birdsPerDay); i++ {
		TotalBirdCount += birdsPerDay[i]
	}
	return TotalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var BirdsInWeek int

	for i := 0; i < 7; i++ {
		BirdsInWeek += birdsPerDay[i+(week-1)*7]
	}
	return BirdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
