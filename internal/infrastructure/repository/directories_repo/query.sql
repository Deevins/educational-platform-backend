-- name: GetCategoriesAndSubcategories :many
 SELECT c.id as category_id, c.name as category_name, s.id as subcategory_id, s.name as subcategory_name FROM human_resources.categories c
    JOIN human_resources.subcategories s ON c.id = s.category_id
 group by c.id, s.id;

-- name: GetLanguages :many
SELECT l.name from human_resources.languages l;