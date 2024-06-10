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

-- name: UpdateCourseGoals :one
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
SELECT c.id, c.title, c.description,
       c.avatar_url, c.subtitle,
       c.rating, c.students_count,
       c.ratings_count,
       c.lectures_length FROM human_resources.courses c
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

-- name: CreateSection :one
INSERT INTO human_resources.sections (title, description, course_id, serial_number) VALUES (@title, @description, @course_id, @serial_number) RETURNING id;

-- name: RemoveSection :one
DELETE FROM human_resources.sections WHERE id = @id RETURNING id;

-- name: UpdateSectionTitle :one
UPDATE human_resources.sections SET title = @title WHERE id = @id RETURNING id;

-- name: CreateLecture :one
INSERT INTO human_resources.lectures (title, description, section_id, serial_number) VALUES (@title, @description, @section_id, @serial_number) RETURNING id;

-- name: CreateTest :one
INSERT INTO human_resources.tests (name,description, section_id, serial_number) VALUES (@name,@description, @section_id, @serial_number) RETURNING id;

-- name: UpdateLectureVideoUrl :one
UPDATE human_resources.lectures SET video_url = @video_url WHERE id = @id RETURNING id;

-- name: UpdateLectureTitle :one
UPDATE human_resources.lectures SET title = @title WHERE section_id = @section_id AND id = @id RETURNING id;

-- name: UpdateTestTitle :one
UPDATE human_resources.tests SET name = @name WHERE section_id = @section_id AND id = @id RETURNING id;

-- name: RemoveLecture :one
DELETE FROM human_resources.lectures WHERE id = @id AND section_id = @section_id RETURNING id;

-- name: RemoveTest :one
DELETE FROM human_resources.tests WHERE id = @id AND section_id = @section_id RETURNING id;

-- name: GetInstructorCourses :many
SELECT c.id, c.avatar_url, c.title, c.status FROM human_resources.courses c WHERE author_id = @author_id ORDER BY created_at DESC;

-- name: SearchInstructorCoursesByTitle :many
SELECT c.id, c.avatar_url, c.title, c.status FROM human_resources.courses c
WHERE to_tsvector('russian', c.title) @@ plainto_tsquery('russian', @title)
  AND author_id = @author_id;

-- name: RemoveCourse :one
DELETE FROM human_resources.courses WHERE id = @id RETURNING id;

-- name: GetFullCourseInfoWithInstructorByCourseID :one
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
WHERE c.id = @id;


-- name: GetSectionsWithLecturesAndTestsByCourseID :many
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
        q.body AS question_body,
        a.id AS answer_id,
        a.body AS answer_body,
        a.is_correct AS answer_is_correct,
        a.description AS answer_description
    FROM
        human_resources.sections s
            LEFT JOIN
        human_resources.lectures l ON s.id = l.section_id
            LEFT JOIN
        human_resources.tests t ON s.id = t.section_id
            LEFT JOIN
        human_resources.tests_questions q ON t.id = q.test_id
            LEFT JOIN
        human_resources.tests_questions_answers a ON q.id = a.question_id
    WHERE
        s.course_id = @course_id
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
                                    'question_body', question_body,
                                    'answers', json_agg(
                                            json_build_object(
                                                    'answer_id', answer_id,
                                                    'answer_body', answer_body,
                                                    'answer_is_correct', answer_is_correct,
                                                    'answer_description', answer_description
                                            )
                                               )
                            )
                                 )
            )
    ) AS tests
FROM
    recursive_section
GROUP BY
    section_id, section_title, section_description;

-- name: GetCoursesAvatarsByIDs :many
SELECT id, avatar_url FROM human_resources.courses WHERE id = ANY($1::int[]);

-- name: GetCourseAvatarByID :one
SELECT avatar_url FROM human_resources.courses WHERE id = @id;

-- name: GetCoursePreviewVideoByID :one
SELECT preview_video_url FROM human_resources.courses WHERE id = @id;

-- name: CreateQuestion :one
INSERT INTO human_resources.tests_questions (test_id, body) VALUES (@test_id, @body) RETURNING id;

-- name: CreateAnswer :one
INSERT INTO human_resources.tests_questions_answers (question_id, body, description, is_correct) VALUES (@question_id, @body, @description, @is_correct) RETURNING id;
