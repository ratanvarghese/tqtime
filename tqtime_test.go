package tqtime

import (
	"testing"
	"time"
)

var shortTests = []struct {
	gYear  int
	gMonth time.Month
	gDay   int
	output string
}{
	{2000, time.February, 29, "ALD 31"},
	{1972, time.February, 29, "ALD 3"},
	{1970, time.July, 20, "ARM 1"},
	{1969, time.July, 21, "01A 1"},
	{1969, time.July, 20, "MNL 0"},
	{1969, time.July, 19, "28M -1"},
	{1968, time.July, 20, "ARM -2"},
	{1968, time.July, 19, "28M -2"},
	{1968, time.February, 29, "ALD -2"},
	{1967, time.July, 20, "ARM -3"},
	{1967, time.July, 19, "28M -3"},
	{1900, time.February, 29, "28H -70"},
}

func TestShortDate(t *testing.T) {
	for _, tt := range shortTests {
		y := tt.gYear
		m := tt.gMonth
		d := tt.gDay
		gt := time.Date(y, m, d, 1, 1, 1, 1, time.UTC)
		actual := ShortDate(gt.Unix())
		expected := tt.output
		if actual != expected {
			//t.Error("Short date", y, "-", m, "-", d, "expected '", expected, "', actual '", actual, "'.")
			t.Errorf("Short date %s; expected %s; actual %s.", gt.Format("2006-01-02"), expected, actual)
		}
	}
}
