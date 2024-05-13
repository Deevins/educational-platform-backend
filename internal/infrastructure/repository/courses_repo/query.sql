-- name: CreateCourseBase :exec
INSERT INTO human_resources.courses (title,type, category_id, time_planned, description, created_at, updated_at) VALUES ($1, $2, $3, $4,$5, now(), now()) RETURNING id;

-- name: AddCourseGoals :exec
UPDATE human_resources.courses SET course_goals = $1, requirements = $2, target_audience = $3, updated_at = now() WHERE id = $4;

-- name: AddCourseBasicInfo :exec
UPDATE human_resources.courses SET title = $1, subtitle = $2,  description = $3, language = $4,level = $5, category_id = $6, subcategory_id = $7, updated_at = now() WHERE id = $8;
