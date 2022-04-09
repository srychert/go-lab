package guess

import "time"

func todayDate() Date {
	year, month, day := time.Now().Date()
	d := Date{Day: day, Month: int(month), Year: year}

	return d
}
