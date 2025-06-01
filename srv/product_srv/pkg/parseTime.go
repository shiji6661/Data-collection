package pkg

import "time"

const TIME = "2006-01-02"

func ParseTime(str string) (time.Time, error) {
	parse, err := time.Parse(TIME, str)
	if err != nil {
		return time.Time{}, err
	}
	return parse, err
}
