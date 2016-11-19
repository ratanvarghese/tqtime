//Package tqtime interprets UNIX timestamps (seconds since 00:00:00 UTC on January 1970) and outputs dates in the Tranquility calendar. The Tranquility calendar is a perennial calendar developed by Jeff Siggins. A copy of the article proposing this calendar is at www.webcitation.org/6WtW38bAU
package tqtime

import (
	"fmt"
	"strconv"
	"time"
)

//TqWeekday represents a day of the Tranquility week.
type TqWeekday int

//In the Tranquility calendar, all months start on a Friday. Armstrong Day, Aldrin Day and Moon Landing Day do not have associated days of the week. These days are represented as SpecialWeekday when their positions in the week are requested.
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

//TqMonth represents a Tranquility month.
type TqMonth int

//The months of the Tranquility calendar are named after scientists and are in alphabetical order. Armstrong Day, Aldrin Day and Moon Landing Day do not have associated months. These days are represented as SpecialDay when their month is requested.
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

//ArmstrongDay is the last day of each Tranquility year, except for 1 Before Tranquility (BT). It is 20 July in the Gregorian calendar. 20 July 1969 is not part of a year, and thus not an Armstrong Day. 20 July 1968 is considered Armstrong Day 2 BT by this package, but is considered Armstrong Day 1 BT by tranquilityDate.c (by Scott M Harrison).
const ArmstrongDay int = -1

//AldrinDay an extra day added during leap years. It is inserted before the last day of Hippocrates, interrupting the month and week. It is 29 February in the Gregorian calendar.
const AldrinDay int = -2

//MoonLandingDay is 20 July 1969, the day Neil Armstrong and Edwin "Buzz" Aldrin landed on the moon. It is not part of any week, month or year, although for convenience it is treated as year 0 by this package.
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
	y, m, d := g.Date()
	return y == mlYear && m == mlMonth && d == mlDay
}

func isAfterArmstrongDay(g time.Time) bool {
	_, m, d := g.Date() //Avoid YearDay due to leap year
	return (m > mlMonth) || (m == mlMonth && d > mlDay)
}

//IsBeforeTranquility returns true if and only if unixTime is before 20:18:01.2 on Moon Landing Day. This is the exact moment that Neil Armstrong said the word "Tranquility" in the phrase "Houston, Tranquility Base here. The Eagle has landed."
func IsBeforeTranquility(unixTime int64) bool {
	const unixMoonLanding int64 = -14182919
	return unixTime < unixMoonLanding
}

//Year returns the Tranquility year of the given unixTime. This is defined as the years since the first moon landing. Years before Moon Landing Day are represented as negative, Moon Landing Day itself is represented with 0, and years after Moon Landing Day are represented as positive.
func Year(unixTime int64) int {
	g := time.Unix(unixTime, 0).UTC()
	if isMoonLandingDay(g) {
		return 0
	}
	yearDiff := g.Year() - mlYear
	if isAfterArmstrongDay(g) {
		yearDiff++
	}
	if yearDiff < 1 {
		yearDiff--
	}
	return yearDiff
}

//Month returns the Tranquility month of the given unixTime. If unixTime does not fall on a month, SpecialDay is returned.
func Month(unixTime int64) TqMonth {
	g := time.Unix(unixTime, 0).UTC()
	if isMoonLandingDay(g) {
		return SpecialDay //Moon Landing Day
	}
	yd := YearDay(unixTime)
	if isGregorianLeapYear(g) {
		const leapDay int = tqMonthLen * int(Hippocrates)
		if yd == leapDay {
			return SpecialDay //Aldrin Day
		} else if yd > leapDay {
			yd--
		}
	}
	if yd == commonYearLen {
		return SpecialDay //Armstrong Day
	}
	return TqMonth(((yd - 1) / tqMonthLen) + 1)
}

//Day returns the day of the Tranquility month of the given unixTime. If the unixTime does not fall on a month, a special negative value is returned: one of MoonLandingDay, ArmstrongDay or AldrinDay.
func Day(unixTime int64) int {
	g := time.Unix(unixTime, 0).UTC()
	if isMoonLandingDay(g) {
		return MoonLandingDay
	}
	yd := YearDay(unixTime)
	if isGregorianLeapYear(g) {
		const leapDay int = tqMonthLen * int(Hippocrates)
		if yd == leapDay {
			return AldrinDay
		} else if yd > leapDay {
			yd--
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

//YearDay returns the day of the Tranquility year of the given unixTime.
func YearDay(unixTime int64) int {
	g := time.Unix(unixTime, 0).UTC()
	yearLen := commonYearLen
	armstrongYearDay := mlYearDay
	if isGregorianLeapYear(g) {
		yearLen++
		armstrongYearDay++
	}
	diff := g.YearDay() - armstrongYearDay
	if diff > 0 || yearLen == diff {
		return diff
	}
	return yearLen + diff
}

//Weekday returns the day of the week of the given unixTime. If unixTime does not fall on a week, the value SpecialWeekday is returned.
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

//MonthName returns the English name of the given Tranquility month. If m is not a valid month, a blank string is returned.
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

//MonthLetter returns the first letter of the name of the given Tranquility month. If m is not a valid month, a blank string is returned.
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

//DayName returns the string representation of a day of Tranquility Month, or one of the following special strings when the corresponding special day constant is provided: "Armstrong Day", "Aldrin Day" or "Moon Landing Day".
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

//DayCode returns the string representation of a day of the Tranquility Month, or one of the following special strings when the corresponding special day constant is provided: "ARM" for ArmstrongDay, "ALD" for AldrinDay and "MNL" for MoonLandingDay.
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

//WeekdayName returns the English name of a day of the week. Invalid inputs produce blank strings.
func WeekdayName(w TqWeekday) string {
	if w < Friday || w > Thursday {
		return ""
	}
	return time.Weekday((int(w) + 4) % 7).String()
}

//ShortDate returns the string representation of unixTime in a compact format. On special days, the result is "DDD %y", where DDD is a 3 character day code. On other days, the result is "DDM %y" where DD is the zero-padded day of the month, M is the first letter of the month name. In both cases, %y is a variable-length integer representing the year. %y is preceded by '-' on years Before Tranquility.
func ShortDate(unixTime int64) string {
	d := Day(unixTime)
	y := Year(unixTime)
	if d < 0 {
		return fmt.Sprintf("%s %d", DayCode(d), y)
	}
	ml := MonthLetter(Month(unixTime))
	return fmt.Sprintf("%02d%s %d", d, ml, y)
}

//LongDate returns the string representation of the unixTime in a descriptive format.
func LongDate(unixTime int64) string {
	d := Day(unixTime)
	if d == MoonLandingDay {
		return "Moon Landing Day"
	}

	y := Year(unixTime)
	var suffix string
	if y < 0 {
		suffix = "Before Tranquility"
		y = -1 * y
	} else {
		suffix = "After Tranquility"
	}
	if d < 0 {
		return fmt.Sprintf("%s, %d %s", DayName(d), y, suffix)
	}
	w := WeekdayName(Weekday(unixTime))
	m := MonthName(Month(unixTime))
	return fmt.Sprintf("%s, %s %s, %d %s", w, DayName(d), m, y, suffix)
}
