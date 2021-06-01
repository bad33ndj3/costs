package sharedexpenses

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSplitEqual_Fill(t *testing.T) {
	type args struct {
		cost Costs
	}
	tests := []struct {
		name   string
		args   args
		inspect func(t *testing.T, s *SplitEqual)
	}{
		{
			name:    "happy",
			args:    args{
				cost: Costs{
					Person1:       decimal.NewFromInt(7000),
					Person2:       decimal.NewFromInt(2000),
					TotalExpenses: decimal.NewFromInt(1500),
					TotalIncome:   decimal.NewFromInt(9000),
				},
			},
			inspect: func(t *testing.T, s *SplitEqual) {
				assert.Equal(t, decimal.NewFromInt(750).String(), s.dividedExpenses.String())
			},
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SplitEqual{}
			s.Fill(tt.args.cost)
			tt.inspect(t, s)
		})
	}
}
