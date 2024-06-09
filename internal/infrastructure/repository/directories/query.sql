-- name: GetCategoriesAndSubcategories :many
SELECT
    c.id AS category_id,
    c.name AS category_name,
    s.id AS subcategory_id,
    s.name AS subcategory_name
FROM
    human_resources.categories c
        LEFT JOIN
    human_resources.subcategories s ON c.id = s.category_id
ORDER BY
    c.id, s.id;

-- name: GetLanguages :many
SELECT l.id, l.name from human_resources.languages l;