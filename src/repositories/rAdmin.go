package repositories

import (
	"errors"

	"github.com/herEmanuel/go-api/database"
	"github.com/herEmanuel/go-api/models"
)

func CreateProduct(bodyFields map[string]interface{}, varToStore *models.Product) error {

	var categoriesExist []models.Category

	if len(bodyFields["categories"].([]string)) != 0 {

		database.Db.Model(&models.Category{}).
			Where("name IN ?", bodyFields["categories"].([]string)).
			Find(&categoriesExist)

		if len(categoriesExist) != len(bodyFields["categories"].([]string)) {
			return errors.New("Insert valid categories")
		}

	}

	newProduct := models.Product{
		Name:         bodyFields["name"].(string),
		Stock:        bodyFields["totalStock"].(int),
		Price:        bodyFields["productPrice"].(float32),
		Rating:       0.0,
		RatingAmount: 0,
		Categories:   categoriesExist,
	}

	result := database.Db.Create(&newProduct)
	if result.Error != nil {
		return errors.New("Could not save the new product")
	}

	if len(categoriesExist) != 0 {
		for _, category := range categoriesExist {
			category.ProductsAmount++

			result = database.Db.Save(&category)
			if result.Error != nil {
				return errors.New("Could not save the categories")
			}
		}
	}

	*varToStore = newProduct

	return nil
}
