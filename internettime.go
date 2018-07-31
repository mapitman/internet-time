// Package internettime is a simple library to convert time of day to Swatch Internet Time in beats. (See https://en.wikipedia.org/wiki/Swatch_Internet_Time)
package internettime

import (
	"time"
)

// ConvertToBeats returns the argument time converted to beats.
func ConvertToBeats(t time.Time) int {
	var utc = t.UTC()
	var seconds = utc.Second()
	var minutes = utc.Minute()
	var hours = utc.Hour()
	if hours < 24 {
		hours = hours + 1
	}

	var beatsFloat = float32(seconds+(minutes*60)+(hours*3600)) / float32(86.4)
	var beats = int(beatsFloat)
	if beats == 1000 {
		beats = 0
	}

	return beats
}
