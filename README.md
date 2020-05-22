# iso8601

[![Go Report Card](https://goreportcard.com/badge/github.com/spatialtime/iso8601?style=flat-square)](https://goreportcard.com/report/github.com/spatialtime/iso8601)
![Top Language](https://img.shields.io/github/languages/top/spatialtime/iso8601?style=flat-square)
[![License][license-image]][license-url]

Format and parse ISO 8601 dates, times, datetimes, time zones, weeks, ordinal dates and durations.   Use this right out of the box, or as a starting point to build upon for more specialized needs.

## Contents

- [iso8601](#iso8601)
  - [Contents](#contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Notes](#notes)
  - [Author](#author)


## Installation
```
$ go get github.com/spatialtime/iso8601
```

## Usage
```
import(
   "fmt"
    "github.com/spatialtime/iso8601"
    "time"
)

func main(){
    // ISO weeks
    formatted := iso8601.FormatWeek(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), false)
    fmt.Println(formatted) // 2020-W53-5
    date, err := iso8601.ParseWeek(formatted)
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Println(date) // 2021-01-01 00:00:00 +0000 UTC

    // Ordinal dates 
    formatted = iso8601.FormatOrdinalDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
    fmt.Println(formatted) // 2020-001
    date, err = iso8601.ParseOrdinalDate(formatted)
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Println(date) // 2020-01-01 00:00:00 +0000 UTC

    // Durations
    t1 := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
    t2 := time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)
    formatted = iso8601.FormatDuration(t2.Sub(t1))
    fmt.Println(formatted) // PT24H0M0S
    duration, err := iso8601.ParseDuration(formatted)
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Println(duration) // 24h0m0s
}
```
## Notes
I included nine ISO-specific layout strings to expedite parsing and formatting of dates, times and datetimes.  Use these in calls to time.Parse() and time.Format().
```
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
)
```
Example usage:
```
t := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
fmt.Println(t.Format(iso8601.ISOFullDate)) // 2020-01-01
```

## Author

* Matt Savage matt@spatialtime.com 

Copyright © 2020 Matt Savage | MIT license

[license-image]: https://img.shields.io/:license-mit-blue.svg?style=flat-square
[license-url]: LICENSE
