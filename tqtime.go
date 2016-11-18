//g_* == Gregorian
package tqtime

import (
	"fmt"
	"strconv"
	"time"
)

type TqWeekday int

const (
	SpecialWeekday TqWeekday = iota
	Friday
	Saturday
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
)

type TqMonth int

const (
	SpecialDay TqMonth = iota
	Archimedes
	Brahe
	Copernicus
	Darwin
	Einstein
	Faraday
	Galileo
	Hippocrates
	Imhotep
	Jung
	Kepler
	Lavoisier
	Mendel
)

const ArmstrongDay int = -1
const AldrinDay int = -2
const MoonLandingDay int = -3

const mlYear int = 1969
const mlMonth time.Month = time.July
const mlDay int = 20
const mlYearDay int = 201

const commonYearLen int = 365
const tqMonthLen int = 28

func isGregorianLeapYear(g time.Time) bool {
	y := g.Year()
	return (y%400 == 0) || (y%4 == 0 && y%100 != 0)
}

func isMoonLandingDay(g time.Time) bool {
	g_y, g_m, g_d := g.Date()
	return g_y == mlYear && g_m == mlMonth && g_d == mlDay
}

func isAfterArmstrongDay(g time.Time) bool {
	_, g_m, g_d := g.Date() //Avoid YearDay due to leap year
	return (g_m > mlMonth) || (g_m == mlMonth && g_d > mlDay)
}

func IsBeforeTranquility(unixTime int64) bool {
	const unixMoonLanding int64 = -14182919
	return unixTime < unixMoonLanding
}

func Year(unixTime int64) int {
	g := time.Unix(unixTime, 0)
	if isMoonLandingDay(g) {
		return 0
	}
	yearDiff := g.Year() - mlYear
	if isAfterArmstrongDay(g) {
		yearDiff += 1
	}
	if yearDiff < 1 {
		yearDiff -= 1
	}
	return yearDiff
}

func Month(unixTime int64) TqMonth {
	g := time.Unix(unixTime, 0)
	if isMoonLandingDay(g) {
		return SpecialDay //Moon Landing Day
	}
	yd := YearDay(unixTime)
	if isGregorianLeapYear(g) {
		const leapDay int = tqMonthLen * int(Hippocrates)
		if yd == leapDay {
			return SpecialDay //Aldrin Day
		} else if yd > leapDay {
			yd -= 1
		}
	}
	if yd == commonYearLen {
		return SpecialDay //Armstrong Day
	}
	return TqMonth(((yd - 1) / tqMonthLen) + 1)
}

func Day(unixTime int64) int {
	g := time.Unix(unixTime, 0)
	if isMoonLandingDay(g) {
		return MoonLandingDay
	}
	yd := YearDay(unixTime)
	if isGregorianLeapYear(g) {
		const leapDay int = tqMonthLen * int(Hippocrates)
		if yd == leapDay {
			return AldrinDay
		} else if yd > leapDay {
			yd -= 1
		}
	}
	if yd == commonYearLen {
		return ArmstrongDay
	} else if yd%tqMonthLen == 0 {
		return tqMonthLen
	} else {
		return (yd % tqMonthLen)
	}
}

func YearDay(unixTime int64) int {
	g := time.Unix(unixTime, 0)
	yearLen := commonYearLen
	armstrongYearDay := mlYearDay
	if isGregorianLeapYear(g) {
		yearLen += 1
		armstrongYearDay += 1
	}
	diff := g.YearDay() - armstrongYearDay
	if diff > 0 || yearLen == diff {
		return diff
	} else {
		return yearLen - diff
	}
}

func Weekday(unixTime int64) TqWeekday {
	d := Day(unixTime)
	switch {
	case d < 0:
		return SpecialWeekday
	case d%7 == 0:
		return Thursday
	default:
		return TqWeekday(d % 7)
	}
}

func MonthName(m TqMonth) string {
	switch m {
	case Archimedes:
		return "Archimedes"
	case Brahe:
		return "Brahe"
	case Copernicus:
		return "Copernicus"
	case Darwin:
		return "Darwin"
	case Einstein:
		return "Einstein"
	case Faraday:
		return "Faraday"
	case Galileo:
		return "Galileo"
	case Hippocrates:
		return "Hippocrates"
	case Imhotep:
		return "Imhotep"
	case Jung:
		return "Jung"
	case Kepler:
		return "Kepler"
	case Lavoisier:
		return "Lavoisier"
	case Mendel:
		return "Mendel"
	default:
		return ""
	}
}

func MonthLetter(m TqMonth) string {
	switch m {
	case Archimedes:
		return "A"
	case Brahe:
		return "B"
	case Copernicus:
		return "C"
	case Darwin:
		return "D"
	case Einstein:
		return "E"
	case Faraday:
		return "F"
	case Galileo:
		return "G"
	case Hippocrates:
		return "H"
	case Imhotep:
		return "I"
	case Jung:
		return "J"
	case Kepler:
		return "K"
	case Lavoisier:
		return "L"
	case Mendel:
		return "M"
	default:
		return ""
	}
}

func DayName(d int) string {
	switch d {
	case ArmstrongDay:
		return "Armstrong Day"
	case AldrinDay:
		return "Aldrin Day"
	case MoonLandingDay:
		return "Moon Landing Day"
	default:
		return strconv.Itoa(d)
	}
}

func DayCode(d int) string {
	switch d {
	case ArmstrongDay:
		return "ARM"
	case AldrinDay:
		return "ALD"
	case MoonLandingDay:
		return "MNL"
	default:
		return strconv.Itoa(d)
	}
}

func WeekdayName(w TqWeekday) string {
	return time.Weekday((int(w) + 4) % 7).String()
}

func ShortDate(unixTime int64) string {
	d := Day(unixTime)
	y := Year(unixTime)
	if d < 0 {
		return fmt.Sprintf("%s %d", DayCode(d), y)
	} else {
		ml := MonthLetter(Month(unixTime))
		return fmt.Sprintf("%02d%s %d", d, ml, y)
	}
}

func LongDate(unixTime int64) string {
	w := WeekdayName(Weekday(unixTime))
	d := DayName(Day(unixTime))
	m := MonthName(Month(unixTime))
	y := strconv.Itoa(Year(unixTime))
	var suffix string
	if IsBeforeTranquility(unixTime) {
		suffix = "Before Tranquility"
	} else {
		suffix = "After Tranquility"
	}
	return fmt.Sprintf("%s, %s %s, %s %s", w, d, m, y, suffix)
}
