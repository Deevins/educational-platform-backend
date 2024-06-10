// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package directories

import (
	"context"
)

type Querier interface {
	GetCategoriesAndSubcategories(ctx context.Context) ([]*GetCategoriesAndSubcategoriesRow, error)
	GetLanguages(ctx context.Context) ([]*GetLanguagesRow, error)
	GetMetasCount(ctx context.Context) (*GetMetasCountRow, error)
}

var _ Querier = (*Queries)(nil)
