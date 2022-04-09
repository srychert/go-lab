package guess

import (
	"fmt"
	"strconv"
	"strings"
)

var filePath string = "website/public/hallOfFame.json"

type Input struct {
	guess   int
	command string
}

type Result struct {
	Name  string
	Score int
}

type Date struct {
	Day   int
	Month int
	Year  int
}

type WrongDateError struct{}

func (m *WrongDateError) Error() string {
	return "Date was in wrong format! It should be `dd-mm-yyyy`"
}

func (s *Date) toString() string {
	str := fmt.Sprintf("%02d-%02d-%d", s.Day, s.Month, s.Year)
	return str
}

func (s Date) MarshalText() (text []byte, err error) {
	return []byte(s.toString()), nil
}

func (s *Date) UnmarshalText(text []byte) error {
	str := string(text)
	arr := strings.Split(str, "-")

	if len(arr) != 3 {
		return &WrongDateError{}
	}

	var ints []int
	for _, v := range arr {
		x, err := strconv.Atoi(v)
		if err != nil {
			return &WrongDateError{}
		}
		ints = append(ints, x)
	}
	d := Date{ints[0], ints[1], ints[2]}
	*s = d
	return nil
}
