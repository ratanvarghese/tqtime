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
	{1972, time.July, 19, "28M 3"},
	{1972, time.June, 22, "01M 3"},
	{1972, time.June, 21, "28L 3"},
	{1972, time.May, 25, "01L 3"},
	{1972, time.May, 24, "28K 3"},
	{1972, time.April, 27, "01K 3"},
	{1972, time.April, 26, "28J 3"},
	{1972, time.March, 30, "01J 3"},
	{1972, time.March, 29, "28I 3"},
	{1972, time.March, 2, "01I 3"},
	{1972, time.March, 1, "28H 3"},
	{1972, time.February, 29, "ALD 3"},
	{1972, time.February, 28, "27H 3"},
	{1972, time.February, 2, "01H 3"},
	{1972, time.February, 1, "28G 3"},
	{1972, time.January, 5, "01G 3"},
	{1972, time.January, 4, "28F 3"},
	{1971, time.December, 8, "01F 3"},
	{1971, time.December, 7, "28E 3"},
	{1971, time.November, 10, "01E 3"},
	{1971, time.November, 9, "28D 3"},
	{1971, time.October, 13, "01D 3"},
	{1971, time.October, 12, "28C 3"},
	{1971, time.September, 15, "01C 3"},
	{1971, time.September, 14, "28B 3"},
	{1971, time.August, 18, "01B 3"},
	{1971, time.August, 17, "28A 3"},
	{1971, time.July, 21, "01A 3"},
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
		actual := ShortDate(gt.Year(), gt.YearDay())
		expected := tt.output
		if actual != expected {
			t.Errorf("Short date %s; expected %s; actual %s.", gt.Format("2006-01-02"), expected, actual)
		}
	}
}

var longTests = []struct {
	gYear  int
	gMonth time.Month
	gDay   int
	output string
}{
	{2000, time.February, 29, "Aldrin Day, 31 After Tranquility"},
	{1972, time.July, 19, "Thursday, 28 Mendel, 3 After Tranquility"},
	{1972, time.June, 22, "Friday, 1 Mendel, 3 After Tranquility"},
	{1972, time.June, 21, "Thursday, 28 Lavoisier, 3 After Tranquility"},
	{1972, time.May, 25, "Friday, 1 Lavoisier, 3 After Tranquility"},
	{1972, time.May, 24, "Thursday, 28 Kepler, 3 After Tranquility"},
	{1972, time.April, 27, "Friday, 1 Kepler, 3 After Tranquility"},
	{1972, time.April, 26, "Thursday, 28 Jung, 3 After Tranquility"},
	{1972, time.March, 30, "Friday, 1 Jung, 3 After Tranquility"},
	{1972, time.March, 29, "Thursday, 28 Imhotep, 3 After Tranquility"},
	{1972, time.March, 2, "Friday, 1 Imhotep, 3 After Tranquility"},
	{1972, time.March, 1, "Thursday, 28 Hippocrates, 3 After Tranquility"},
	{1972, time.February, 29, "Aldrin Day, 3 After Tranquility"},
	{1972, time.February, 28, "Wednesday, 27 Hippocrates, 3 After Tranquility"},
	{1972, time.February, 2, "Friday, 1 Hippocrates, 3 After Tranquility"},
	{1972, time.February, 1, "Thursday, 28 Galileo, 3 After Tranquility"},
	{1972, time.January, 5, "Friday, 1 Galileo, 3 After Tranquility"},
	{1972, time.January, 4, "Thursday, 28 Faraday, 3 After Tranquility"},
	{1971, time.December, 8, "Friday, 1 Faraday, 3 After Tranquility"},
	{1971, time.December, 7, "Thursday, 28 Einstein, 3 After Tranquility"},
	{1971, time.November, 10, "Friday, 1 Einstein, 3 After Tranquility"},
	{1971, time.November, 9, "Thursday, 28 Darwin, 3 After Tranquility"},
	{1971, time.October, 13, "Friday, 1 Darwin, 3 After Tranquility"},
	{1971, time.October, 12, "Thursday, 28 Copernicus, 3 After Tranquility"},
	{1971, time.September, 15, "Friday, 1 Copernicus, 3 After Tranquility"},
	{1971, time.September, 14, "Thursday, 28 Brahe, 3 After Tranquility"},
	{1971, time.August, 18, "Friday, 1 Brahe, 3 After Tranquility"},
	{1971, time.August, 17, "Thursday, 28 Archimedes, 3 After Tranquility"},
	{1971, time.July, 27, "Thursday, 7 Archimedes, 3 After Tranquility"},
	{1971, time.July, 26, "Wednesday, 6 Archimedes, 3 After Tranquility"},
	{1971, time.July, 25, "Tuesday, 5 Archimedes, 3 After Tranquility"},
	{1971, time.July, 24, "Monday, 4 Archimedes, 3 After Tranquility"},
	{1971, time.July, 23, "Sunday, 3 Archimedes, 3 After Tranquility"},
	{1971, time.July, 22, "Saturday, 2 Archimedes, 3 After Tranquility"},
	{1971, time.July, 21, "Friday, 1 Archimedes, 3 After Tranquility"},
	{1970, time.July, 20, "Armstrong Day, 1 After Tranquility"},
	{1969, time.July, 21, "Friday, 1 Archimedes, 1 After Tranquility"},
	{1969, time.July, 20, "Moon Landing Day"},
	{1969, time.July, 19, "Thursday, 28 Mendel, 1 Before Tranquility"},
	{1968, time.July, 20, "Armstrong Day, 2 Before Tranquility"},
	{1968, time.July, 19, "Thursday, 28 Mendel, 2 Before Tranquility"},
	{1968, time.February, 29, "Aldrin Day, 2 Before Tranquility"},
	{1967, time.July, 20, "Armstrong Day, 3 Before Tranquility"},
	{1967, time.July, 19, "Thursday, 28 Mendel, 3 Before Tranquility"},
	{1900, time.February, 29, "Thursday, 28 Hippocrates, 70 Before Tranquility"},
}

func TestLongDate(t *testing.T) {
	for _, tt := range longTests {
		y := tt.gYear
		m := tt.gMonth
		d := tt.gDay
		gt := time.Date(y, m, d, 1, 1, 1, 1, time.UTC)
		actual := LongDate(gt.Year(), gt.YearDay())
		expected := tt.output
		if actual != expected {
			t.Errorf("Long date %s; expected %s; actual %s.", gt.Format("2006-01-02"), expected, actual)
		}
	}
}

var TranquilityBoundaryTests = []struct {
	minute int
	second int
	output bool
}{
	{18, 2, false},
	{18, 1, true},
	{17, 59, true},
	{17, 58, true},
}

func TestTranquilityBoundary(t *testing.T) {
	for _, tt := range TranquilityBoundaryTests {
		min := tt.minute
		sec := tt.second
		gt := time.Date(1969, time.July, 20, 20, min, sec, 0, time.UTC)
		actual := IsBeforeTranquility(gt.Year(), gt.YearDay(), gt.Hour(), gt.Minute(), gt.Second(), gt.Nanosecond()*1000000)
		expected := tt.output
		if actual != expected {
			t.Errorf("Time %s on Moon Landing Day on wrong side of Tranquility Boundary.", gt.Format("15:04:05"))
		}
	}
}

func TestMonthSpecialDay(t *testing.T) {
	gt1 := time.Date(1969, time.July, 20, 20, 1, 1, 1, time.UTC)
	if Month(gt1.Year(), gt1.YearDay()) != SpecialDay {
		t.Error("Month did not return SpecialDay on Moon Landing Day.")
	}
	gt2 := time.Date(1968, time.July, 20, 20, 1, 1, 1, time.UTC)
	if Month(gt2.Year(), gt2.YearDay()) != SpecialDay {
		t.Error("Month did not return SpecialDay on Armstrong Day.")
	}
	gt3 := time.Date(2000, time.February, 29, 1, 1, 1, 1, time.UTC)
	if Month(gt3.Year(), gt3.YearDay()) != SpecialDay {
		t.Error("Month did not return SpecialDay on Aldrin Day.")
	}
}

func TestWeekdayNameInvalid(t *testing.T) {
	if WeekdayName(-99) != "" {
		t.Error("WeekdayName did not return blank with invalid value.")
	}
}

func TestWeekdaySpecial(t *testing.T) {
	gt := time.Date(2000, time.February, 29, 1, 1, 1, 1, time.UTC)
	if Weekday(gt.Year(), gt.YearDay()) != SpecialWeekday {
		t.Error("Weekday did not return SpecialWeekday on Aldrin Day.")
	}
}

func TestDayCodeOrdinary(t *testing.T) {
	if DayCode(11) != "11" {
		t.Error("DayCode did not return '11' with input 11.")
	}
}

func TestDayNameOrdinary(t *testing.T) {
	if DayName(11) != "11" {
		t.Error("DayName did not return '11' with input 11.")
	}
}

func TestMonthNameInvalid(t *testing.T) {
	if TqMonth(-99).String() != "" {
		t.Error("TqMonth.String did not return blank with invalid value.")
	}
}

func TestMonthLetterInvalid(t *testing.T) {
	if MonthLetter(-99) != "" {
		t.Error("MonthLetter did not return blank with invalid value.")
	}
}
