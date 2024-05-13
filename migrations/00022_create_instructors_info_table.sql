-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE human_resources.instructors_info (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    has_previous_experience TEXT NOT NULL default '',
    has_video_knowledge text not null default '',
    current_audience_count TEXT NOT NULL default '',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);


-- +goose Down

DROP TABLE IF EXISTS human_resources.instructors_info;
