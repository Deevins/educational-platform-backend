// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package directories

import (
	"context"
)

const getCategoriesAndSubcategories = `-- name: GetCategoriesAndSubcategories :many
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
    c.id, s.id
`

type GetCategoriesAndSubcategoriesRow struct {
	CategoryID      int32
	CategoryName    string
	SubcategoryID   *int32
	SubcategoryName *string
}

func (q *Queries) GetCategoriesAndSubcategories(ctx context.Context) ([]*GetCategoriesAndSubcategoriesRow, error) {
	rows, err := q.db.Query(ctx, getCategoriesAndSubcategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCategoriesAndSubcategoriesRow
	for rows.Next() {
		var i GetCategoriesAndSubcategoriesRow
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.SubcategoryID,
			&i.SubcategoryName,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLanguages = `-- name: GetLanguages :many
SELECT l.id, l.name from human_resources.languages l
`

type GetLanguagesRow struct {
	ID   int32
	Name string
}

func (q *Queries) GetLanguages(ctx context.Context) ([]*GetLanguagesRow, error) {
	rows, err := q.db.Query(ctx, getLanguages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetLanguagesRow
	for rows.Next() {
		var i GetLanguagesRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLevels = `-- name: GetLevels :many
SELECT l.id, l.name from human_resources.skill_levels l
`

type GetLevelsRow struct {
	ID   int32
	Name string
}

func (q *Queries) GetLevels(ctx context.Context) ([]*GetLevelsRow, error) {
	rows, err := q.db.Query(ctx, getLevels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetLevelsRow
	for rows.Next() {
		var i GetLevelsRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMetasCount = `-- name: GetMetasCount :one
SELECT
    (SELECT COUNT(*) FROM human_resources.users) AS usersCount,
    (SELECT COUNT(*) FROM human_resources.users WHERE has_user_tried_instructor = false) AS studentsCount,
    (SELECT COUNT(*) FROM human_resources.courses) AS coursesCount
`

type GetMetasCountRow struct {
	Userscount    int64
	Studentscount int64
	Coursescount  int64
}

func (q *Queries) GetMetasCount(ctx context.Context) (*GetMetasCountRow, error) {
	row := q.db.QueryRow(ctx, getMetasCount)
	var i GetMetasCountRow
	err := row.Scan(&i.Userscount, &i.Studentscount, &i.Coursescount)
	return &i, err
}
