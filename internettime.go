// Package internettime is a simple library to convert time of day to Swatch Internet Time in beats. (See https://en.wikipedia.org/wiki/Swatch_Internet_Time)
package internettime

import (
	"time"
)

// ConvertToBeats returns the argument time converted to beats.
func ConvertToBeats(t time.Time) int {
	t = t.UTC()
	var beats = int(float64(t.Second()+(t.Minute()*60)+((t.Hour()+1)*3600)) / 86.4)
	return beats
}
