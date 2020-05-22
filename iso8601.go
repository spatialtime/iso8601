// Package iso8601 provides support for parsing and formatting ISO 8601 constructs.
// Formatting methods accept a time.Time instance and return the appropriate ISO 8601 string.
// Parsing methods accept an ISO 8601 string and return a time.Time instance.
package iso8601

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// These are layout constants that supplement the layout constants provided by time.Time.
// Note: golang's time.RFC3339 constant represents a layout string containing
// ISO Date, full ISO time (minus milliseconds) and time zone.
const (
	ISOYear                 = "2006"
	ISOYearMonth            = "2006-01"
	ISOFullDate             = "2006-01-02"
	ISOHoursMinutes         = "15:04"
	ISOHoursMinutesSeconds  = "15:04:05"
	ISOFullTime             = "15:04:05.000"
	ISOTzZulu               = "Z"
	ISOTzOffsetHours        = "-07"
	ISOTzOffsetHoursMinutes = "-07:00"
	MinWeek                 = 1
	MinYear                 = 1
	MaxYear                 = 9999
)

// ErrYearRange is returned when a week is not within our permitted range.
var ErrYearRange = fmt.Errorf("year is out of range (valid range: %d–%d inclusive)", MinYear, MaxYear)

// ErrWeekRange is returned when a week is not within our permitted range.
var ErrWeekRange = fmt.Errorf("week is out of range (valid range: %d–number of iso weeks in the given year inclusive)", MinWeek)

// Weekday returns day of week with Monday=0...Sunday=6.
// Utilizes Zeller's Congruence.  see: https://en.wikipedia.org/wiki/Zeller%27s_congruence
func Weekday(year, month, day int) int {
	if month == 1 {
		month = 13
		year--
	} else if month == 2 {
		month = 14
		year--
	}

	dow := (day + (13 * (month + 1) / 5) + year +
		(year / 4) - (year / 100) +
		(year / 400)) % 7

	return (7 + (dow - 2)) % 7
}

// FormatOrdinalDate returns an ISO 8601 ordinal date string.
// In this context, Ordinal date represents the nth day of the year.
func FormatOrdinalDate(date time.Time) string {
	return date.Format("2006-002")
}

// ParseOrdinalDate parses an ISO 8601 string representing a ordinal date,
// and returns the resultant golang time.Time insance.
func ParseOrdinalDate(isoOrdinalDate string) (time.Time, error) {
	return time.Parse("2006-002", isoOrdinalDate)
}

// ParseDuration parses an ISO 8601 string representing a duration,
// and returns the resultant golang time.Duration instance.
func ParseDuration(isoDuration string) (time.Duration, error) {
	re := regexp.MustCompile(`^P(?:(\d+)Y)?(?:(\d+)M)?(?:(\d+)D)?T(?:(\d+)H)?(?:(\d+)M)?(?:(\d+(?:.\d+)?)S)?$`)
	matches := re.FindStringSubmatch(isoDuration)
	if matches == nil {
		return 0, errors.New("duration string is of incorrect format")
	}

	seconds := 0.0

	//skipping years and months

	//days
	if matches[3] != "" {
		f, err := strconv.ParseFloat(matches[3], 32)
		if err != nil {
			return 0, err
		}

		seconds += (f * 24 * 60 * 60)
	}
	//hours
	if matches[4] != "" {
		f, err := strconv.ParseFloat(matches[4], 32)
		if err != nil {
			return 0, err
		}

		seconds += (f * 60 * 60)
	}
	//minutes
	if matches[5] != "" {
		f, err := strconv.ParseFloat(matches[5], 32)
		if err != nil {
			return 0, err
		}

		seconds += (f * 60)
	}
	//seconds & milliseconds
	if matches[6] != "" {
		f, err := strconv.ParseFloat(matches[6], 32)
		if err != nil {
			return 0, err
		}

		seconds += f
	}

	goDuration := strconv.FormatFloat(seconds, 'f', -1, 32) + "s"
	return time.ParseDuration(goDuration)

}

// FormatDuration returns an ISO 8601 duration string.
func FormatDuration(dur time.Duration) string {
	return "PT" + strings.ToUpper(dur.Truncate(time.Millisecond).String())
}

// FormatWeek returns an ISO 8601 week string.
func FormatWeek(date time.Time, shortForm bool) string {
	year, week := date.ISOWeek()
	if shortForm {
		return fmt.Sprintf("%d-W%02d", year, week)
	}

	//date.Weekday() returns a Sunday-started week, with Sunday=0.
	// have to adjust to ISO Monday-started week, with Monday=1.
	dow := ((7 + date.Weekday() - 1) % 7) + 1

	return fmt.Sprintf("%d-W%02d-%d", year, week, dow)
}

func calcP(y int) int {
	return y + (y / 4) - (y / 100) + (y / 400)
}

// ISOYearWeeks returns the number of ISO weeks contained
// in a given Gregorian calendar year.
func ISOYearWeeks(gregYear int) int {
	if (calcP(gregYear)%7 == 4) ||
		(calcP(gregYear-1)%7 == 3) {
		return 53
	}
	return 52
}

// ParseWeek parses an ISO 8601 string representing an ISO week,
// and returns the resultant golang time.Time instance.
// Note: if the ISO week is of the short form (doesn't include day of week),
// this function will return a time.Time instance with day of week of Monday.
func ParseWeek(isoWeek string) (time.Time, error) {
	re := regexp.MustCompile(`^(\d{4})-W([0-5]\d)(?:-([1-7]))?$`)
	matches := re.FindStringSubmatch(isoWeek)
	if matches == nil {
		return time.Time{}, errors.New("isoWeek string is of incorrect format")
	}

	year, err := strconv.Atoi(matches[1])
	if err != nil {
		return time.Time{}, err
	}
	if year < MinYear || year > MaxYear {
		return time.Time{}, ErrYearRange
	}

	week, err := strconv.Atoi(matches[2])
	if week < MinWeek || week > ISOYearWeeks(year) {
		return time.Time{}, ErrWeekRange
	}
	if err != nil {
		return time.Time{}, err
	}

	week--
	daysToAdd := week * 7

	if matches[3] != "" {
		day, err := strconv.Atoi(matches[3])
		if err != nil {
			return time.Time{}, err
		}

		daysToAdd += day - 1
	}

	daysToAdd -= Weekday(year, 1, 4)

	return time.Date(year, time.January, 4+daysToAdd, 0, 0, 0, 0, time.UTC), nil
}

// ParseDateTime parses an ISO 8601 string representing a date or time or date+time,
// and returns the resultant golang time.Time insance.
func ParseDateTime(isoTime, layout string) (time.Time, error) {
	return time.Parse(layout, isoTime)
}

// FormatDateTime returns an ISO 8601 date.
func FormatDateTime(t time.Time, layout string) string {
	return t.Format(layout)
}
