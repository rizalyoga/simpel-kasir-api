package services

import (
	"kasir-api-bootcamp/common/errors"
	"kasir-api-bootcamp/models"
	"kasir-api-bootcamp/repositories"
)

func GetCategories() ([]models.Category, error) {
	return repositories.Categories, nil
}

func GetCategory(id int) (*models.Category, error) {
	for _, c := range repositories.Categories {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, &errors.ErrNotFound{Resource: "Category", ID: id}
}

func CreateCategory(newCategory models.Category) (*models.Category, error) {
	newCategory.ID = len(repositories.Categories) + 1
	repositories.Categories = append(repositories.Categories, newCategory)
	return &newCategory, nil
}

func UpdateCategory(id int, updatedCategory models.Category) (*models.Category, error) {
	for i, c := range repositories.Categories {
		if c.ID == id {
			updatedCategory.ID = id
			repositories.Categories[i] = updatedCategory
			return &updatedCategory, nil
		}
	}
	return nil, &errors.ErrNotFound{Resource: "Category", ID: id}
}

func DeleteCategory(id int) error {
	for i, c := range repositories.Categories {
		if c.ID == id {
			repositories.Categories = append(repositories.Categories[:i], repositories.Categories[i+1:]...)
			return nil
		}
	}
	return &errors.ErrNotFound{Resource: "Category", ID: id}
}
