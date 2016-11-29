package main

import (
	"fmt"
	"github.com/ratanvarghese/tqtime"
	"time"
)

func main() {
	t := time.Now()
	long := tqtime.LongDate(t.Year(), t.YearDay())
	short := tqtime.ShortDate(t.Year(), t.YearDay())
	fmt.Printf("%s\t%s\n", long, short)
}
