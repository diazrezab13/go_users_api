package dateutils

import "time"

const (
	apiDateLayout = "2011-11-02"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
