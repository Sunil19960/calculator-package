package calculator

import (
	"fmt"
)

func Calculator(operation string, firstVal, secondVal float64) (float64, error) {
	switch operation {
	case "add":
		return firstVal + secondVal, nil
	case "subtract":
		return firstVal - secondVal, nil
	case "multiply":
		return firstVal * secondVal, nil
	case "divide":
		if secondVal == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return firstVal / secondVal, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}

}
