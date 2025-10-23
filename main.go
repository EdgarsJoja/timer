package main

import (
	"flag"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	timeSeconds := flag.Int("s", 0, "Timer seconds duration")
	timeMinutes := flag.Int("m", 0, "Timer minutes duration")
	timeHours := flag.Int("h", 0, "Timer hours duration")

	timeString := flag.String("t", "", "Specific time. e.g.: 15:04:05")

	title := flag.String("title", "Time's up!", "Notification title")
	message := flag.String("msg", "", "Notification message")

	flag.Parse()

	var timer *time.Timer

	if *timeString != "" {
		targetTime, err := time.Parse(time.TimeOnly, *timeString)

		if err != nil {
			panic(err)
		}

		now := time.Now()
		targetTime = time.Date(
			now.Year(), now.Month(), now.Day(),
			targetTime.Hour(), targetTime.Minute(), targetTime.Second(),
			0, now.Location(),
		)

		timer = time.NewTimer(time.Until(targetTime))
	} else if *timeSeconds > 0 || *timeMinutes > 0 || *timeHours > 0 {
		timer = time.NewTimer(time.Duration(*timeSeconds)*time.Second + time.Duration(*timeMinutes)*time.Minute + time.Duration(*timeHours)*time.Hour)
	}

	if timer == nil {
		os.Exit(1)
	}

	<-timer.C

	beeep.AppName = "Timer"
	err := beeep.Notify(*title, *message, "stopwatch")
	if err != nil {
		panic(err)
	}
}
