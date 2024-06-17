-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE human_resources.instructors_info (
    user_id INT PRIMARY KEY,
    previous_experience TEXT NOT NULL default '',
    video_knowledge TEXT not null default '',
    current_audience TEXT NOT NULL default '',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);


-- +goose Down

DROP TABLE IF EXISTS human_resources.instructors_info;
