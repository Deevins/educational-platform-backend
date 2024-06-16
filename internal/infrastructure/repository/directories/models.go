// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package directories

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

type HumanResourcesCourseTypes string

const (
	HumanResourcesCourseTypesCourse         HumanResourcesCourseTypes = "course"
	HumanResourcesCourseTypesPracticeCourse HumanResourcesCourseTypes = "practice_course"
)

func (e *HumanResourcesCourseTypes) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = HumanResourcesCourseTypes(s)
	case string:
		*e = HumanResourcesCourseTypes(s)
	default:
		return fmt.Errorf("unsupported scan type for HumanResourcesCourseTypes: %T", src)
	}
	return nil
}

type NullHumanResourcesCourseTypes struct {
	HumanResourcesCourseTypes HumanResourcesCourseTypes
	Valid                     bool // Valid is true if HumanResourcesCourseTypes is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullHumanResourcesCourseTypes) Scan(value interface{}) error {
	if value == nil {
		ns.HumanResourcesCourseTypes, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.HumanResourcesCourseTypes.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullHumanResourcesCourseTypes) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.HumanResourcesCourseTypes), nil
}

type HumanResourcesRoles string

const (
	HumanResourcesRolesADMIN     HumanResourcesRoles = "ADMIN"
	HumanResourcesRolesMODERATOR HumanResourcesRoles = "MODERATOR"
	HumanResourcesRolesUSER      HumanResourcesRoles = "USER"
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
	ID                     int32
	AuthorID               int32
	Title                  string
	Subtitle               *string
	Description            string
	AvatarUrl              *string
	StudentsCount          *int32
	RatingsCount           *int32
	Rating                 *float64
	CategoryTitle          string
	SubcategoryID          *int32
	Language               *string
	Level                  *string
	TimePlanned            *string
	CourseGoals            []string
	Requirements           []string
	TargetAudience         []string
	Type                   HumanResourcesCourseTypes
	Status                 HumanResourcesCourseStatuses
	LecturesLengthInterval pgtype.Interval
	LecturesCount          *int32
	PreviewVideoUrl        *string
	CreatedAt              pgtype.Timestamptz
	UpdatedAt              pgtype.Timestamptz
}

type HumanResourcesCoursesAttendant struct {
	CourseID  int32
	UserID    int32
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
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
	UserID    int32
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

type HumanResourcesInstructorsInfo struct {
	UserID             int32
	PreviousExperience string
	VideoKnowledge     string
	CurrentAudience    string
	CreatedAt          pgtype.Timestamptz
	UpdatedAt          pgtype.Timestamptz
}

type HumanResourcesLanguage struct {
	ID        int32
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type HumanResourcesLastCourseItemsVisitedByUser struct {
	ID         int32
	CourseID   int32
	UserID     int32
	LastItemID int32
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
}

type HumanResourcesLecture struct {
	ID           int32
	Title        string
	Description  string
	SectionID    int32
	SerialNumber int32
	VideoUrl     string
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

type HumanResourcesLecturesResource struct {
	ID          int32
	LectureID   int32
	Title       string
	Extension   string
	ResourceUrl string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type HumanResourcesNotification struct {
	ID        int32
	UserID    int32
	Message   string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type HumanResourcesSection struct {
	ID           int32
	Title        string
	Description  string
	CourseID     int32
	SerialNumber int32
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

type HumanResourcesSkillLevel struct {
	ID        int32
	Name      string
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
	ID           int32
	Name         string
	SectionID    int32
	SerialNumber int32
	Description  string
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

type HumanResourcesTestsQuestion struct {
	ID        int32
	TestID    int32
	Body      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type HumanResourcesTestsQuestionsAnswer struct {
	ID          int32
	QuestionID  int32
	Body        string
	Description string
	IsCorrect   bool
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type HumanResourcesTestsUsersAttempt struct {
	ID            int32
	TestID        int32
	UserID        int32
	AttemptNumber int32
	Score         int32
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
}

type HumanResourcesThreadsTag struct {
	ID   int32
	Name string
}

type HumanResourcesThreadsTagsLink struct {
	ID       int32
	ThreadID int32
	TagID    int32
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
	Role                   NullHumanResourcesRoles
	StudentsCount          int32
	CoursesCount           int32
	InstructorRating       float64
}
