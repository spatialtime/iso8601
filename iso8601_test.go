package iso8601

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOrdinalDateFormatting(t *testing.T) {
	assert := assert.New(t)

	testString := FormatOrdinalDate(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(testString, "2020-001")

	// can we do leap years?
	testString = FormatOrdinalDate(time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC))
	assert.Equal(testString, "2020-366")
}
func TestOrdinalDateParsing(t *testing.T) {
	assert := assert.New(t)

	testDate, err := ParseOrdinalDate("2020-001")
	assert.NoError(err)
	assert.True(testDate.Equal(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)))

	// can we do leap years?
	testDate, err = ParseOrdinalDate("2020-366")
	assert.NoError(err)
	assert.True(testDate.Equal(time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)))

	// make sure it fails bad regular expression
	_, err = ParseOrdinalDate("I love The Durrells series on Prime")
	assert.Error(err)
}

func TestISOWeekFormatting(t *testing.T) {
	assert := assert.New(t)

	//Jan 1, 2000 is tricky as it is in the 1999 ISO year
	testString := FormatWeek(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), false)
	assert.Equal(testString, "1999-W52-6")
	//same, but request short format
	testString = FormatWeek(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), true)
	assert.Equal(testString, "1999-W52")

	//try another year
	testString = FormatWeek(time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC), false)
	assert.Equal(testString, "2020-W09-7")

}

func TestISOWeekParsing(t *testing.T) {
	assert := assert.New(t)

	// this should put us on Jan 1, 2000
	testDate, err := ParseWeek("1999-W52-6")
	assert.NoError(err)
	assert.True(testDate.Equal(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)))

	//short-form should put us on the monday of the 52nd week of 1999
	testDate, err = ParseWeek("1999-W52")
	assert.NoError(err)
	assert.True(testDate.Equal(time.Date(1999, 12, 27, 0, 0, 0, 0, time.UTC)))

	// make sure it fails bad regular expression
	_, err = ParseWeek("Quarantine-is-a-drag")
	assert.Error(err)
}

func TestISODurationFormatting(t *testing.T) {
	assert := assert.New(t)

	t1 := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)
	assert.Equal(FormatDuration(t2.Sub(t1)), "PT24H0M0S")

	t1 = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	t2 = time.Date(2020, 1, 2, 1, 0, 0, 0, time.UTC)
	assert.Equal(FormatDuration(t2.Sub(t1)), "PT1H0M0S")
}

func TestISODurationParsing(t *testing.T) {
	assert := assert.New(t)

	// 1 day and 1 hour
	dur, err := ParseDuration("P1DT1H")
	assert.NoError(err)
	assert.Equal(25*time.Hour, dur)

	// 1 hour and 2 seconds
	dur, err = ParseDuration("PT1H2S")
	assert.NoError(err)
	assert.Equal(time.Hour+time.Second*2, dur)

	// make sure it fails bad regular expression
	_, err = ParseDuration("I-LOVE-CATS")
	assert.Error(err)
}
