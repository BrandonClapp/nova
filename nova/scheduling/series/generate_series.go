package series

import (
	"regexp"
	"strings"
	"time"
)

func (series *Series) GenerateInstances() ([]ScheduledInstance, error) {
	start, err := time.Parse(time.RFC3339, series.Start)

	if err != nil {
		// every series must have a valid start date
		return nil, err
	}

	// default end value for the series
	// project up to 5 years of instances
	end := start.AddDate(5, 0, 0)

	if series.End != nil {
		// override end value with the end provided on the series
		end, err = time.Parse(time.RFC3339, *series.End)

		if err != nil {
			return nil, err
		}
	}

	instances := []ScheduledInstance{}

	if series.Frequency == "monthly" {
		next := start

		for next.Before(end) || next == end {
			instances = append(instances, getInstance(series, next))
			next = bumpMonth(next)
		}
	}

	if series.Frequency == "weekly" {
		next := start

		for next.Before(end) || next == end {
			instances = append(instances, getInstance(series, next))
			next = bumpWeek(next)
		}
	}

	if series.Frequency == "biweekly" {
		next := start

		for next.Before(end) || next == end {
			instances = append(instances, getInstance(series, next))
			next = bumpBiweek(next)
		}
	}

	if series.Frequency == "yearly" {
		next := start

		for next.Before(end) || next == end {
			instances = append(instances, getInstance(series, next))
			next = bumpYear(next)
		}
	}

	if strings.Contains(series.Frequency, "daily") {
		whitelist := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

		r := regexp.MustCompile(`daily\[(.*)\]`)
		matches := r.FindStringSubmatch(series.Frequency)

		if len(matches) == 2 {
			days := strings.Split(strings.ReplaceAll(matches[1], " ", ""), ",")
			whitelist = days
		}

		next := start

		for next.Before(end) || next == end {
			// if the day of the week is in whitelist, add to instance
			day := strings.ToLower(next.Weekday().String())

			for _, val := range whitelist {
				if val == day {
					instances = append(instances, getInstance(series, next))
				}
			}

			next = bumpDay(next)
		}
	}

	return instances, nil
}

func bumpDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func bumpWeek(t time.Time) time.Time {
	return t.AddDate(0, 0, 7)
}

func bumpBiweek(t time.Time) time.Time {
	return t.AddDate(0, 0, 14)
}

func bumpMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, 0)
}

func bumpYear(t time.Time) time.Time {
	return t.AddDate(1, 0, 0)
}

func getInstance(series *Series, start time.Time) ScheduledInstance {
	return ScheduledInstance{
		ID:             -1,
		SeriesID:       series.ID,
		Start:          start.Format(time.RFC3339),
		Duration:       series.Duration,
		Metadata:       series.Metadata,
		IsMaterialized: false,
	}
}
