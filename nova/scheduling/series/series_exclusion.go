package series

type SeriesExclusion struct {
	SeriesID int64

	// start date to exclude from the series
	// must match the start date & time generated from the extrapolation algorithm
	Date string
}
