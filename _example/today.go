package main

import (
    "fmt"
    "time"
    "github.com/ratanvarghese/tqtime"
)

func main() {
    t := time.Now().Unix()
    long := tqtime.LongDate(t)
    short := tqtime.ShortDate(t)
    fmt.Printf("%s\t%s\n", long, short)
}
