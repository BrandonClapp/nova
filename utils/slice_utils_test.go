package utils

import "testing"

func TestPrepend(t *testing.T) {
	teams := []string{
		"team1", "team2", "team3", "team4",
	}

	order := Prepend(teams, "new")

	if order[0] != "new" {
		t.Error("first element should have been new")
	}
}

func TestPop(t *testing.T) {
	teams := []string{
		"team1", "team2", "team3", "team4",
	}

	popped := PopEnd(&teams)
	if popped != "team4" {
		t.Error("popped not expected value")
	}

	if len(teams) != 3 {
		t.Error("teams unexpected length")
	}
}
