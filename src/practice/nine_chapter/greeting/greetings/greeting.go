package greetings

import "time"

func IsAM() bool {
	now := time.Now()
	return now.Hour() <= 12
}

func IsAfternoon() bool {
	now := time.Now()
	return now.Hour() <= 18 && now.Hour() > 12
}

func IsEvening() bool {
	now := time.Now()
	return now.Hour() <= 22 && now.Hour() > 18
}

func GoodDay(name string) string {
	return "Good Day " + name
}

func GoodNight(name string) string {
	return "Good Night " + name
}
