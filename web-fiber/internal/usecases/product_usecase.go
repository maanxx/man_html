package usecases

import (
	"web-fiber/internal/models"
	"web-fiber/internal/repositories"
)

func GetProductsLogic() []models.Product {
	products := repositories.GetProducts()
	return products
}

func CreateProductLogic(p *models.Product) error {
	return repositories.CreateProduct(p)
}

func GetProductByIdLogic(id int) (*models.Product, error) {
	return repositories.GetProductById(id)
}

func DeleteProduct(id int) error {
	return repositories.DeleteProduct(id)
}

func UpdateProduct(id int, updateData *models.Product) error {
	return repositories.UpdateProduct(id, updateData)
}


