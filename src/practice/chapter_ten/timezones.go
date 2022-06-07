package main

import "fmt"

type TZ int

const (
	HOUR TZ = 60 * 60
	UTC  TZ = 0 * HOUR
	EST  TZ = -5 * HOUR
	CST  TZ = -6 * HOUR
)

var timeZone = map[TZ]string{UTC: "Universal Greenwich Time", EST: "Eastern Stand Time", CST: "Central Stand Time"}

func main() {
	fmt.Println(EST)
	fmt.Println(0 * HOUR)
	fmt.Println(-6 * HOUR)
}

func (tz TZ) String() string {
	if zone, ok := timeZone[tz]; ok {
		return zone
	}
	return ""
}
