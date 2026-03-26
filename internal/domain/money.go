package domain

import (
	"fmt"
)

type Money int64

func (m Money) Add(other Money) Money {
	return m + other
}

func (m Money) Sub(other Money) Money {
	return m - other
}

func (m Money) string() string {
	sign := ""
	value := m

	if m < 0 {
		sign = "-"
		value = -m
	}
	major := value / 100
	minor := value % 100

	return fmt.Sprintf("sign %s, value %d, major %d, minor %d", sign, value, major, minor)
}
