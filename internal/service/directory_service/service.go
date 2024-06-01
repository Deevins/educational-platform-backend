package directory_service

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/handler"
	"github.com/deevins/educational-platform-backend/internal/infrastructure/repository/directories_repo"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/pkg/errors"
)

//DirectoryService

var _ handler.DirectoryService = &Service{}

func NewService(repo directories_repo.Querier) *Service {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo directories_repo.Querier
}

func (s *Service) GetLanguages(ctx context.Context) ([]*model.Language, error) {
	languages, err := s.repo.GetLanguages(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get languages")
	}
	output := make([]*model.Language, 0)
	for _, language := range languages {
		output = append(output, &model.Language{
			ID:   language.ID,
			Name: language.Name,
		})
	}

	return output, nil
}

func repackDBToModel(dbCategories []*directories_repo.GetCategoriesAndSubcategoriesRow) []*model.Category {
	categoryMap := make(map[int32]*model.Category)
	var categories []*model.Category

	for _, dbCategory := range dbCategories {
		// Если категория уже существует в карте, берем её, иначе создаем новую
		category, exists := categoryMap[dbCategory.CategoryID]
		if !exists {
			category = &model.Category{
				ID:            dbCategory.CategoryID,
				Name:          dbCategory.CategoryName,
				Subcategories: []*model.Category{},
			}
			categoryMap[dbCategory.CategoryID] = category
			categories = append(categories, category)
		}

		// Если есть подкатегория, добавляем её
		if dbCategory.SubcategoryID != nil {
			subcategory := &model.Category{
				ID:   *dbCategory.SubcategoryID,
				Name: *dbCategory.SubcategoryName,
			}
			category.Subcategories = append(category.Subcategories, subcategory)
		}
	}

	return categories
}
func (s *Service) GetCategoriesWithSubCategories(ctx context.Context) ([]*model.Category, error) {
	categoriesWithSubcategories, err := s.repo.GetCategoriesAndSubcategories(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get categories with subcategories")
	}
	return repackDBToModel(categoriesWithSubcategories), nil
}
