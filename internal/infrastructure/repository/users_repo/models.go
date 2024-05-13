// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package users_repo

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type HumanResourcesCourseStatuses string

const (
	HumanResourcesCourseStatusesDRAFT   HumanResourcesCourseStatuses = "DRAFT"
	HumanResourcesCourseStatusesPENDING HumanResourcesCourseStatuses = "PENDING"
	HumanResourcesCourseStatusesREADY   HumanResourcesCourseStatuses = "READY"
)

func (e *HumanResourcesCourseStatuses) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = HumanResourcesCourseStatuses(s)
	case string:
		*e = HumanResourcesCourseStatuses(s)
	default:
		return fmt.Errorf("unsupported scan type for HumanResourcesCourseStatuses: %T", src)
	}
	return nil
}

type NullHumanResourcesCourseStatuses struct {
	HumanResourcesCourseStatuses HumanResourcesCourseStatuses
	Valid                        bool // Valid is true if HumanResourcesCourseStatuses is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullHumanResourcesCourseStatuses) Scan(value interface{}) error {
	if value == nil {
		ns.HumanResourcesCourseStatuses, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.HumanResourcesCourseStatuses.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullHumanResourcesCourseStatuses) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.HumanResourcesCourseStatuses), nil
}

type HumanResourcesRoles string

const (
	HumanResourcesRolesAdmin     HumanResourcesRoles = "admin"
	HumanResourcesRolesModerator HumanResourcesRoles = "moderator"
	HumanResourcesRolesUsers     HumanResourcesRoles = "users"
)

func (e *HumanResourcesRoles) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = HumanResourcesRoles(s)
	case string:
		*e = HumanResourcesRoles(s)
	default:
		return fmt.Errorf("unsupported scan type for HumanResourcesRoles: %T", src)
	}
	return nil
}

type NullHumanResourcesRoles struct {
	HumanResourcesRoles HumanResourcesRoles
	Valid               bool // Valid is true if HumanResourcesRoles is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullHumanResourcesRoles) Scan(value interface{}) error {
	if value == nil {
		ns.HumanResourcesRoles, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.HumanResourcesRoles.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullHumanResourcesRoles) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.HumanResourcesRoles), nil
}

type HumanResourcesCategory struct {
	ID        int32
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type HumanResourcesCourse struct {
	ID          int32
	Name        string
	Description string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type HumanResourcesCoursesReview struct {
	ID        int32
	CourseID  int32
	UserID    int32
	Rating    int32
	Review    string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type HumanResourcesForumThread struct {
	ID        int32
	Title     string
	Body      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type HumanResourcesForumThreadsMessage struct {
	ID         int32
	ThreadID   int32
	Body       string
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
	UserSentID int32
}

type HumanResourcesLecture struct {
	ID          int32
	Title       string
	Description string
	VideoUrl    string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type HumanResourcesNotification struct {
	ID        int32
	Message   string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type HumanResourcesSubcategory struct {
	ID         int32
	Name       string
	CategoryID int32
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
}

type HumanResourcesTest struct {
	ID        int32
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type HumanResourcesUser struct {
	ID                     int32
	FullName               string
	Description            *string
	AvatarUrl              *string
	Email                  string
	PasswordHashed         string
	CreatedAt              pgtype.Timestamptz
	UpdatedAt              pgtype.Timestamptz
	HasUserTriedInstructor *bool
	PhoneNumber            string
}
