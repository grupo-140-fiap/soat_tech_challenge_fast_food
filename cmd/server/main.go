package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db, err := initDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	router := gin.Default()

	routerConfig := routers.RouterConfig{
		Engine: router,
		DB:     db,
	}
	routers.SetupRoutes(routerConfig)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initDatabase() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" {
		dbUser = "root"
	}

	if dbPassword == "" {
		dbPassword = "password"
	}

	if dbHost == "" {
		dbHost = "localhost"
	}

	if dbPort == "" {
		dbPort = "3306"
	}

	if dbName == "" {
		dbName = "fastfood"
	}

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
