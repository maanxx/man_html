package entity

import "github.com/teoit/gosctx/core"

type UserGroup struct {
	core.SQLModel
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      int        `json:"status"`
	CreatedBy   int64      `json:"created_by"`
	UpdatedBy   int64      `json:"updated_by"`
	UserCreated *UserGroup `gorm:"foreignKey:CreatedBy"`
	UserUpdated *UserGroup `gorm:"foreignKey:UpdatedBy"`
}

func (UserGroup) TableName() string {
	return "user_groups"
}
