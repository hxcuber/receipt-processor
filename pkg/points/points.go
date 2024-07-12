package points

import (
	"math"
	"github.com/hxcuber/receipt-processor/pkg/model"
	"strings"
)

func CalcPoints(receipt model.Receipt) int {
	points := 0

	// One point for every alphanumeric char in retailer name
	for _, r := range receipt.Retailer {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			points++
		}
	}

	// 50 points if the total is a round dollar amount with no cents.
	if math.Floor(receipt.Total) == receipt.Total {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += len(receipt.Items) / 2 * 5

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, v := range receipt.Items {
		if len(strings.Trim(v.ShortDescription, " "))%3 == 0 {
			points += int(math.Ceil(v.Price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd.
	if _, _, day := receipt.PurchaseDate.Date(); day%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	hours, mins, _ := receipt.PurchaseTime.Clock()
	if (hours > 14 || hours == 14 && mins > 0) && (hours < 16) {
		points += 10
	}

	return points
}
