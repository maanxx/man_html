package rules

import (
	"app/modules/admins/common/consts"
	"slices"

	"github.com/go-playground/validator/v10"
)

func RuleStatusUser(fl validator.FieldLevel) bool {
	status := fl.Field().Interface().(int)
	return slices.Contains(consts.StatusUserToViewCreate, status)
}

func RuleStatusUserCreate(fl validator.FieldLevel) bool {
	status := fl.Field().Interface().(int)
	return slices.Contains(consts.StatusUserToViewCreate, status)
}
