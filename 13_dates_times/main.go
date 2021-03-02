package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	t1 := time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(t)                 // current local time
	fmt.Println(time.Now().Date()) // date
	fmt.Println(t.Clock())         // time
	fmt.Println(t.Hour())          // hour
	fmt.Println(t.Before(t1))      // False
	fmt.Println(t.After(t1))       // True

	// Epoch
	secs := t.Unix()
	nanos := t.UnixNano()
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

	// Time Formatting / Parsing
	// t := time.Now()
	// fmt.Println(t.Format(time.ANSIC))
	// fmt.Println(t.Format(time.Kitchen))
	// fmt.Println(t.Format(time.RFC1123))
	// fmt.Println(t.Format("Mon-Jan")) // custom format

	// t1, _ := time.Parse(time.Kitchen, "9:11AM")
	// fmt.Println(t1)
	// fmt.Println(t1.Hour())
}
