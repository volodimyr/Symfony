package domain

import "fmt"

var (
	ErrCarsNotFound      = fmt.Errorf("no cars were found")
	ErrOrdersNotFound    = fmt.Errorf("no orders were found")
	ErrCarIsNotAvailable = fmt.Errorf("car is not available")
)
