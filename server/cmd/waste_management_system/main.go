package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/akers1023/Smart-Waste-Management-System/internal/app/handlers"
	"github.com/akers1023/Smart-Waste-Management-System/internal/app/models"
	"github.com/akers1023/Smart-Waste-Management-System/internal/app/repository"
	"github.com/akers1023/Smart-Waste-Management-System/internal/app/service"
	"github.com/akers1023/Smart-Waste-Management-System/internal/connections"
	"github.com/akers1023/Smart-Waste-Management-System/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("DB_PORT must be a valid integer")
	}
	sql := &connections.Sql{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Password: os.Getenv("DB_PASS"),
		UserName: os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("DB_NAME"),
	}

	db, _ := sql.Connect()
	defer sql.Close()

	err = models.MigrateUser(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}
	fmt.Println("ok con de")

	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo) // Điền thêm thông tin cần thiết

	// Tạo đối tượng UserHandler và truyền vào UserService
	userHandler := handlers.NewUserHandler(userService) //ervice) //

	// Gọi hàm UserRoutes với đối tượng UserHandler vừa tạo
	app := fiber.New()
	routes.UserRoutes(app, userHandler)
	// routes.UserRoutes(app, repo)
	log.Fatal(app.Listen(":1234"))
}
