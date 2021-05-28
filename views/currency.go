package views

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Currency struct {
	ID    int     `json:"vid"`
	Code  string  `json:"vcode"`
	Name  string  `json:"vname"`
	Price float64 `json:"vprice"`
}

func (c *Currency) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Code, validation.Required, validation.Length(1, 3)),
		validation.Field(&c.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&c.Price, validation.Required, validation.Min(float64(0)), validation.Max(float64(99999999.99))),
	)
}
