// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package courses

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addCourseBasicInfo = `-- name: AddCourseBasicInfo :one
UPDATE human_resources.courses SET
        title = $1,
        subtitle = $2,
        description = $3,
        language = $4,
        level = $5,
        category_id = $6,
        updated_at = now()
                        WHERE id = $7 RETURNING id
`

type AddCourseBasicInfoParams struct {
	Title       string
	Subtitle    *string
	Description string
	Language    *string
	Level       *string
	CategoryID  *int32
	ID          int32
}

func (q *Queries) AddCourseBasicInfo(ctx context.Context, arg *AddCourseBasicInfoParams) (int32, error) {
	row := q.db.QueryRow(ctx, addCourseBasicInfo,
		arg.Title,
		arg.Subtitle,
		arg.Description,
		arg.Language,
		arg.Level,
		arg.CategoryID,
		arg.ID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const approveCourse = `-- name: ApproveCourse :one
UPDATE human_resources.courses SET status = 'READY', updated_at = now()  WHERE id = $1 RETURNING id
`

func (q *Queries) ApproveCourse(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, approveCourse, id)
	err := row.Scan(&id)
	return id, err
}

const createCourseBase = `-- name: CreateCourseBase :one
INSERT INTO human_resources.courses (
                title,
                type,
                author_id,
                category_id,
                time_planned)
VALUES ($1,
        $2,
        $3,
        $4,
        $5
       ) RETURNING id
`

type CreateCourseBaseParams struct {
	Title       string
	Type        HumanResourcesCourseTypes
	AuthorID    int32
	CategoryID  *int32
	TimePlanned *string
}

func (q *Queries) CreateCourseBase(ctx context.Context, arg *CreateCourseBaseParams) (int32, error) {
	row := q.db.QueryRow(ctx, createCourseBase,
		arg.Title,
		arg.Type,
		arg.AuthorID,
		arg.CategoryID,
		arg.TimePlanned,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createSection = `-- name: CreateSection :one
INSERT INTO human_resources.sections (title, description, course_id) VALUES ($1, $2, $3) RETURNING id
`

type CreateSectionParams struct {
	Title       string
	Description string
	CourseID    int32
}

func (q *Queries) CreateSection(ctx context.Context, arg *CreateSectionParams) (int32, error) {
	row := q.db.QueryRow(ctx, createSection, arg.Title, arg.Description, arg.CourseID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getAllDraftCourses = `-- name: GetAllDraftCourses :many
SELECT id, author_id, title, subtitle, description, avatar_url, students_count, ratings_count, rating, category_id, subcategory_id, language, level, time_planned, course_goals, requirements, target_audience, type, status, lectures_length, lectures_count, preview_video_url, created_at, updated_at FROM human_resources.courses c
WHERE c.status != 'PENDING'
  AND c.status != 'READY'
ORDER BY
    c.created_at
`

func (q *Queries) GetAllDraftCourses(ctx context.Context) ([]*HumanResourcesCourse, error) {
	rows, err := q.db.Query(ctx, getAllDraftCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*HumanResourcesCourse
	for rows.Next() {
		var i HumanResourcesCourse
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Title,
			&i.Subtitle,
			&i.Description,
			&i.AvatarUrl,
			&i.StudentsCount,
			&i.RatingsCount,
			&i.Rating,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.Language,
			&i.Level,
			&i.TimePlanned,
			&i.CourseGoals,
			&i.Requirements,
			&i.TargetAudience,
			&i.Type,
			&i.Status,
			&i.LecturesLength,
			&i.LecturesCount,
			&i.PreviewVideoUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPendingCourses = `-- name: GetAllPendingCourses :many
SELECT id, author_id, title, subtitle, description, avatar_url, students_count, ratings_count, rating, category_id, subcategory_id, language, level, time_planned, course_goals, requirements, target_audience, type, status, lectures_length, lectures_count, preview_video_url, created_at, updated_at  FROM human_resources.courses c
WHERE c.status != 'DRAFT'
  AND c.status != 'READY'
ORDER BY
    c.created_at
`

func (q *Queries) GetAllPendingCourses(ctx context.Context) ([]*HumanResourcesCourse, error) {
	rows, err := q.db.Query(ctx, getAllPendingCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*HumanResourcesCourse
	for rows.Next() {
		var i HumanResourcesCourse
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Title,
			&i.Subtitle,
			&i.Description,
			&i.AvatarUrl,
			&i.StudentsCount,
			&i.RatingsCount,
			&i.Rating,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.Language,
			&i.Level,
			&i.TimePlanned,
			&i.CourseGoals,
			&i.Requirements,
			&i.TargetAudience,
			&i.Type,
			&i.Status,
			&i.LecturesLength,
			&i.LecturesCount,
			&i.PreviewVideoUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllReadyCourses = `-- name: GetAllReadyCourses :many
SELECT id, author_id, title, subtitle, description, avatar_url, students_count, ratings_count, rating, category_id, subcategory_id, language, level, time_planned, course_goals, requirements, target_audience, type, status, lectures_length, lectures_count, preview_video_url, created_at, updated_at FROM human_resources.courses c
WHERE c.status != 'DRAFT'
  AND c.status != 'PENDING'
ORDER BY
    c.created_at
`

func (q *Queries) GetAllReadyCourses(ctx context.Context) ([]*HumanResourcesCourse, error) {
	rows, err := q.db.Query(ctx, getAllReadyCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*HumanResourcesCourse
	for rows.Next() {
		var i HumanResourcesCourse
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Title,
			&i.Subtitle,
			&i.Description,
			&i.AvatarUrl,
			&i.StudentsCount,
			&i.RatingsCount,
			&i.Rating,
			&i.CategoryID,
			&i.SubcategoryID,
			&i.Language,
			&i.Level,
			&i.TimePlanned,
			&i.CourseGoals,
			&i.Requirements,
			&i.TargetAudience,
			&i.Type,
			&i.Status,
			&i.LecturesLength,
			&i.LecturesCount,
			&i.PreviewVideoUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCourseAvatarByID = `-- name: GetCourseAvatarByID :one
SELECT avatar_url FROM human_resources.courses WHERE id = $1
`

func (q *Queries) GetCourseAvatarByID(ctx context.Context, id int32) (*string, error) {
	row := q.db.QueryRow(ctx, getCourseAvatarByID, id)
	var avatar_url *string
	err := row.Scan(&avatar_url)
	return avatar_url, err
}

const getCoursePreviewVideoByID = `-- name: GetCoursePreviewVideoByID :one
SELECT preview_video_url FROM human_resources.courses WHERE id = $1
`

func (q *Queries) GetCoursePreviewVideoByID(ctx context.Context, id int32) (*string, error) {
	row := q.db.QueryRow(ctx, getCoursePreviewVideoByID, id)
	var preview_video_url *string
	err := row.Scan(&preview_video_url)
	return preview_video_url, err
}

const getCourseSections = `-- name: GetCourseSections :many
SELECT s.id, s.title, s.description FROM human_resources.sections s
                    INNER JOIN human_resources.courses cs ON s.course_id = cs.id
WHERE cs.id = $1
ORDER BY
    s.created_at DESC
`

type GetCourseSectionsRow struct {
	ID          int32
	Title       string
	Description string
}

func (q *Queries) GetCourseSections(ctx context.Context, courseID int32) ([]*GetCourseSectionsRow, error) {
	rows, err := q.db.Query(ctx, getCourseSections, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCourseSectionsRow
	for rows.Next() {
		var i GetCourseSectionsRow
		if err := rows.Scan(&i.ID, &i.Title, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCoursesAvatarsByIDs = `-- name: GetCoursesAvatarsByIDs :many
SELECT id, avatar_url FROM human_resources.courses WHERE id = ANY($1::int[])
`

type GetCoursesAvatarsByIDsRow struct {
	ID        int32
	AvatarUrl *string
}

func (q *Queries) GetCoursesAvatarsByIDs(ctx context.Context, dollar_1 []int32) ([]*GetCoursesAvatarsByIDsRow, error) {
	rows, err := q.db.Query(ctx, getCoursesAvatarsByIDs, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCoursesAvatarsByIDsRow
	for rows.Next() {
		var i GetCoursesAvatarsByIDsRow
		if err := rows.Scan(&i.ID, &i.AvatarUrl); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFullCourseInfoWithInstructorByCourseID = `-- name: GetFullCourseInfoWithInstructorByCourseID :one
SELECT c.id as course_id, c.title, c.subtitle,
       c.description, c.language, c.level, u.description as instructor_description,
       c.rating, c.students_count, c.ratings_count, c.lectures_count,
       c.lectures_length, c.avatar_url as course_avatar_url, c.preview_video_url,
       c.status as course_status, c.created_at as course_created_at, c.course_goals, c.requirements,
       c.target_audience, c.author_id, cg.name as category_title, u.full_name as instructor_full_name,
       u.students_count as instructor_students_count, u.courses_count as instructor_courses_count, u.instructor_rating,
       u.avatar_url as instructor_avatar_url, u.id as instructor_id FROM human_resources.courses c
                     JOIN human_resources.users u ON c.author_id = u.id
                     JOIN human_resources.categories cg on cg.id = c.category_id
WHERE c.id = $1
`

type GetFullCourseInfoWithInstructorByCourseIDRow struct {
	CourseID                int32
	Title                   string
	Subtitle                *string
	Description             string
	Language                *string
	Level                   *string
	InstructorDescription   *string
	Rating                  *float64
	StudentsCount           *int32
	RatingsCount            *int32
	LecturesCount           *int32
	LecturesLength          *int32
	CourseAvatarUrl         *string
	PreviewVideoUrl         *string
	CourseStatus            HumanResourcesCourseStatuses
	CourseCreatedAt         pgtype.Timestamptz
	CourseGoals             []string
	Requirements            []string
	TargetAudience          []string
	AuthorID                int32
	CategoryTitle           string
	InstructorFullName      string
	InstructorStudentsCount int32
	InstructorCoursesCount  int32
	InstructorRating        int32
	InstructorAvatarUrl     *string
	InstructorID            int32
}

func (q *Queries) GetFullCourseInfoWithInstructorByCourseID(ctx context.Context, id int32) (*GetFullCourseInfoWithInstructorByCourseIDRow, error) {
	row := q.db.QueryRow(ctx, getFullCourseInfoWithInstructorByCourseID, id)
	var i GetFullCourseInfoWithInstructorByCourseIDRow
	err := row.Scan(
		&i.CourseID,
		&i.Title,
		&i.Subtitle,
		&i.Description,
		&i.Language,
		&i.Level,
		&i.InstructorDescription,
		&i.Rating,
		&i.StudentsCount,
		&i.RatingsCount,
		&i.LecturesCount,
		&i.LecturesLength,
		&i.CourseAvatarUrl,
		&i.PreviewVideoUrl,
		&i.CourseStatus,
		&i.CourseCreatedAt,
		&i.CourseGoals,
		&i.Requirements,
		&i.TargetAudience,
		&i.AuthorID,
		&i.CategoryTitle,
		&i.InstructorFullName,
		&i.InstructorStudentsCount,
		&i.InstructorCoursesCount,
		&i.InstructorRating,
		&i.InstructorAvatarUrl,
		&i.InstructorID,
	)
	return &i, err
}

const getInstructorCourses = `-- name: GetInstructorCourses :many
SELECT c.id, c.avatar_url, c.title, c.status FROM human_resources.courses c WHERE author_id = $1 ORDER BY created_at DESC
`

type GetInstructorCoursesRow struct {
	ID        int32
	AvatarUrl *string
	Title     string
	Status    HumanResourcesCourseStatuses
}

func (q *Queries) GetInstructorCourses(ctx context.Context, authorID int32) ([]*GetInstructorCoursesRow, error) {
	rows, err := q.db.Query(ctx, getInstructorCourses, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetInstructorCoursesRow
	for rows.Next() {
		var i GetInstructorCoursesRow
		if err := rows.Scan(
			&i.ID,
			&i.AvatarUrl,
			&i.Title,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSectionsWithLecturesAndTestsByCourseID = `-- name: GetSectionsWithLecturesAndTestsByCourseID :many
WITH recursive_section AS (
    SELECT
        s.id AS section_id,
        s.title AS section_title,
        s.description AS section_description,
        s.course_id AS course_id,
        l.id AS lecture_id,
        l.title AS lecture_title,
        l.description AS lecture_description,
        l.video_url AS lecture_video_url,
        t.id AS test_id,
        t.name AS test_name,
        q.id AS question_id,
        q.body AS question_body
    FROM
        human_resources.sections s
            LEFT JOIN
        human_resources.lectures l ON s.id = l.section_id
            LEFT JOIN
        human_resources.tests t ON s.id = t.section_id
            LEFT JOIN
        human_resources.tests_questions q ON t.id = q.test_id
    WHERE
        s.course_id = $1
)
SELECT
    section_id,
    section_title,
    section_description,
    json_agg(
            json_build_object(
                    'lecture_id', lecture_id,
                    'lecture_title', lecture_title,
                    'lecture_description', lecture_description,
                    'lecture_video_url', lecture_video_url
            )
    ) AS lectures,
    json_agg(
            json_build_object(
                    'test_id', test_id,
                    'test_name', test_name,
                    'questions', json_agg(
                            json_build_object(
                                    'question_id', question_id,
                                    'question_body', question_body
                            )
                                 )
            )
    ) AS tests
FROM
    recursive_section
GROUP BY
    section_id, section_title, section_description
`

type GetSectionsWithLecturesAndTestsByCourseIDRow struct {
	SectionID          int32
	SectionTitle       string
	SectionDescription string
	Lectures           []byte
	Tests              []byte
}

func (q *Queries) GetSectionsWithLecturesAndTestsByCourseID(ctx context.Context, courseID int32) ([]*GetSectionsWithLecturesAndTestsByCourseIDRow, error) {
	rows, err := q.db.Query(ctx, getSectionsWithLecturesAndTestsByCourseID, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetSectionsWithLecturesAndTestsByCourseIDRow
	for rows.Next() {
		var i GetSectionsWithLecturesAndTestsByCourseIDRow
		if err := rows.Scan(
			&i.SectionID,
			&i.SectionTitle,
			&i.SectionDescription,
			&i.Lectures,
			&i.Tests,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserCourses = `-- name: GetUserCourses :many
SELECT c.id, c.title, c.description,
       c.avatar_url, c.subtitle,
       c.rating, c.students_count,
       c.ratings_count,
       c.lectures_length FROM human_resources.courses c
                    INNER JOIN human_resources.courses_attendants uc ON c.id = uc.course_id
WHERE uc.user_id = $1
  AND c.status != 'DRAFT'
  AND c.status != 'PENDING'
ORDER BY
    c.created_at
`

type GetUserCoursesRow struct {
	ID             int32
	Title          string
	Description    string
	AvatarUrl      *string
	Subtitle       *string
	Rating         *float64
	StudentsCount  *int32
	RatingsCount   *int32
	LecturesLength *int32
}

func (q *Queries) GetUserCourses(ctx context.Context, userID int32) ([]*GetUserCoursesRow, error) {
	rows, err := q.db.Query(ctx, getUserCourses, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetUserCoursesRow
	for rows.Next() {
		var i GetUserCoursesRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.AvatarUrl,
			&i.Subtitle,
			&i.Rating,
			&i.StudentsCount,
			&i.RatingsCount,
			&i.LecturesLength,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertLectureTitleAndDescription = `-- name: InsertLectureTitleAndDescription :one
UPDATE human_resources.lectures SET title = $1, description = $2 WHERE id = $3 RETURNING id
`

type InsertLectureTitleAndDescriptionParams struct {
	Title       string
	Description string
	ID          int32
}

func (q *Queries) InsertLectureTitleAndDescription(ctx context.Context, arg *InsertLectureTitleAndDescriptionParams) (int32, error) {
	row := q.db.QueryRow(ctx, insertLectureTitleAndDescription, arg.Title, arg.Description, arg.ID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertLectureVideoUrl = `-- name: InsertLectureVideoUrl :one
UPDATE human_resources.lectures SET video_url = $1 WHERE id = $2 RETURNING id
`

type InsertLectureVideoUrlParams struct {
	VideoUrl string
	ID       int32
}

func (q *Queries) InsertLectureVideoUrl(ctx context.Context, arg *InsertLectureVideoUrlParams) (int32, error) {
	row := q.db.QueryRow(ctx, insertLectureVideoUrl, arg.VideoUrl, arg.ID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const rejectCourse = `-- name: RejectCourse :one
UPDATE human_resources.courses SET status = 'DRAFT', updated_at = now()  WHERE id = $1 RETURNING id
`

func (q *Queries) RejectCourse(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, rejectCourse, id)
	err := row.Scan(&id)
	return id, err
}

const removeCourse = `-- name: RemoveCourse :one
DELETE FROM human_resources.courses WHERE id = $1 RETURNING id
`

func (q *Queries) RemoveCourse(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, removeCourse, id)
	err := row.Scan(&id)
	return id, err
}

const removeLecture = `-- name: RemoveLecture :one
DELETE FROM human_resources.lectures WHERE id = $1 AND section_id = $2 RETURNING id
`

type RemoveLectureParams struct {
	ID        int32
	SectionID int32
}

func (q *Queries) RemoveLecture(ctx context.Context, arg *RemoveLectureParams) (int32, error) {
	row := q.db.QueryRow(ctx, removeLecture, arg.ID, arg.SectionID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const removeSection = `-- name: RemoveSection :one
DELETE FROM human_resources.sections WHERE id = $1 RETURNING id
`

func (q *Queries) RemoveSection(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, removeSection, id)
	err := row.Scan(&id)
	return id, err
}

const searchCoursesByTitle = `-- name: SearchCoursesByTitle :many
SELECT id, title, rating, students_count, avatar_url FROM human_resources.courses
WHERE to_tsvector('russian', title) @@ plainto_tsquery('russian', $1)
  AND status = 'READY'
`

type SearchCoursesByTitleRow struct {
	ID            int32
	Title         string
	Rating        *float64
	StudentsCount *int32
	AvatarUrl     *string
}

func (q *Queries) SearchCoursesByTitle(ctx context.Context, title string) ([]*SearchCoursesByTitleRow, error) {
	rows, err := q.db.Query(ctx, searchCoursesByTitle, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*SearchCoursesByTitleRow
	for rows.Next() {
		var i SearchCoursesByTitleRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Rating,
			&i.StudentsCount,
			&i.AvatarUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchInstructorCoursesByTitle = `-- name: SearchInstructorCoursesByTitle :many
SELECT c.id, c.avatar_url, c.title, c.status FROM human_resources.courses c
WHERE to_tsvector('russian', c.title) @@ plainto_tsquery('russian', $1)
  AND author_id = $2
`

type SearchInstructorCoursesByTitleParams struct {
	Title    string
	AuthorID int32
}

type SearchInstructorCoursesByTitleRow struct {
	ID        int32
	AvatarUrl *string
	Title     string
	Status    HumanResourcesCourseStatuses
}

func (q *Queries) SearchInstructorCoursesByTitle(ctx context.Context, arg *SearchInstructorCoursesByTitleParams) ([]*SearchInstructorCoursesByTitleRow, error) {
	rows, err := q.db.Query(ctx, searchInstructorCoursesByTitle, arg.Title, arg.AuthorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*SearchInstructorCoursesByTitleRow
	for rows.Next() {
		var i SearchInstructorCoursesByTitleRow
		if err := rows.Scan(
			&i.ID,
			&i.AvatarUrl,
			&i.Title,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const sendCourseToCheck = `-- name: SendCourseToCheck :one
UPDATE human_resources.courses SET status = 'PENDING', updated_at = now() WHERE id = $1 RETURNING id
`

func (q *Queries) SendCourseToCheck(ctx context.Context, id int32) (int32, error) {
	row := q.db.QueryRow(ctx, sendCourseToCheck, id)
	err := row.Scan(&id)
	return id, err
}

const updateCourseAvatar = `-- name: UpdateCourseAvatar :one
UPDATE human_resources.courses SET avatar_url = $1 WHERE id = $2 RETURNING avatar_url
`

type UpdateCourseAvatarParams struct {
	AvatarUrl *string
	ID        int32
}

func (q *Queries) UpdateCourseAvatar(ctx context.Context, arg *UpdateCourseAvatarParams) (*string, error) {
	row := q.db.QueryRow(ctx, updateCourseAvatar, arg.AvatarUrl, arg.ID)
	var avatar_url *string
	err := row.Scan(&avatar_url)
	return avatar_url, err
}

const updateCourseGoals = `-- name: UpdateCourseGoals :one
UPDATE human_resources.courses SET course_goals = $1, requirements = $2, target_audience = $3, updated_at = now() WHERE id = $4 RETURNING id
`

type UpdateCourseGoalsParams struct {
	CourseGoals    []string
	Requirements   []string
	TargetAudience []string
	ID             int32
}

func (q *Queries) UpdateCourseGoals(ctx context.Context, arg *UpdateCourseGoalsParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateCourseGoals,
		arg.CourseGoals,
		arg.Requirements,
		arg.TargetAudience,
		arg.ID,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateCoursePreviewVideo = `-- name: UpdateCoursePreviewVideo :one
UPDATE human_resources.courses SET preview_video_url = $1, updated_at = now() WHERE id = $2 RETURNING preview_video_url
`

type UpdateCoursePreviewVideoParams struct {
	PreviewVideoUrl *string
	ID              int32
}

func (q *Queries) UpdateCoursePreviewVideo(ctx context.Context, arg *UpdateCoursePreviewVideoParams) (*string, error) {
	row := q.db.QueryRow(ctx, updateCoursePreviewVideo, arg.PreviewVideoUrl, arg.ID)
	var preview_video_url *string
	err := row.Scan(&preview_video_url)
	return preview_video_url, err
}

const updateLectureTitle = `-- name: UpdateLectureTitle :one
UPDATE human_resources.lectures SET title = $1 WHERE id = $2 RETURNING id
`

type UpdateLectureTitleParams struct {
	Title string
	ID    int32
}

func (q *Queries) UpdateLectureTitle(ctx context.Context, arg *UpdateLectureTitleParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateLectureTitle, arg.Title, arg.ID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateSectionTitle = `-- name: UpdateSectionTitle :one
UPDATE human_resources.sections SET title = $1 WHERE id = $2 RETURNING id
`

type UpdateSectionTitleParams struct {
	Title string
	ID    int32
}

func (q *Queries) UpdateSectionTitle(ctx context.Context, arg *UpdateSectionTitleParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateSectionTitle, arg.Title, arg.ID)
	var id int32
	err := row.Scan(&id)
	return id, err
}