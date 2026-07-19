package requests

import (
	"app/modules/admins/common/errs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PayloadUserDelete struct {
	Ids []primitive.ObjectID `json:"ids" validate:"required"`
}

func (pl *PayloadUserDelete) Validate() error {
	validate := validator.New()
	err := validate.Struct(pl)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Ids":
				return errs.ErrIDUserValidate
			}
		}
	}
	return nil
}
