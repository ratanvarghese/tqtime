//Package tqtime interprets UNIX timestamps (seconds since 00:00:00 UTC on January 1970) and outputs dates in the Tranquility calendar. The Tranquility calendar is a perennial calendar developed by Jeff Siggins. A copy of the article proposing this calendar is at www.webcitation.org/6WtW38bAU
package tqtime

//Note to modders: For nonexported symbols in this package, the prefix 'g' is Gregorian, and the prefix 'tq' is Tranquility.

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

const commonYearLen int = 365
const gCommonYearArmstrongDay int = 201
const tqMonthLen int = 28
const gMoonLandingYear int = 1969

//gYearDay gets the Gregorian year and day of year from unixTime.
func gYearDay(unixTime int64) (gy, gyd int) {
	gt := time.Unix(unixTime, 0)
	return gt.Year(), gt.YearDay()
}

//gLeapYear returns true if gy is a Gregorian leap year.
func gLeapYear(gy int) bool {
	return (gy%400 == 0) || (gy%4 == 0 && gy%100 != 0)
}

//gYearLen returns the length of Gregorian year gy, taking into account leap years.
func gYearLen(gy int) int {
	if gLeapYear(gy) {
		return commonYearLen + 1
	}
	return commonYearLen
}

//gArmstrongDay returns the day of the Gregorian year gy that corresponds to Armstrong Day (July 20), taking into account leap years.
func gArmstrongDay(gy int) int {
	if gLeapYear(gy) {
		return gCommonYearArmstrongDay + 1
	}
	return gCommonYearArmstrongDay
}

//clockModulo returns the modulo as a number in range [1,b] rather than a number in range [0,b-1]. If a % b is zero, b is returned. Otherwise a % b is returned. This is important because calendars tend to have cycles but rarely count from 0.
func clockModulo(a, b int) int {
	mod := a % b
	if mod == 0 {
		return b
	}
	return mod
}

//tqYearDay converts a Gregorian year and day of year into a Tranquility day of year.
func tqYearDay(gy, gyd int) int {
	shift := commonYearLen - gCommonYearArmstrongDay
	return clockModulo((gyd + shift), gYearLen(gy))
}

//tqLeapAdjustedYearDay converts a Tranquility day of year and a Gregorian year into a value which is easier to calculate with. If the Gregorian year and Tranquility day of year corresponds to a special day, then that day's constant is returned. Otherwise, the corresponding day of common Tranquility year is returned. For instance if tqyd = 300 and gy = 2000, that represents a day after Aldrin Day on a leap year: the corresponding day of common Tranqility year is 299.
func tqLeapAdjustedYearDay(tqyd, gy int) int {
	if gLeapYear(gy) {
		const tqydAldrin int = tqMonthLen * int(Hippocrates)
		if tqyd == tqydAldrin {
			return AldrinDay
		} else if tqyd > tqydAldrin {
			tqyd--
		}
	}
	if tqyd == commonYearLen {
		if gy == gMoonLandingYear {
			return MoonLandingDay
		}
		return ArmstrongDay
	}
	return tqyd
}

//YearDay returns the day of the Tranquility year of the given unixTime.
func YearDay(unixTime int64) int {
	gy, gyd := gYearDay(unixTime)
	return tqYearDay(gy, gyd)
}

//IsBeforeTranquility returns true if and only if unixTime is before 20:18:01 on Moon Landing Day. This is the exact moment that Neil Armstrong said the word "Tranquility" in the phrase "Houston, Tranquility Base here. The Eagle has landed."
func IsBeforeTranquility(unixTime int64) bool {
	const unixMoonLanding int64 = -14182919
	return unixTime < unixMoonLanding
}

//Year returns the Tranquility year of the given unixTime. This is defined as the years since the first moon landing. Years before Moon Landing Day are represented as negative, Moon Landing Day itself is represented with 0, and years after Moon Landing Day are represented as positive.
func Year(unixTime int64) int {
	gy, gyd := gYearDay(unixTime)
	if gy == gMoonLandingYear && gyd == gCommonYearArmstrongDay {
		return 0
	}
	diff := gy - gMoonLandingYear
	if gyd > gArmstrongDay(gy) {
		diff++
	}
	if diff < 1 { //For 1 AT, depends on previous if statement.
		diff--
	}
	return diff
}

//Month returns the Tranquility month of the given unixTime. If unixTime does not fall on a month, SpecialDay is returned.
func Month(unixTime int64) TqMonth {
	gy, gyd := gYearDay(unixTime)
	tqyd := tqLeapAdjustedYearDay(tqYearDay(gy, gyd), gy)
	if tqyd < 0 {
		return SpecialDay
	}
	return TqMonth(((tqyd - 1) / tqMonthLen) + 1)
}

//Day returns the day of the Tranquility month of the given unixTime. If the unixTime does not fall on a month, a special negative value is returned: one of MoonLandingDay, ArmstrongDay or AldrinDay.
func Day(unixTime int64) int {
	gy, gyd := gYearDay(unixTime)
	tqyd := tqLeapAdjustedYearDay(tqYearDay(gy, gyd), gy)
	if tqyd < 0 {
		return tqyd
	}
	return clockModulo(tqyd, tqMonthLen)
}

//Weekday returns the day of the week of the given unixTime. If unixTime does not fall on a week, the value SpecialWeekday is returned.
func Weekday(unixTime int64) TqWeekday {
	tqd := Day(unixTime)
	if tqd < 0 {
		return SpecialWeekday
	}
	return TqWeekday(clockModulo(tqd, 7))
}

//String returns the English name of the given Tranquility month. If m is not a valid month, a blank string is returned.
func (tqm TqMonth) String() string {
	if tqm < Archimedes || tqm > Mendel {
		return ""
	}
	names := [Mendel]string{
		"Archimedes",
		"Brahe",
		"Copernicus",
		"Darwin",
		"Einstein",
		"Faraday",
		"Galileo",
		"Hippocrates",
		"Imhotep",
		"Jung",
		"Kepler",
		"Lavoisier",
		"Mendel",
	}
	return names[tqm-1]
}

//MonthLetter returns the first letter of the name of the given Tranquility month. If m is not a valid month, a blank string is returned.
func MonthLetter(tqm TqMonth) string {
	name := tqm.String()
	if len(name) > 0 {
		return name[:1]
	}
	return ""
}

//DayName returns the string representation of a day of Tranquility Month, or one of the following special strings when the corresponding special day constant is provided: "Armstrong Day", "Aldrin Day" or "Moon Landing Day".
func DayName(tqmd int) string {
	switch tqmd {
	case ArmstrongDay:
		return "Armstrong Day"
	case AldrinDay:
		return "Aldrin Day"
	case MoonLandingDay:
		return "Moon Landing Day"
	default:
		return strconv.Itoa(clockModulo(tqmd, tqMonthLen))
	}
}

//DayCode returns the string representation of a day of the Tranquility Month, or one of the following special strings when the corresponding special day constant is provided: "ARM" for ArmstrongDay, "ALD" for AldrinDay and "MNL" for MoonLandingDay.
func DayCode(tqmd int) string {
	switch tqmd {
	case ArmstrongDay:
		return "ARM"
	case AldrinDay:
		return "ALD"
	case MoonLandingDay:
		return "MNL"
	default:
		return strconv.Itoa(clockModulo(tqmd, tqMonthLen))
	}
}

//WeekdayName returns the English name of a day of the week. Invalid inputs produce blank strings.
func WeekdayName(tqwd TqWeekday) string {
	if tqwd < Friday || tqwd > Thursday {
		return ""
	}
	return time.Weekday((int(tqwd) + 4) % 7).String()
}

//ShortDate returns the string representation of unixTime in a compact format. On special days, the result is "DDD %y", where DDD is a 3 character day code. On other days, the result is "DDM %y" where DD is the zero-padded day of the month, M is the first letter of the month name. In both cases, %y is a variable-length integer representing the year. %y is preceded by '-' on years Before Tranquility.
func ShortDate(unixTime int64) string {
	tqmd := Day(unixTime)
	tqy := Year(unixTime)
	if tqmd < 0 {
		return fmt.Sprintf("%s %d", DayCode(tqmd), tqy)
	}
	tqml := MonthLetter(Month(unixTime))
	return fmt.Sprintf("%02d%s %d", tqmd, tqml, tqy)
}

//LongDate returns the string representation of the unixTime in a descriptive format.
func LongDate(unixTime int64) string {
	tqmd := Day(unixTime)
	if tqmd == MoonLandingDay {
		return DayName(tqmd)
	}

	tqy := Year(unixTime)
	var suffix string
	if tqy < 0 {
		suffix = "Before Tranquility"
		tqy = -1 * tqy
	} else {
		suffix = "After Tranquility"
	}
	if tqmd < 0 {
		return fmt.Sprintf("%s, %d %s", DayName(tqmd), tqy, suffix)
	}
	tqwd := WeekdayName(Weekday(unixTime))
	tqmn := Month(unixTime)
	return fmt.Sprintf("%s, %s %v, %d %s", tqwd, DayName(tqmd), tqmn, tqy, suffix)
}
