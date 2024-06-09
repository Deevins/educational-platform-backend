package fixtures

import (
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	models_dto "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/models"
)

type UserBuilder struct {
	instance *models_dto.UserDTO
}

func User() *UserBuilder {
	return &UserBuilder{
		instance: &models_dto.UserDTO{},
	}
}

func (ob *UserBuilder) ID(v models.UserID) *UserBuilder {
	ob.instance.ID = v
	return ob
}

func (ob *UserBuilder) Username(v models.Username) *UserBuilder {
	ob.instance.Username = v
	return ob
}

func (ob *UserBuilder) Email(v models.UserEmail) *UserBuilder {
	ob.instance.Email = v
	return ob
}

func (ob *UserBuilder) P() *models_dto.UserDTO {
	return ob.instance
}

func (ob *UserBuilder) V() models_dto.UserDTO {
	return *ob.instance
}

func (ob *UserBuilder) Valid() *UserBuilder {
	return ob.
		ID(1).
		Username("Div").
		Email("tests@tests.tests")
}
