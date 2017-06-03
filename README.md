# TqTime

This package generates dates and date components of the [Tranquility Calendar](http://www.webcitation.org/6WtW38bAU) from standard Gregorian dates.

## What is the Tranquility Calendar?
The Tranquility Calendar is a calendar system developed by Jeff Siggins. It is [perennial](https://en.wikipedia.org/wiki/Perennial_calendar): the days of the week occur on the same dates every month and every year.

### Moon Landing Day and Tranquility
Instead of being centered on a religious event in the ancient
past, the Tranquility Calendar starts on the date of a recent, 
well-recorded event: the day Neil Armstrong and Buzz Aldrin landed
 on the moon. This is 20 July, 1969 in the conventional Gregorian
 calendar system. This day is called Moon Landing Day, and is not 
part of any month or year. At 20:18:01 UTC on Moon Landing Day, 
Neil Armstrong said the word "Tranquility" in the phrase "Houston,
 Tranquility Base here, *The Eagle* has landed." All time before
 this moment is considered "Before Tranquility", and all time
after is considered "After Tranquility". 

### Months and Days of the Week
The months are named after great scientists, instead of [Roman autocrats](https://en.wikipedia.org/wiki/August) or [misnumbered months from previous calendar systems](https://en.wikipedia.org/wiki/September).

1. Archimedes
2. Brahe
3. Copernicus
4. Darwin
5. Einstein
6. Faraday
7. Galileo
8. Hippocrates
9. Imhotep
10. Jung
11. Kepler
12. Lavoisier
13. Mendel

Notice that the months are in alphabetical order: they could be
unambiguously identified with just the first letter. Every month
starts on a Friday, and has 28 days. 1 Archimedes, 1 After
Tranquility is 21 July 1969 in the Gregorian calendar.

### Special Days
In addition to Moon Landing Day, there are two other special days 
which are not bound to any month or week. Armstrong Day is the 
last day of every year (with one possible exception, see below). 
It is right after 28 Mendel. In the Gregorian calendar it is 20 
July.

Aldrin Day is added after 27th Hippocrates and before 28th 
Hippocrates on leap years. Leap years occur every 4 years before 
and after 31 After Tranquility, unless the difference is divisible
 by 100 and not divisible by 400. It corresponds exactly to 29th 
February in the Gregorian calendar.

## Using this package
All commands are assumed to be run from the root source directory.

### Installation
To install this package, install the [standard go tools](https://golang.org/doc/install). Then run: `go get github.com/ratanvarghese/tqtime`.

### Functions
All exported functions are documented in the standard Go format,
so for more information run `go doc`.

One of the less obvious functions is `ShortDate`, which prints the
 Tranquility date in an original format. Since each Tranquility
month has a unique starting letter, any date in a month can be
presented with a unique 3-character code:

    Thursday, 28 Mendel                         28M
    Sunday, 3 Copernicus                        03C

Special Days have unique codes:

    Armstrong Day                               ARM
    Aldrin Day                                  ALD
    Moon Landing Day                            MNL

ShortDate appends this code with a space and a variable-length 
year number. Years Before Tranquility are negative numbers, years 
After Tranquility are positive.

    Thursday, 28 Mendel, 3 After Tranquility    28M 3
    Thursday, 28 Mendel, 3 Before Tranquility   28M -3
    Tuesday, 12 Einstein, 28 After Tranquility  12E 28
    Armstrong Day, 28 After Tranquility         ARM 28

If you do not like the provided `LongDate` and `ShortDate`, you 
can gather all the individual date components and print them as 
you wish.

A basic utility to print the current day exists in `_example`: 
`go run _example/today.go`

### Testing
There is a basic test script called tqcheck which requires [gometalinter](https://github.com/alecthomas/gometalinter) and a UNIX shell. This is convenient if you already have both of those. If not, just use the standard Go tools and whatever else is in your setup:
`go test`

## License
TqTime is provided under the MIT license.
See LICENSE.txt for details.

## 20 July 1968: Which Tranquility Year Does it Belong to?
The [Wikipedia article](https://en.wikipedia.org/wiki/Tranquility_Calendar) about the Tranquility calendar says:
> The year ending the day before Moon Landing Day, and starting on
> the previous Armstrong Day, is 1 Before Tranquility, or 1 BT.

This would suggest that dates Before Tranquility start on
Armstrong Day, instead of ending on it. But Jeff Siggins' article 
states that:
> The last day of each Tranquility year is called Armstrong Day...

Days Before Tranquility are hardly mentioned by Siggins.

When using [tranquilityDate.c](http://www.mithrandir.com/Tranquility/tranquilityDate.c) by Scott M Harrison, 20 July 1968 is considered Armstrong Day 1 Before Tranquility... but 20 July 1967 is considered Armstrong Day 3 Before Tranquility!

As far as I know, the biggest users of the Tranquility calendar today are the [Orion's Arm collaborative science fiction project](http://www.orionsarm.com) but they do not care about anything that happened on 20 July 1968.

On the matter of Armstrong Days Before Tranquility, this package 
assumes the following:
* 20 July 1969 is Moon Landing Day, and not part of any year.
* 1 Before Tranquility has no Armstrong Day. The year starts at 1 
  Archimedes Before Tranquility (21 July 1968), and ends on 28 
  Mendel Before Tranquility (19 July 1969). The day that *would* 
  be Armstrong Day 1 Before Tranquility is Moon Landing Day, but 
  that is not part of any year.
* Every other year, stretching all the way from the Big Bang to 
  the Heat Death of the Universe, ends on Armstrong Day. Armstrong
  Day always corresponds to the Gregorian date 20 July.

## Authors
Ratan Varghese
