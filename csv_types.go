package nostradamus

import (
	"math"
	"strconv"
	"time"

	"github.com/aodin/date"
)

type Date struct {
	date.Date
}

func NewDate(year int, month time.Month, day int) Date {
	d := date.New(year, month, day)
	return Date{Date: d}
}

func (d *Date) UnmarshalCSV(text string) (err error) {
	if text == "" {
		return nil
	}

	if text == "0000-00-00" {
		return nil
	}

	layout := "2006-01-02"
	time, err := time.Parse(layout, text)
	d.Date = date.FromTime(time)
	return err
}

type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalCSV(text string) (err error) {
	if text == "" {
		return nil
	}

	if text == "0000-00-00 00:00:00" {
		return nil
	}

	layout := "2006-01-02 15:04:05"
	time, err := time.Parse(layout, text)
	dt.Time = time
	return err
}

type CrappyInt int

func (i *CrappyInt) UnmarshalCSV(text string) (err error) {
	if text == "" {
		return nil
	}

	f, err := strconv.ParseFloat(text, 32)
	if err != nil {
		return err
	}

	f = math.Round(f)
	ci := CrappyInt(int(f))
	*i = ci
	return nil
}
