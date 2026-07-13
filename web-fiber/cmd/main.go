package main

import (
	"fmt"
	"log"
	"web-fiber/internal/config"
	"web-fiber/internal/models"
	"web-fiber/internal/routes"

	"github.com/gofiber/fiber/v2"
)


func main(){
	config.ConnectDB()

	config.DB.AutoMigrate(&models.Product{})

	fmt.Println("Tạo bảng thành công!")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Web Fiber đã sẵn dàng chiến đấu!")
	}) 

	routes.SetupProductRoutes(app)
	
	log.Println("Đang chạy server ở http://localhost:5000")
	log.Fatal(app.Listen(":5000"))
}