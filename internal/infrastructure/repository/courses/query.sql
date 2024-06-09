-- name: CreateCourseBase :one
INSERT INTO human_resources.courses (
                title,
                type,
                author_id,
                category_id,
                time_planned)
VALUES (@title,
        @type,
        @author_id,
        @category_id,
        @time_planned
       ) RETURNING id;

-- name: AddCourseGoals :one
UPDATE human_resources.courses SET course_goals = @course_goals, requirements = @requirements, target_audience = @target_audience, updated_at = now() WHERE id = @id RETURNING id;

-- name: AddCourseBasicInfo :one
UPDATE human_resources.courses SET
        title = @title,
        subtitle = @subtitle,
        description = @description,
        language = @language,
        level = @level,
        category_id = @category_id,
        updated_at = now()
                        WHERE id = @id RETURNING id;

-- name: GetUserCourses :many
SELECT c.id, c.title, c.description, c.avatar_url, c.subtitle, c.rating, c.students_count, c.ratings_count, c.lectures_length FROM human_resources.courses c
                    INNER JOIN human_resources.courses_attendants uc ON c.id = uc.course_id
WHERE uc.user_id = @user_id
  AND c.status != 'DRAFT'
  AND c.status != 'PENDING'
ORDER BY
    c.created_at;

-- name: GetAllReadyCourses :many
SELECT * FROM human_resources.courses c
WHERE c.status != 'DRAFT'
  AND c.status != 'PENDING'
ORDER BY
    c.created_at;

-- name: GetAllPendingCourses :many
SELECT *  FROM human_resources.courses c
WHERE c.status != 'DRAFT'
  AND c.status != 'READY'
ORDER BY
    c.created_at;

-- name: GetAllDraftCourses :many
SELECT * FROM human_resources.courses c
WHERE c.status != 'PENDING'
  AND c.status != 'READY'
ORDER BY
    c.created_at;

-- name: SendCourseToCheck :one
UPDATE human_resources.courses SET status = 'PENDING', updated_at = now() WHERE id = @id RETURNING id;

-- name: RejectCourse :one
UPDATE human_resources.courses SET status = 'DRAFT', updated_at = now()  WHERE id = @id RETURNING id;

-- name: ApproveCourse :one
UPDATE human_resources.courses SET status = 'READY', updated_at = now()  WHERE id = @id RETURNING id;


-- name: SearchCoursesByTitle :many
SELECT id, title, rating, students_count, avatar_url FROM human_resources.courses
WHERE to_tsvector('russian', title) @@ plainto_tsquery('russian', @title)
  AND status = 'READY';


-- name: UpdateCourseAvatar :one
UPDATE human_resources.courses SET avatar_url = @avatar_url WHERE id = @id RETURNING avatar_url;

-- name: UpdateCoursePreviewVideo :one
UPDATE human_resources.courses SET preview_video_url = @preview_video_url, updated_at = now() WHERE id = @id RETURNING preview_video_url;


-- name: InsertLectureAndCourseLecture :one
-- Вставляет лекцию в таблицу lectures и соответствующую запись в таблицу courses_lectures
WITH inserted_lecture AS (
    INSERT INTO human_resources.lectures (video_url)
        VALUES (@video_url)
        RETURNING id
)
INSERT INTO human_resources.courses_lectures (course_id, lecture_id)
VALUES (@course_id, (SELECT id FROM inserted_lecture))
RETURNING id, course_id, lecture_id;


-- name: InsertLectureTitleAndDescription :one
UPDATE human_resources.lectures SET title = @title, description = @description WHERE id = @id RETURNING id;

-- name: UpdateLectureTitle :one
UPDATE human_resources.lectures SET title = @title WHERE id = @id RETURNING id;

-- name: RemoveLecture :one
DELETE FROM human_resources.lectures WHERE id = @id RETURNING id;

-- name: RemoveLectureFromCourse :one
DELETE FROM human_resources.courses_lectures WHERE id = @id RETURNING id;