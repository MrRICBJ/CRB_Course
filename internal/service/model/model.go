package model

type Response struct {
	MaxValue, SumValue, MinValue, AvgValue float64
	MaxName, MaxDate, MinName, MinDate     string
	Count                                  int
}
