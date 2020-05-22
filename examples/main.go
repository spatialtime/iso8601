package main

import (
	"fmt"
	"github.com/spatialtime/iso8601"
	"time"
)

// ISOWeekExample demonstrates formatting a time.Time to ISO,
// and then parsing that string back into a time.Time.
func ISOWeekExample() {
	fmt.Println("\nISO Week example:")

	formatted := iso8601.FormatWeek(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), true)
	fmt.Println(formatted)

	parsed, err := iso8601.ParseWeek(formatted)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parsed)
}

// ISODurationExample demonstrates parsing an ISO duration string to a time.Duration,
// and then formatting that time.Duration back out to an ISO duration string.
func ISODurationExample() {
	fmt.Println("\nISO Duration example:")

	parsed, err := iso8601.ParseDuration("P45DT3H3.266662S")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(parsed)

	formatted := iso8601.FormatDuration(parsed)
	fmt.Println(formatted)
}

// ISOOrdinalDateExample demonstrates formatting a time.Time to an ISO ordinal date,
// and then parsing that string back into a time.Time.
func ISOOrdinalDateExample() {
	fmt.Println("\nISO Ordinal date example:")

	s := iso8601.FormatOrdinalDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	fmt.Println("formatted as:  ", s)

	date, err := iso8601.ParseOrdinalDate(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(date)
}

// ISOFormatDateTimeExample demonstrates formatting a time.Time instance
// via a variety of layouts.
func ISOFormatDateTimeExample() {
	fmt.Println("\nFormat DateTime example:")
	loc, _ := time.LoadLocation("America/Los_Angeles")

	t := time.Date(2020, 1, 1, 15, 30, 10, 693000000, loc)
	fmt.Println("YYYY", iso8601.FormatDateTime(t, iso8601.ISOYear))
	fmt.Println("YYYY-MM", iso8601.FormatDateTime(t, iso8601.ISOYearMonth))
	fmt.Println("YYYY-MM-DD", iso8601.FormatDateTime(t, iso8601.ISOFullDate))
	fmt.Println("hh:mm", iso8601.FormatDateTime(t, iso8601.ISOHoursMinutes))
	fmt.Println("hh:mm:ss", iso8601.FormatDateTime(t, iso8601.ISOHoursMinutesSeconds))
	fmt.Println("hh:mm:ss.sss", iso8601.FormatDateTime(t, iso8601.ISOFullTime))

	fmt.Println("hh:mm:ss.sssZ", iso8601.FormatDateTime(t,
		iso8601.ISOFullDate+"T"+iso8601.ISOFullTime+iso8601.ISOTzZulu))

	fmt.Println("YYYY-MM-DDThh:mm:ss-hh:mm", iso8601.FormatDateTime(t, time.RFC3339))
}

func main() {
	ISOWeekExample()
	ISODurationExample()
	ISOOrdinalDateExample()
	ISOFormatDateTimeExample()
}
