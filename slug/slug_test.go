package slug

import (
	"fmt"
	"testing"
)

func TestMakeSlug(t *testing.T) {
	testCases := []struct {
		title    string
		expected string
	}{
		{
			title:    "This is a tournament",
			expected: "this-is-a-tournament",
		},
		{
			title:    "Rocket League 3v3 Summer Kickoff",
			expected: "rocket-league-3v3-summer-kickoff",
		},
	}

	for _, c := range testCases {
		got := MakeSlug(c.title)

		fmt.Println(got)
		if got != c.expected {
			t.Errorf("expected %s but got %s", c.expected, got)
		}
	}
}
