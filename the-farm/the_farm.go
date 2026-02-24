package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood calculates food per cow.
func DivideFood(calc FodderCalculator, cows int) (float64, error) {
	total, err := calc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := calc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (total * factor) / float64(cows), nil
}

// ValidateInputAndDivideFood validates cows count then delegates to DivideFood.
func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calc, cows)
}

type InvalidCowsError struct {
	cows int
	msg  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}

// ValidateNumberOfCows validates count and returns a custom error if invalid.
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows: cows,
			msg:  "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			cows: cows,
			msg:  "no cows don't need food",
		}
	}
	return nil
}
