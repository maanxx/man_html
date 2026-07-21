package requests

import (
	"app/modules/admins/common/consts"
	"app/modules/admins/common/errs"
	"app/modules/admins/users/transport/rules"
	"context"
	"strings"

	"github.com/go-playground/validator/v10"
)

// payload
type PayloadUserCreation struct {
	ID        int    `json:"id"`
	LastName  string `json:"last_name" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Status    int    `json:"status" validate:"required,ruleStatusUserCreate"`
	StatusStr string `json:"-"`
}

func (req *PayloadUserCreation) Validation(ctx context.Context) []*string {
	validation := validator.New()
	validation.RegisterValidation("ruleStatusUserCreate", rules.RuleStatusUserCreate)
	err := validation.Struct(req)
	var validationErrors []*string
	if err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {
			switch vErr.Field() {
			case "LastName":
				errLastName := errs.ErrLastNameValidate.Error()
				validationErrors = append(validationErrors, &errLastName)
			case "FirstName":
				errFirstName := errs.ErrFirstNameValidate.Error()
				validationErrors = append(validationErrors, &errFirstName)
			case "Email":
				errEmail := errs.ErrEmailValidate.Error()
				validationErrors = append(validationErrors, &errEmail)
			case "Status":
				errStatus := errs.ErrStatusNotFound.Error()
				validationErrors = append(validationErrors, &errStatus)
			}
		}
	}
	if validationErrors != nil {
		return validationErrors
	}

	req.StatusStr = consts.MapStatusInt[req.Status]

	req.LastName = strings.TrimSpace(req.LastName)
	req.FirstName = strings.TrimSpace(req.FirstName)
	req.Email = strings.TrimSpace(req.Email)

	return nil
}
