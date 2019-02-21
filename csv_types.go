package nostradamus

import (
	"encoding/json"
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

func (d *Date) UnmarshalCSV(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	layout := "2006-01-02"
	time, err := time.Parse(layout, value)
	d.Date = date.FromTime(time)
	return err
}
