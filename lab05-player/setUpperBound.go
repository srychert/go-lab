package main

import (
	"regexp"
	"strconv"
)

func setUpperBound(boundsMsg string, upperBound *int) {
	// regex for finding bounds
	rB, _ := regexp.Compile("[0-9]+")
	bounds := rB.FindAllString(boundsMsg, -1)
	var ints []int

	for _, v := range bounds {
		vi, err := strconv.Atoi(v)
		if err == nil {
			ints = append(ints, vi)
		}
	}

	if len(ints) > 0 {
		*upperBound = ints[0]
	}

	for _, e := range ints {
		if e > *upperBound {
			*upperBound = e
		}
	}
}
