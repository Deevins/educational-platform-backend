-- name: GetUsers :many
SELECT * from human_resources.users;

-- name: GetUserByID :one
SELECT * from human_resources.users WHERE id = @id;

-- name: CreateUser :one
INSERT INTO human_resources.users (full_name, email,description, password_hashed, phone_number) VALUES (@full_name, @email, @description, @password_hashed, @phone_number) RETURNING id;

-- name: GetUserByEmailAndHashedPassword :one
SELECT * from human_resources.users WHERE email = @email AND password_hashed = @password_hashed LIMIT 1;

-- name: UpdateHasUserTriedInstructor :one
UPDATE human_resources.users SET has_user_tried_instructor = true WHERE id = @id RETURNING id;

-- name: GetHasUserTriedInstructor :one
SELECT has_user_tried_instructor from human_resources.users WHERE id = @id;

-- name: UpdateUserFullName :one
UPDATE human_resources.users SET full_name = @full_name WHERE id = @id RETURNING id;

-- name: UpdateUserEmail :one
UPDATE human_resources.users SET email = @email WHERE id = @id RETURNING id;

-- name: UpdateUserDescription :one
UPDATE human_resources.users SET description = @description WHERE id = @id RETURNING id;

-- name: UpdateUserPhone :one
UPDATE human_resources.users SET phone_number = @phone_number WHERE id = @id RETURNING id;

-- name: AddTeachingExperience :one
INSERT INTO human_resources.instructors_info (user_id, video_knowledge, current_audience, previous_experience) VALUES (@user_id, @video_knowledge, @current_audience, @previous_experience) RETURNING user_id;

-- name: UpdateAvatar :one
UPDATE human_resources.users SET avatar_url = @avatar_url WHERE id = @id RETURNING id;

-- name: RegisterToCourse :one
INSERT INTO human_resources.courses_attendants (user_id, course_id) VALUES (@user_id, @course_id) ON CONFLICT do nothing RETURNING user_id;

-- name: CheckIfUserRegisteredToCourse :one
SELECT * from human_resources.courses_attendants WHERE user_id = @user_id AND course_id = @course_id;

-- name: UpdateCourseStudentsCount :one
UPDATE human_resources.courses SET students_count = students_count + 1 WHERE id = @id RETURNING id;