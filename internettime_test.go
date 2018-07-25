package internettime

import (
	"testing"
	"time"
)

func TestConvertToBeats(t *testing.T) {
	loc, _ := time.LoadLocation("US/Pacific")
	input, _ := time.ParseInLocation(time.RFC3339, "2018-07-24T22:12:00-07:00", loc)
	expected := 258
	actual := ConvertToBeats(input)
	if actual != expected {
		t.Errorf("ConvertToBeats() = %d, expected %d", actual, expected)
	}
}
