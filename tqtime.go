//g_* == Gregorian
package tqtime

import "time"

type Weekday int

const (
	SpecialWeekday Weekday = iota
	Friday
	Saturday
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
)

type Month int

const (
	SpecialDay Month = iota
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

const ArmstrongDay = -1
const AldrinDay = -2
const MoonLandingDay = -3

func moonLandingMoment() time.Time {
    return time.Date(1969, time.July, 20, 20, 18, 1, 200000000, time.UTC)
}

func gregorianLeapYear(g time.Time) bool {
	y := g.Year()
	feb29 := time.Date(y, time.February, 29, 1, 0, 0, 0, g.Location())
	return (feb29.Month() == time.February)
}

func isMoonLandingDay(g time.Time) bool {
    mlm := moonLandingMoment()
    g_y, g_m, g_d := g.Date()
    mlm_y, mlm_m, mlm_d := mlm.Date()
    return g_y == mlm.y && g_m == mlm.m && g_d == mlm_d
}

func isAfterArmstrongDay(g time.Time) bool {
    mlm := moonLandingMoment()
    g_y, g_m, g_d := g.Date()
    mlm_y, mlm_m, mlm_d := mlm.Date()
    return (g_m > mlm_m) || (g_m == mlm_m && g_d > mlm_d)
}

func gregYearDayToTqYearDay(g_yd, yearLen int) int {
    mlm := moonLandingMoment()
    baseShift := mlm.YearDay()
    diff := g_yd - baseShift
    if diff > 0 || yearLen == diff {
        return diff
    } else {
        return yearLen - diff
    }
}

func IsBeforeTranquility(unixTime int64) bool {
    t := time.Unix(unixTime, 0)
    return t.Before(moonLandingMoment())
}

func Year(unixTime int64) int {
    g := time.Unix(unixTime, 0)
    if isMoonLandingDay(g) {
        return 0
    }
    mlm = moonLandingMoment()
    yearDiff := g_y - mlm.Year()
    if isAfterArmstrongDay(g) {
        yearDiff += 1
    }
    if g.Before(mlm) {
        yearDiff -= 1
    }
    return yearDiff
}

func Month(unixTime int64) Month {
    yd := YearDay(unixTime)
    g := time.Unix(unixTime, 0)
    if gregorianLeapYear(g) {
        const leapDay = Hippocrates*28
        if yd == leapDay {
            return SpecialDay //Aldrin Day
        } else if yd > leapDay {
            yd -= 1
        }
    }
    if yd == 365 {
        return SpecialDay //Armstrong Day
    }
    return ((yd - 1)/ 28) + 1
}

func Day(unixTime int64) int {
    g := time.Unix(unixTime)
    if isMoonLandingDay(g) {
        return MoonLandingDay
    }
    yd := YearDay(unixTime)
    if gregorianLeapYear(g) {
        const leapDay = Hippocrates*28
        if yd == leapDay {
            return SpecialDay //Aldrin Day
        } else if yd > leapDay {
            yd -= 1
        }
    }
    if yd == 365 {
        return ArmstrongDay
    } else if yd % 28 == 0 {
        return 28
    } else {
        return (yd % 28)
    }
}

func YearDay(unixTime int64) int {
    g := time.Unix(unixTime, 0)
    g_yd := g.YearDay()
    yearLen := 365
    if gregorianLeapYear(g) {
        yearLen += 1
    }
    return gregYearDayToTqYearDay(g_yd, yearLen)
}

func Weekday(unixTime int64) Weekday {
    d := Day(unixTime)
    if d < 0 {
        return SpecialWeekday
    } else if d % 7 == 0 {
        return Thursday
    } else {
        return d % 7
    }
}

func LongDate(unixTime int64) string {
}

func ShortDate(unixTime int64) string {
}
