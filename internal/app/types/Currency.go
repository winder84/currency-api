package types

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

type Currency struct {
	From      string    `db:"currency_from,pk"`
	To        string    `db:"currency_to,pk"`
	Well      float64   `db:"well"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (c *Currency) validateCurrency() error {
	var err error
	if len(c.From) != 3 {
		err = fmt.Errorf("%s - len error", c.From)
	}
	if len(c.To) != 3 {
		err = fmt.Errorf("%s - len error", c.To)
	}
	for _, curr := range []string{
		c.From,
		c.To,
	} {
		for _, s := range curr {
			if s > unicode.MaxASCII {
				err = fmt.Errorf("%s - unicode error", curr)
			}
		}
		curr = strings.ToUpper(curr)
	}
	return err
}
