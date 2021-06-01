package sharedexpenses

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shopspring/decimal"
)

const hundred = 100
const two = 2

type Costs struct {
	Person1       decimal.Decimal
	Person2       decimal.Decimal
	TotalExpenses decimal.Decimal
	TotalIncome   decimal.Decimal
}

var errMissingParameters = fmt.Errorf("missing parameters")

func (c *Costs) Init(person1, person2, totalExpenses *int) error {
	if person1 == nil || person2 == nil || totalExpenses == nil {
		return errMissingParameters
	}
	c.Person1 = decimal.NewFromInt(int64(*person1))
	c.Person2 = decimal.NewFromInt(int64(*person2))
	c.TotalExpenses = decimal.NewFromInt(int64(*totalExpenses))
	c.TotalIncome = c.Person1.Add(c.Person2)

	return nil
}

func (c *Costs) footer() table.Row {
	return table.Row{"total", c.TotalIncome.Round(2).String(), "", c.TotalExpenses.Round(2).String()}
}
