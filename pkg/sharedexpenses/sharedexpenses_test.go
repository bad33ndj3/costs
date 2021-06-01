package sharedexpenses

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCosts_Init(t *testing.T) {
	type args struct {
		person1       *int
		person2       *int
		totalExpenses *int
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "happy",
			args: args{
				person1: func() *int {
					x := 1000
					return &x
				}(),
				person2: func() *int {
					x := 1000
					return &x
				}(),
				totalExpenses: func() *int {
					x := 1000
					return &x
				}(),
			},
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				cost, ok := i[0].(*Costs)
				if !assert.True(t, ok) {
					t.Errorf("assertion is not on type Costs")
				}
				assert.Equal(t, nil, err)
				assert.Equal(t, decimal.NewFromInt(1000), cost.Person1)
				assert.Equal(t, decimal.NewFromInt(1000), cost.Person2)
				assert.Equal(t, decimal.NewFromInt(2000), cost.TotalIncome)
				assert.Equal(t, decimal.NewFromInt(1000), cost.TotalExpenses)

				return true
			},
		},
		{
			name: "total-expenses-not-set",
			args: args{
				person1: func() *int {
					x := 1000
					return &x
				}(),
				person2: func() *int {
					x := 1000
					return &x
				}(),
				totalExpenses: nil,
			},
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				assert.Equal(t, errMissingParameters, err)

				return true
			},
		},
		{
			name: "30/70",
			args: args{
				person1: func() *int {
					x := 300
					return &x
				}(),
				person2: func() *int {
					x := 700
					return &x
				}(),
				totalExpenses: func() *int {
					x := 1005
					return &x
				}(),
			},
			assertion: func(t assert.TestingT, err error, i ...interface{}) bool {
				cost, ok := i[0].(*Costs)
				if !assert.True(t, ok) {
					t.Errorf("assertion is not on type Costs")
				}
				assert.Equal(t, nil, err)
				assert.Equal(t, decimal.NewFromInt(300), cost.Person1)
				assert.Equal(t, decimal.NewFromInt(700), cost.Person2)
				assert.Equal(t, decimal.NewFromInt(1000), cost.TotalIncome)
				assert.Equal(t, decimal.NewFromInt(1005), cost.TotalExpenses)

				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Costs{}
			tt.assertion(t, c.Init(tt.args.person1, tt.args.person2, tt.args.totalExpenses), c)
		})
	}
}
