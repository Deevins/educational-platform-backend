// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package users

import (
	"context"
)

type Querier interface {
	AddTeachingExperience(ctx context.Context, arg *AddTeachingExperienceParams) (int32, error)
	CheckIfUserRegisteredToCourse(ctx context.Context, arg *CheckIfUserRegisteredToCourseParams) (*HumanResourcesCoursesAttendant, error)
	CreateUser(ctx context.Context, arg *CreateUserParams) (int32, error)
	GetHasUserTriedInstructor(ctx context.Context, id int32) (*bool, error)
	GetUserByEmailAndHashedPassword(ctx context.Context, arg *GetUserByEmailAndHashedPasswordParams) (*HumanResourcesUser, error)
	GetUserByID(ctx context.Context, id int32) (*HumanResourcesUser, error)
	GetUsers(ctx context.Context) ([]*HumanResourcesUser, error)
	RegisterToCourse(ctx context.Context, arg *RegisterToCourseParams) (int32, error)
	UpdateAvatar(ctx context.Context, arg *UpdateAvatarParams) (int32, error)
	UpdateHasUserTriedInstructor(ctx context.Context, id int32) (int32, error)
	UpdateUserDescription(ctx context.Context, arg *UpdateUserDescriptionParams) (int32, error)
	UpdateUserEmail(ctx context.Context, arg *UpdateUserEmailParams) (int32, error)
	UpdateUserFullName(ctx context.Context, arg *UpdateUserFullNameParams) (int32, error)
	UpdateUserPhone(ctx context.Context, arg *UpdateUserPhoneParams) (int32, error)
}

var _ Querier = (*Queries)(nil)
