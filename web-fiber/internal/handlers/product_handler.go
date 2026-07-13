package handlers

import (
	"web-fiber/internal/models"
	"web-fiber/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	products := usecases.GetProductsLogic()

	return c.JSON(fiber.Map{"message": "Lấy danh sách sản phẩm thành công!", "data": products})
}

func CreateProduct(c *fiber.Ctx) error {
	var p models.Product

	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Dữ liệu gửi lên không hợp lệ"})
	}

	if err := usecases.CreateProductLogic(&p); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Không thể lưu vào DB!"})
	}

	return c.JSON(fiber.Map{"message": "Thêm sản phẩm thành công!", "data": p})
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID sẽ là số nguyên hợp lệ"})
	}

	product, err := usecases.GetProductByIdLogic(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy sản phẩm này trong DB!"})
	}

	return c.JSON(fiber.Map{"message": "Tìm thấy sản phẩm",  "data": product,})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID phải là số nguyên!"})
	}

	err = usecases.DeleteProduct(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Không thể xóa sản phẩm này!"})
	}

	return c.JSON(fiber.Map{"message": "Đã xóa sản phẩm thành công"})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID phải là số nguyên"})
	}

	var p models.Product

	if 	err := usecases.UpdateProduct(id, &p); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Không tìm thấy hoặc không cập nhật được sản phẩm"})
	}

	return c.JSON(fiber.Map{"message": "Cập nhật sản phẩm thành công"})
}