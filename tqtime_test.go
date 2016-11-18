package tqtime

import (
	"fmt"
	"testing"
	"time"
)

func TestNowShort(t *testing.T) {
	fmt.Println(ShortDate(time.Now().Unix()))
}

func TestNowLong(t *testing.T) {
	fmt.Println(LongDate(time.Now().Unix()))
}

func TestYearDay(t *testing.T) {
	fmt.Println(YearDay(time.Now().Unix()))
}
