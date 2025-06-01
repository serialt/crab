package crab

import (
	"errors"

	"github.com/shopspring/decimal"
)

// GetDecimal get decimal from string
func GetDecimal(data string) (d decimal.Decimal) {
	d, _ = decimal.NewFromString(data)
	return
}

// GetDecimalOrZero get decimal from string,when decimal less than or equal zero, return zero
func GetDecimalOrZero(data string) (d decimal.Decimal) {
	d, _ = decimal.NewFromString(data)
	if d.LessThanOrEqual(decimal.Zero) {
		d = decimal.Zero
	}
	return
}

// GetDecimalToString decimal to string, when decimal is zero return ""
func GetDecimalToString(data decimal.Decimal) (st string) {
	st = data.String()
	if st == "0" {
		st = ""
	}
	return
}

// GetDecimalSum decimal sum
func GetDecimalSum(items ...decimal.Decimal) (st string) {
	sum := decimal.Decimal{}
	for _, v := range items {
		sum = sum.Add(v)
	}
	st = sum.String()
	return
}

// GetNumRoundDecimal get decimal string, 取小数位数
func GetNumRoundDecimal(num decimal.Decimal, round int64) (data string) {
	if num.IsZero() {
		return
	}
	if round == -1 {
		data = num.String()
		return
	} else {
		data = num.Round(int32(round)).String()
	}
	return
}

// GetNumRoundDecimalV2 get decimal string, 取小数位数
func GetNumRoundDecimalV2(num string, round int64) (data string) {
	if num == "" {
		return
	}
	if round == -1 {
		data = num
		return
	} else {
		data = GetDecimal(num).Round(int32(round)).String()
	}
	return
}

func MinDecimal(list ...decimal.Decimal) (d decimal.Decimal, err error) {
	if len(list) == 0 {
		err = errors.New("list is nil")
		return
	}
	d = list[0]
	for _, item := range list {
		if d.LessThan(item) {
			d = item
		}
	}
	return
}

func MaxDecimal(list ...decimal.Decimal) (d decimal.Decimal, err error) {
	if len(list) == 0 {
		err = errors.New("list is nil")
		return
	}
	d = list[0]
	for _, item := range list {
		if d.GreaterThan(item) {
			d = item
		}
	}
	return
}

func SumDecimal[V any](list []V, f func(item V) decimal.Decimal) decimal.Decimal {
	sum := decimal.Zero
	for i := range list {
		sum = sum.Add(f(list[i]))
	}
	return sum
}
