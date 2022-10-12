package keeper

import (
	"strconv"
	t "time"
)

// Global time layout
const timeLayout = t.UnixDate

// Get seconds in int64 to time.Duration
func GetDuration(sec int64) t.Duration {
	duration, _ := t.ParseDuration(ToSecondsStr(sec))
	return duration
}

// Convert seconds in int64 to string and append 's'
// End result: "{seconds}s"
func ToSecondsStr(sec int64) string {
	return strconv.FormatInt(sec, 10) + "s"
}

// Global conversion of time to string.
func TimeToString(time t.Time) string {
	return time.Format(timeLayout)
}

// Global conversion of string to time
func StringToTime(time string) (t.Time, error) {
	return t.Parse(timeLayout, time)
}
