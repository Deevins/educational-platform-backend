-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE human_resources.courses (
    id SERIAL PRIMARY KEY,
    author_id INT NOT NULL, -- references to user who create course
    title VARCHAR(255) NOT NULL,
    subtitle VARCHAR(255) default '',
    description TEXT NOT NULL,
    avatar_url VARCHAR(255) default '',
    students_count INT DEFAULT 0,
    ratings_count INT DEFAULT 0,
    rating FLOAT DEFAULT 0,
    category_id int default 0,
    subcategory_id int default 0,
    language VARCHAR(255) default 'Русский',
    level VARCHAR(255) default 'Начинающий',
    time_planned TEXT DEFAULT '0',
    course_goals TEXT[] DEFAULT '{}',
    requirements TEXT[] DEFAULT '{}',
    target_audience TEXT[] DEFAULT '{}',
    type human_resources.course_types DEFAULT 'course',
    status human_resources.course_statuses DEFAULT 'DRAFT',
    lectures_length INT DEFAULT 0,
    lectures_count INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);


-- +goose Down

DROP TABLE IF EXISTS human_resources.courses;
