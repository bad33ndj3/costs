package sharedexpenses

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shopspring/decimal"
)

type ToRatio struct {
	Costs
	percentageOfIncome1 decimal.Decimal
	percentageOfIncome2 decimal.Decimal
	expensesToRatio1    decimal.Decimal
	expensesToRatio2    decimal.Decimal
}

func (s *ToRatio) Fill(cost Costs) {
	s.Costs = cost
	s.percentageOfIncome1 = s.percentageOfTotalIncome(s.Person1)
	s.percentageOfIncome2 = s.percentageOfTotalIncome(s.Person2)

	s.expensesToRatio1 = s.expensesToRatio(s.percentageOfIncome1)
	s.expensesToRatio2 = s.expensesToRatio(s.percentageOfIncome2)
}

func (s *ToRatio) Render() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleBold)
	t.AppendHeader(table.Row{"#", "income", "Inc. Ratio", "Pays"})
	t.AppendRows([]table.Row{
		{1, s.Person1.Round(2).String(), s.percentageOfIncome1.Round(2).String(), s.expensesToRatio1.Round(2).String()},
		{2, s.Person2.Round(2).String(), s.percentageOfIncome2.Round(2).String(), s.expensesToRatio2.Round(2).String()},
	})
	t.AppendFooter(s.footer())
	t.Render()
}

func (s *ToRatio) expensesToRatio(incomeRatio decimal.Decimal) decimal.Decimal {
	percent := s.TotalExpenses.Div(decimal.NewFromInt(hundred))
	return percent.Mul(incomeRatio)
}

func (s *ToRatio) percentageOfTotalIncome(income decimal.Decimal) decimal.Decimal {
	x := income.Div(s.TotalIncome)
	return x.Mul(decimal.NewFromInt(hundred))
}
