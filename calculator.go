package calculator

import (
	"fmt"
)

// Add returns the sum of two float64 numbers.
//
// It takes two float64 values as input and returns their sum along with an error value.
// Example usage:
// result, err := Add(3.5, 2.5)
//
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//
//	    fmt.Println("Result:", result)
//	}
//
// Output: Result: 6.0
func Add(firstVal, secondVal float64) (float64, error) {
	return firstVal + secondVal, nil
}

// Subtract returns the difference of two float64 numbers.
//
// It takes two float64 values as input and returns their difference along with an error value.
// It does not handle any specific error cases, so it always returns nil for the error.
// Example usage:
// result, err := Subtract(10.5, 2.3)
//
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//
//	    fmt.Println("Result:", result)
//	}
//
// Output: Result: 8.2
func Subtract(firstVal, secondVal float64) (float64, error) {
	return firstVal - secondVal, nil
}

// Multiply returns the product of two float64 numbers.
//
// It takes two float64 values as input and returns their product along with an error value.
// Example usage:
// result, err := Multiply(3.0, 4.0)
//
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//
//	    fmt.Println("Result:", result)
//	}
//
// Output: Result: 12.0
func Multiply(firstVal, secondVal float64) (float64, error) {
	return firstVal * secondVal, nil
}

// Divide returns the quotient of two float64 numbers.
// It takes two float64 values as input and returns their quotient along with an error value.
// If the second value is zero, it returns an error indicating division by zero.
// Example usage:
// result, err := Divide(10.0, 2.0)
//
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//
//	    fmt.Println("Result:", result)
//	}
//
// Output: Result: 5.0
func Divide(firstVal, secondVal float64) (float64, error) {
	if secondVal == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return firstVal / secondVal, nil
}

// Modulus returns the modulus of two float64 numbers.
// It takes two float64 values as input and returns the modulus along with an error value.
// If the second value is zero, it returns an error indicating division by zero.
// Example usage:
// result, err := Modulus(10.0, 3.0)
//
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//
//	    fmt.Println("Result:", result)
//	}
//
// Output: Result: 1.0
func Modulus(firstVal, secondVal float64) (float64, error) {
	if secondVal == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return float64(int(firstVal) % int(secondVal)), nil
}

// Square returns the square of a float64 number.
// It takes a float64 value as input and returns its square along with an error value.
// Example usage:
// result, err := Square(4.0)
//
//	if err != nil {
//	    fmt.Println("Error:", err)
//	} else {
//
//	    fmt.Println("Result:", result)
//	}
//
// Output: Result: 16.0
func Square(value float64) (float64, error) {
	return value * value, nil
}
