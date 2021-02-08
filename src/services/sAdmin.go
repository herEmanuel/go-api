package services

import (
	"github.com/herEmanuel/go-api/models"
	"github.com/herEmanuel/go-api/repositories"
)

func CreateProduct(bodyFields map[string]interface{}) (models.Product, error) {

	var product models.Product

	err := repositories.CreateProduct(bodyFields, &product)
	if err != nil {
		return product, err
	}

	return product, nil
}
