-- +goose NO TRANSACTION
-- +goose Up

CREATE TABLE human_resources.last_course_items_visited_by_users (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    user_id INT NOT NULL,
    last_item_id INT NOT NULL, -- here will be stored the last item visited by the user(lecture/test)
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);


-- +goose Down

DROP TABLE IF EXISTS human_resources.last_course_items_visited_by_users;
