package tools

import (
	"github.com/shopspring/decimal"
	"math"
)

//go 精度计算

// Add 精确加法
func Add(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Add(decimal.NewFromFloat(f2)).Float64()
	return res
}

// Sub 精确减法
func Sub(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Sub(decimal.NewFromFloat(f2)).Float64()
	return res
}
func Int64SubFloat(f1 int64, f2 float64) float64 {
	res, _ := decimal.NewFromInt(f1).Sub(decimal.NewFromFloat(f2)).Float64()
	return res
}
func Int64SubInt64(f1, f2 int64) int64 {
	res := decimal.NewFromInt(f1).Sub(decimal.NewFromInt(f2)).IntPart()
	return res
}

// Mul 精确乘法
func Mul(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2)).Float64()
	return res
}
func Float64MulInt64(f1 float64, i1 int64) float64 {
	res, _ := decimal.NewFromFloat(f1).Mul(decimal.NewFromInt(i1)).Float64()
	return res
}
func Int64MulInt64(i1 int64, i2 int64) float64 {
	res, _ := decimal.NewFromInt(i1).Mul(decimal.NewFromInt(i2)).Float64()
	return res
}
func Int64MulFloat64(i1 int64, f1 float64) float64 {
	res, _ := decimal.NewFromInt(i1).Mul(decimal.NewFromFloat(f1)).Float64()
	return res
}

// Div 精确除法
func Div(f1, f2 float64) float64 {
	res, _ := decimal.NewFromFloat(f1).Div(decimal.NewFromFloat(f2)).Float64()
	return res
}
func Int64DivInt64(i1, i2 int64) float64 {
	res, _ := decimal.NewFromInt(i1).Div(decimal.NewFromInt(i2)).Float64()
	return res
}
func Float64DivInt64(f1 float64, i2 int64) float64 {
	res, _ := decimal.NewFromFloat(f1).Div(decimal.NewFromInt(i2)).Float64()
	return res
}
func Int64DivFloat64(i2 int64, f1 float64) float64 {
	res, _ := decimal.NewFromInt(i2).Div(decimal.NewFromFloat(f1)).Float64()
	return res
}

// IntPart 返回小数的整数部分
func IntPart(f float64) int64 {
	return decimal.NewFromFloat(f).IntPart()
}

//decimal=>float64
func Float(d decimal.Decimal) float64 {
	f, exact := d.Float64()
	if !exact {
		return f
	}
	return 0
}
func Bccomp(a, b float64) int64 {
	var Accuracy float64
	if math.Abs(a-b) < Accuracy {
		return 0
	}
	if math.Max(a, b) == a {
		return 1
	} else {
		return -1
	}
}
