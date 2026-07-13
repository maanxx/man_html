package entity

import "gorm.io/gorm"


type UserEntity struct {
  gorm.Model
  Name         string
  CompanyRefer int
  Company      Company `gorm:"foreignKey:CompanyRefer"`
  CreatedBy	uint
  CreatedInfo *UserEntity `gorm:"foreignKey:CreatedBy"`
  // use CompanyRefer as foreign key
}

type Company struct {
  ID   int
  Name string
}