package points

import (
	"github.com/hxcuber/receipt-processor/pkg/model"
	"testing"
	"time"
)

func TestCalcPoints(t *testing.T) {
	type args struct {
		receipt model.Receipt
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example_1",
			args: args{
				receipt: model.Receipt{
					Retailer:     "Target",
					PurchaseDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					PurchaseTime: time.Date(0, 0, 0, 13, 1, 0, 0, time.UTC),
					Items: []model.Item{
						{
							ShortDescription: "Mountain Dew 12PK",
							Price:            6.49,
						}, {
							ShortDescription: "Emils Cheese Pizza",
							Price:            12.25,
						}, {
							ShortDescription: "Knorr Creamy Chicken",
							Price:            1.26,
						}, {
							ShortDescription: "Doritos Nacho Cheese",
							Price:            3.35,
						}, {
							ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
							Price:            12.00,
						},
					},
					Total: 35.35,
				},
			},
			want: 28,
		}, {
			name: "example_2",
			args: args{
				model.Receipt{
					Retailer:     "M&M Corner Market",
					PurchaseDate: time.Date(2022, 03, 20, 0, 0, 0, 0, time.UTC),
					PurchaseTime: time.Date(0, 0, 0, 14, 33, 0, 0, time.UTC),
					Items: []model.Item{
						{
							ShortDescription: "Gatorade",
							Price:            2.25,
						}, {
							ShortDescription: "Gatorade",
							Price:            2.25,
						}, {
							ShortDescription: "Gatorade",
							Price:            2.25,
						}, {
							ShortDescription: "Gatorade",
							Price:            2.25,
						},
					},
					Total: 9.00,
				},
			},
			want: 109,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcPoints(tt.args.receipt); got != tt.want {
				t.Errorf("CalcPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
