package sharedexpenses

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/shopspring/decimal"
)

type SplitEqual struct {
	Costs
	dividedExpenses decimal.Decimal
}


func (s *SplitEqual) Fill(cost Costs) {
	s.Costs = cost
	s.dividedExpenses = s.TotalExpenses.Div(decimal.NewFromInt(two))
}

func (s *SplitEqual) Render() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "income", "Pays"})
	t.AppendRows([]table.Row{
		{1, s.Person1.Round(2).String(), s.dividedExpenses.Round(2).String()},
		{2, s.Person2.Round(2).String(), s.dividedExpenses.Round(2).String()},
	})
	t.AppendFooter(table.Row{"total", s.TotalIncome.String(), s.TotalExpenses.String()})
	t.Render()
}
