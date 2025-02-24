package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)


type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	MobileNo string `json:"mobileNo" validate:"required"`
	Product  string `json:"product" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}


func connectToDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/user?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTable() {
	db, err := connectToDB()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS user_orders (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		mobile_no VARCHAR(20) NOT NULL,
		product VARCHAR(50) NOT NULL,
		quantity INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating table:", err)
	} else {
		log.Println("Table 'user_orders' is ready!")
	}
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	createTable()

	app.Post("/api/register", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		if user.Name == "" || user.Email == "" || user.MobileNo == "" || user.Product == "" || user.Quantity == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
		}

		db, err := connectToDB()
		if err != nil {
			log.Println("Database connection error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
		}
		defer db.Close()

		query := `INSERT INTO user_orders (name, email, mobile_no, product, quantity) 
				  VALUES (?, ?, ?, ?, ?)`
		_, err = db.Exec(query, user.Name, user.Email, user.MobileNo, user.Product, user.Quantity)
		if err != nil {
			log.Println("Error inserting data:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error inserting data into database"})
		}

		return c.JSON(fiber.Map{"message": "User registered and order stored successfully!"})
	})

	log.Println("Server running on port 8000")
	log.Fatal(app.Listen(":8000"))
}
