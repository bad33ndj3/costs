package sharedexpenses

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestToRatio_Fill(t *testing.T) {
	type args struct {
		cost Costs
	}
	tests := []struct {
		name    string
		args    args
		inspect func(t *testing.T, toRatio *ToRatio)
	}{
		{
			name: "50/50",
			args: args{
				cost: Costs{
					Person1:       decimal.NewFromInt(1000),
					Person2:       decimal.NewFromInt(1000),
					TotalExpenses: decimal.NewFromInt(1000),
					TotalIncome:   decimal.NewFromInt(2000),
				},
			},
			inspect: func(t *testing.T, toRatio *ToRatio) {
				assert.Equal(t, decimal.NewFromInt(50).String(), toRatio.percentageOfIncome1.String())
				assert.Equal(t, decimal.NewFromInt(50).String(), toRatio.percentageOfIncome2.String())
				assert.Equal(t, decimal.NewFromInt(500).String(), toRatio.expensesToRatio1.String())
				assert.Equal(t, decimal.NewFromInt(500).String(), toRatio.expensesToRatio2.String())
			},
		},
		{
			name: "30/70",
			args: args{
				cost: Costs{
					Person1:       decimal.NewFromInt(700),
					Person2:       decimal.NewFromInt(300),
					TotalExpenses: decimal.NewFromInt(1000),
					TotalIncome:   decimal.NewFromInt(1000),
				},
			},
			inspect: func(t *testing.T, toRatio *ToRatio) {
				assert.Equal(t, decimal.NewFromInt(70).String(), toRatio.percentageOfIncome1.String())
				assert.Equal(t, decimal.NewFromInt(30).String(), toRatio.percentageOfIncome2.String())
				assert.Equal(t, decimal.NewFromInt(700).String(), toRatio.expensesToRatio1.String())
				assert.Equal(t, decimal.NewFromInt(300).String(), toRatio.expensesToRatio2.String())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ToRatio{}
			s.Fill(tt.args.cost)
			tt.inspect(t, s)
		})
	}
}

func TestToRatio_expensesToRatio(t *testing.T) {
	type fields struct {
		Costs               Costs
		percentageOfIncome1 decimal.Decimal
		percentageOfIncome2 decimal.Decimal
		expensesToRatio1    decimal.Decimal
		expensesToRatio2    decimal.Decimal
	}
	type args struct {
		incomeRatio decimal.Decimal
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   decimal.Decimal
	}{
		{
			name: "50/50",
			fields: fields{
				Costs: Costs{
					Person1:       decimal.NewFromInt(1000),
					Person2:       decimal.NewFromInt(1000),
					TotalExpenses: decimal.NewFromInt(2000),
					TotalIncome:   decimal.NewFromInt(2000),
				},
			},
			args: args{
				incomeRatio: decimal.NewFromInt(50),
			},
			want: decimal.NewFromInt(1000),
		},
		{
			name: "30/70",
			fields: fields{
				Costs: Costs{
					Person1:       decimal.NewFromInt(300),
					Person2:       decimal.NewFromInt(700),
					TotalExpenses: decimal.NewFromInt(500),
					TotalIncome:   decimal.NewFromInt(1000),
				},
			},
			args: args{
				incomeRatio: decimal.NewFromInt(30),
			},
			want: decimal.NewFromInt(150),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ToRatio{
				Costs:               tt.fields.Costs,
				percentageOfIncome1: tt.fields.percentageOfIncome1,
				percentageOfIncome2: tt.fields.percentageOfIncome2,
				expensesToRatio1:    tt.fields.expensesToRatio1,
				expensesToRatio2:    tt.fields.expensesToRatio2,
			}
			assert.Equal(t, tt.want.String(), s.expensesToRatio(tt.args.incomeRatio).String())
		})
	}
}

func TestToRatio_percentageOfTotalIncome(t *testing.T) {
	type fields struct {
		Costs               Costs
		percentageOfIncome1 decimal.Decimal
		percentageOfIncome2 decimal.Decimal
		expensesToRatio1    decimal.Decimal
		expensesToRatio2    decimal.Decimal
	}
	type args struct {
		income decimal.Decimal
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   decimal.Decimal
	}{
		{
			name: "50/50",
			fields: fields{
				Costs: Costs{
					TotalIncome: decimal.NewFromInt(2000),
				},
			},
			args: args{
				income: decimal.NewFromInt(1000),
			},
			want: decimal.NewFromInt(50),
		},
		{
			name: "25-percentage-of-total-income",
			fields: fields{
				Costs: Costs{
					TotalIncome: decimal.NewFromInt(2000),
				},
			},
			args: args{
				income: decimal.NewFromInt(500),
			},
			want: decimal.NewFromInt(25),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ToRatio{
				Costs:               tt.fields.Costs,
				percentageOfIncome1: tt.fields.percentageOfIncome1,
				percentageOfIncome2: tt.fields.percentageOfIncome2,
				expensesToRatio1:    tt.fields.expensesToRatio1,
				expensesToRatio2:    tt.fields.expensesToRatio2,
			}
			assert.Equal(t, tt.want.String(), s.percentageOfTotalIncome(tt.args.income).String())
		})
	}
}
