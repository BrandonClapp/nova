package series

import (
	"testing"
)

func TestGenerateMonthlySeriesWithEnd(t *testing.T) {
	end := "2022-10-01T00:00:00Z"

	series := Series{
		ID:        1,
		Start:     "2021-10-01T00:00:00Z",
		End:       &end,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "monthly",
	}

	instances, _ := series.GenerateInstances()
	expected := 13

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateMonthlySeriesWithoutEnd(t *testing.T) {
	series := Series{
		ID:        1,
		Start:     "2021-10-01T00:00:00Z",
		End:       nil,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "monthly",
	}

	instances, _ := series.GenerateInstances()
	expected := 61

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateDailySeriesWithEnd(t *testing.T) {
	end := "2022-10-01T00:00:00Z"

	series := Series{
		ID:        1,
		Start:     "2021-10-01T00:00:00Z",
		End:       &end,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "daily",
	}

	instances, _ := series.GenerateInstances()
	expected := 366

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateDailySeriesWithoutEnd(t *testing.T) {
	series := Series{
		ID:        1,
		Start:     "2021-10-01T00:00:00Z",
		End:       nil,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "daily",
	}

	instances, _ := series.GenerateInstances()
	expected := 1827

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateDailySeriesWeekdays(t *testing.T) {
	end := "2022-10-04T00:00:00Z"

	series := Series{
		ID:        1,
		Start:     "2021-10-04T00:00:00Z",
		End:       &end,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "daily[monday, tuesday, wednesday, thursday, friday]",
	}

	instances, _ := series.GenerateInstances()
	expected := 262

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateDailySeriesWeekdaysWithoutEnd(t *testing.T) {
	series := Series{
		ID:        1,
		Start:     "2021-10-04T00:00:00Z",
		End:       nil,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "daily[monday, tuesday, wednesday, thursday, friday]",
	}

	instances, _ := series.GenerateInstances()
	expected := 1305

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateWeeklySeries(t *testing.T) {
	end := "2022-10-04T00:00:00Z"

	series := Series{
		ID:        1,
		Start:     "2021-10-04T00:00:00Z",
		End:       &end,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "weekly",
	}

	instances, _ := series.GenerateInstances()
	expected := 53

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateWeeklySeriesWithoutEnd(t *testing.T) {
	series := Series{
		ID:        1,
		Start:     "2021-10-04T00:00:00Z",
		End:       nil,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "weekly",
	}

	instances, _ := series.GenerateInstances()
	expected := 261

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateBiweeklySeries(t *testing.T) {
	end := "2022-10-04T00:00:00Z"

	series := Series{
		ID:        1,
		Start:     "2021-10-04T00:00:00Z",
		End:       &end,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "biweekly",
	}

	instances, _ := series.GenerateInstances()
	expected := 27

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateBiweeklySeriesWithoutEnd(t *testing.T) {
	series := Series{
		ID:        1,
		Start:     "2021-10-04T00:00:00Z",
		End:       nil,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "biweekly",
	}

	instances, _ := series.GenerateInstances()
	expected := 131

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateYearlySeries(t *testing.T) {
	end := "2022-10-04T00:00:00Z"

	series := Series{
		ID:        1,
		Start:     "2021-10-04T00:00:00Z",
		End:       &end,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "yearly",
	}

	instances, _ := series.GenerateInstances()
	expected := 2

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}

func TestGenerateYearlySeriesWithoutEnd(t *testing.T) {
	series := Series{
		ID:        1,
		Start:     "2021-10-04T00:00:00Z",
		End:       nil,
		Duration:  3600,
		Metadata:  nil,
		Frequency: "yearly",
	}

	instances, _ := series.GenerateInstances()
	expected := 6

	if len(instances) != expected {
		t.Errorf("expected %d but got %d", expected, len(instances))
	}
}
