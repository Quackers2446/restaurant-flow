package util

import (
	"github.com/labstack/echo/v4"
)

// ValidateInput binds context to the input type T and validates it
// If input is a nil pointer, the function will instantiate a struct of type T for you
func ValidateInput[T interface{}](context *echo.Context, input *T) (*T, error) {
	if input == nil {
		var newInput T

		if err := (*context).Bind(&newInput); err != nil {
			return nil, err
		}
		if err := (*context).Validate(&newInput); err != nil {
			return nil, err
		}

		return &newInput, nil
	}

	if err := (*context).Bind(&input); err != nil {
		return nil, err
	}
	if err := (*context).Validate(&input); err != nil {
		return nil, err
	}

	return input, nil
}
