package main

import (
    "fmt"
    "time"
    "github.com/ratanvarghese/tqtime"
)

func main() {
    fmt.Println(tqtime.LongDate(time.Now().Unix()))
}
