package series

type ScheduledInstance struct {
	ID       int64
	SeriesID int64
	Start    string
	Duration int64
	Metadata *string

	// virtual means calculated, materialized means additive
	IsMaterialized bool
}
