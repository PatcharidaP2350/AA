package entity

import "gorm.io/gorm"

type TakeAHistory struct {
	gorm.Model
	Weight float32 `valid:"required~Weight is required"`
	Email string `valid:"required~Email is required,email~Email is correct"`
	PhoneNumber string `valid:"required~PhoneNumber is required,matches(^[0-9]{10}$)~PhoneNumber is correct"`
}