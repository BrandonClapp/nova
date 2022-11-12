package series

type Series struct {
	ID        int64   `json:"id"`
	Start     string  `json:"start"`     // when the series starts
	End       *string `json:"end"`       // when the series ends
	Duration  int64   `json:"duration"`  // duration in seconds
	Frequency string  `json:"frequency"` // the frequently to extrapolate instances. "daily" | "weekly" | "biweekly" | "monthly"
	Metadata  *string `json:"metadata"`  // arbitrary info about the series
}
