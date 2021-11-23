package converter

import "github.com/leekchan/accounting"

func GoalAmountFormatIDR(number int) string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(number)
}