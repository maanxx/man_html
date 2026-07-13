package repositories

import (
	"web-fiber/internal/config"
	"web-fiber/internal/models"
)

func GetProducts() []models.Product{
	var products[]models.Product
	config.DB.Find(&products)
	return products
}

func CreateProduct(p *models.Product) error {
	result := config.DB.Create(p)
	return result.Error
}

func GetProductById(id int) (*models.Product, error) {
	var p models.Product
	result := config.DB.First(&p, id)
	return &p, result.Error
}

func DeleteProduct(id int) error {
	result := config.DB.Delete(&models.Product{}, id)
	return result.Error
}

func UpdateProduct(id int, updateData *models.Product) error {
	var p models.Product

	if err := config.DB.First(&p, id).Error; err != nil {
		return err
	}

	p.Name = updateData.Name
	p.Price = updateData.Price

	result := config.DB.Save(&p)
	
	return result.Error
} 