package requests

import (
	"context"

	"github.com/go-playground/validator/v10"
)

// payload
type PayloadChangeStatus struct {
	ID     int `json:"id" validate:"required"`
	Status int `json:"status" validate:"required"`
}

func (req *PayloadChangeStatus) Validation(ctx context.Context) []*string {
	validation := validator.New()
	err := validation.Struct(req)
	var validationErrors []*string
	if err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {
			switch vErr.Field() {
			case "ID":
				errID := "ID is required"
				validationErrors = append(validationErrors, &errID)
			case "Status":
				errStatus := "Status is required"
				validationErrors = append(validationErrors, &errStatus)
			}
		}
	}
	if validationErrors != nil {
		return validationErrors
	}
	return nil
}
