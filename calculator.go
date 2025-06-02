package calculator

import (
	"fmt"
)

func Add(firstVal, secondVal float64) (float64, error) {
	return firstVal + secondVal, nil
}

func Subtract(firstVal, secondVal float64) (float64, error) {
	return firstVal - secondVal, nil
}

func Multiply(firstVal, secondVal float64) (float64, error) {
	return firstVal * secondVal, nil
}
func Divide(firstVal, secondVal float64) (float64, error) {
	if secondVal == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return firstVal / secondVal, nil
}

func Modulus(firstVal, secondVal float64) (float64, error) {
	if secondVal == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(int(firstVal) % int(secondVal)), nil
}
