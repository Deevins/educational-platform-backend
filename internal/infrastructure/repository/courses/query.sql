-- name: CreateCourseBase :one
INSERT INTO human_resources.courses (
                title,
                type,
                author_id,
                category_title,
                time_planned)
VALUES (@title,
        @type,
        @author_id,
        @category_title,
        @time_planned
       ) RETURNING id;

-- name: UpdateCourseGoals :one
UPDATE human_resources.courses SET course_goals = @course_goals, requirements = @requirements, target_audience = @target_audience, updated_at = now() WHERE id = @id RETURNING id;

-- name: GetCourseGoals :one
SELECT course_goals, requirements, target_audience FROM human_resources.courses WHERE id = @id;

-- name: AddCourseBasicInfo :one
UPDATE human_resources.courses SET
        title = @title,
        subtitle = @subtitle,
        description = @description,
        language = @language,
        level = @level,
        category_title = @category_title,
        updated_at = now()
                        WHERE id = @id RETURNING id;

-- name: GetCourseBasicInfo :one
SELECT c.title, c.subtitle, c.description, c.language, c.level, c.category_title FROM human_resources.courses c WHERE id = @id;

-- name: GetUserCourses :many
SELECT c.id, c.title, c.description,
       c.avatar_url, c.subtitle,
       c.rating, c.students_count,
       c.ratings_count,
       c.lectures_length_interval FROM human_resources.courses c
                    INNER JOIN human_resources.courses_attendants uc ON c.id = uc.course_id
WHERE uc.user_id = @user_id
  AND c.status != 'DRAFT'
  AND c.status != 'PENDING'
ORDER BY
    c.created_at;


-- name: GetAllReadyCourses :many
SELECT
    c.id AS course_id,
    c.title,
    c.description,
    c.level,
    c.lectures_count,
    c.preview_video_url AS course_preview_video_url,
    c.avatar_url AS course_avatar_url,
    c.subtitle,
    c.rating,
    c.students_count,
    c.ratings_count,
    c.lectures_length_interval,
    u.full_name AS instructor_name,
    u.avatar_url AS instructor_avatar_url,
    c.created_at -- Добавляем created_at в селекцию и группировку
FROM
    human_resources.courses c
        JOIN
    human_resources.users u ON c.author_id = u.id
WHERE
    c.status != 'PENDING'
  AND c.status != 'DRAFT'
GROUP BY
    c.id,
    c.title,
    c.description,
    c.level,
    c.preview_video_url,
    c.avatar_url,
    c.subtitle,
    c.rating,
    c.lectures_count,
    c.students_count,
    c.ratings_count,
    c.lectures_length_interval,
    u.full_name,
    u.avatar_url,
    c.created_at -- Указали created_at в секции GROUP BY
ORDER BY
    c.created_at;


-- name: GetAllPendingCourses :many
SELECT
    c.id AS course_id,
    c.title,
    c.description,
    c.level,
    c.preview_video_url AS course_preview_video_url,
    c.avatar_url AS course_avatar_url,
    c.subtitle,
    c.lectures_count,
    c.rating,
    c.students_count,
    c.ratings_count,
    c.lectures_length_interval,
    u.full_name AS instructor_name,
    u.avatar_url AS instructor_avatar_url,
    c.created_at -- Добавляем created_at в селекцию и группировку
FROM
    human_resources.courses c
        JOIN
    human_resources.users u ON c.author_id = u.id
WHERE
    c.status != 'DRAFT'
  AND c.status != 'READY'
GROUP BY
    c.id,
    c.title,
    c.description,
    c.lectures_count,
    c.level,
    c.preview_video_url,
    c.avatar_url,
    c.subtitle,
    c.rating,
    c.students_count,
    c.ratings_count,
    c.lectures_length_interval,
    u.full_name,
    u.avatar_url,
    c.created_at -- Указали created_at в секции GROUP BY
ORDER BY
    c.created_at;

-- name: GetAllDraftCourses :many
SELECT
    c.id AS course_id,
    c.title,
    c.description,
    c.lectures_count,
    c.level,
    c.preview_video_url AS course_preview_video_url,
    c.avatar_url AS course_avatar_url,
    c.subtitle,
    c.rating,
    c.students_count,
    c.ratings_count,
    c.lectures_length_interval,
    u.full_name AS instructor_name,
    u.avatar_url AS instructor_avatar_url,
    c.created_at -- Добавляем created_at в селекцию и группировку
FROM
    human_resources.courses c
        JOIN
    human_resources.users u ON c.author_id = u.id
WHERE
    c.status != 'PENDING'
  AND c.status != 'READY'
GROUP BY
    c.id,
    c.title,
    c.description,
    c.level,
    c.lectures_count,
    c.preview_video_url,
    c.avatar_url,
    c.subtitle,
    c.rating,
    c.students_count,
    c.ratings_count,
    c.lectures_length_interval,
    u.full_name,
    u.avatar_url,
    c.created_at -- Указали created_at в секции GROUP BY
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

-- name: UpdateLecturesInfo :one
UPDATE human_resources.courses
SET lectures_count = lectures_count + @lectures_count,
    lectures_length_interval = lectures_length_interval + @lectures_length_interval
WHERE id = @id RETURNING id;

-- name: UpdateLectureVideoAddedInfo :one
UPDATE human_resources.lectures SET lecture_video_length = @lecture_video_length WHERE id = @id RETURNING id;

-- name: GetLectureByID :one
SELECT id, title, description, video_url, lecture_video_length FROM human_resources.lectures WHERE id = @id;

-- name: UpdateLectureTitle :one
UPDATE human_resources.lectures SET title = @title WHERE section_id = @section_id AND id = @id RETURNING id;

-- name: UpdateTestTitle :one
UPDATE human_resources.tests SET name = @name WHERE section_id = @section_id AND id = @id RETURNING id;

-- name: RemoveLecture :one
DELETE FROM human_resources.lectures WHERE id = @id AND section_id = @section_id RETURNING id;

-- name: RemoveTest :one
DELETE FROM human_resources.tests WHERE id = @id AND section_id = @section_id RETURNING id;

-- name: GetInstructorCourses :many
SELECT c.id, c.avatar_url, c.title, c.ratings_count, c.rating FROM human_resources.courses c WHERE author_id = @author_id ORDER BY created_at DESC;

-- name: SearchInstructorCoursesByTitle :many
SELECT c.id, c.avatar_url, c.title, c.status, c.title, c.ratings_count, c.rating FROM human_resources.courses c
WHERE to_tsvector('russian', c.title) @@ plainto_tsquery('russian', @title)
  AND author_id = @author_id;

-- name: RemoveCourse :one
DELETE FROM human_resources.courses WHERE id = @id RETURNING id;

-- name: GetFullCourseInfoWithInstructorByCourseID :one
SELECT
    c.id AS course_id,
    c.title,
    c.subtitle,
    c.description,
    c.language,
    c.level,
    u.description AS instructor_description,
    c.rating,
    c.students_count,
    c.ratings_count,
    c.lectures_count,
    c.lectures_length_interval,
    c.avatar_url AS course_avatar_url,
    c.preview_video_url,
    c.status AS course_status,
    c.created_at AS course_created_at,
    c.course_goals,
    c.requirements,
    c.target_audience,
    c.author_id,
    c.category_title AS category_title,
    u.full_name AS instructor_full_name,
    (SELECT COUNT(*) FROM human_resources.courses WHERE author_id = u.id) AS instructor_courses_count,
    (SELECT SUM(students_count) FROM human_resources.courses WHERE author_id = u.id) AS instructor_students_count,
    (SELECT COUNT(*) FROM human_resources.courses_reviews cr
     JOIN human_resources.courses cc ON cr.course_id = cc.id
     WHERE cc.author_id = u.id) AS total_reviews_count,
    u.instructor_rating,
    u.avatar_url AS instructor_avatar_url,
    u.id AS instructor_id
FROM
    human_resources.courses c
    JOIN human_resources.users u ON c.author_id = u.id
WHERE
    c.id = @id;

-- name: GetCourseReviewsByCourseID :many
SELECT
    cr.rating AS review_rating,
    cr.review AS review_text,
    cr.created_at AS review_created_at,
    ru.full_name AS reviewer_full_name,
    ru.avatar_url AS reviewer_avatar_url
FROM
    human_resources.courses_reviews cr
        JOIN
    human_resources.users ru ON cr.user_id = ru.id
WHERE
    cr.course_id = @id
    ORDER by cr.created_at ASC;


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

-- name: UpdateQuestion :one
UPDATE human_resources.tests_questions SET body = @body WHERE id = @id RETURNING id;

-- name: UpdateAnswer :one
UPDATE human_resources.tests_questions_answers SET body = @body, description = @description, is_correct = @is_correct WHERE id = @id RETURNING id;

-- name: RemoveQuestion :one
DELETE FROM human_resources.tests_questions WHERE id = @id RETURNING id;

-- name: RemoveAnswer :one
DELETE FROM human_resources.tests_questions_answers WHERE id = @id RETURNING id;

-- name: GetFullCourseByID :one
SELECT
    c.id AS course_id,
    c.title,
    c.subtitle,
    c.description,
    c.language,
    c.level,
    c.rating,
    c.students_count,
    c.ratings_count,
    c.lectures_count,
    c.lectures_length_interval,
    c.avatar_url AS course_avatar_url,
    c.preview_video_url,
    c.status AS course_status,
    c.created_at AS course_created_at,
    c.course_goals,
    c.requirements,
    c.target_audience,
    c.author_id,
    c.category_title AS category_title,
    u.full_name AS instructor_full_name,
    u.avatar_url AS instructor_avatar_url,
    u.id AS instructor_id
FROM
    human_resources.courses c
    JOIN human_resources.users u ON c.author_id = u.id
WHERE
    c.id = @id;