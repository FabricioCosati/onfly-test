package validator

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Datetime struct {
	time.Time
}

func (d *Datetime) UnmarshalJSON(input []byte) error {
	date := strings.Trim(string(input), `"`)
	newTime, err := time.Parse("2006-01-02", date)

	if err != nil {
		return err
	}

	d.Time = newTime
	return nil
}

func GteDate(field validator.FieldLevel) bool {
	initDate := field.Parent().FieldByName(field.Param()).Interface().(Datetime).Time
	finalDate := field.Field().Interface().(Datetime).Time

	return finalDate.After(initDate) || finalDate.Equal(initDate)
}

func GtToday(field validator.FieldLevel) bool {
	date := field.Field().Interface().(Datetime).Time
	now := time.Now().Truncate(24 * time.Hour)

	return date.After(now) || date.Equal(now)
}

func DateRequired(field validator.FieldLevel) bool {
	date, ok := field.Field().Interface().(Datetime)
	if !ok {
		return false
	}

	return !date.IsZero()
}
