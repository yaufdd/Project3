package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yaufdd/project3/internal/handler"
	"github.com/yaufdd/project3/internal/repository"
	"github.com/yaufdd/project3/internal/service"
)

func main() {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal(err)
	}
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("invalid DB_PORT: %v", err)
	}
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		user, pass, host, port, dbname, sslmode,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	boil.SetDB(db)

	r := gin.Default()

	repo := repository.NewRepository(db)
	service := service.NewServise(repo)
	handler := handler.NewHandler(service)

	//User CRUD
	r.POST("user", handler.CreateUser)
	r.GET("user", handler.ReadUser)
	r.PUT("user", handler.UpdateUsername)
	r.DELETE("user", handler.DeleteUser)

	//Company CRUD
	r.POST("company", handler.AddCompanyToFounder)
	r.GET("company", handler.ReadCompany)
	r.PUT("company", handler.UpdateCompanyname)
	r.DELETE("company", handler.DeleteCompany)

	//Project CRUD
	r.POST("project", handler.AddProject)
	r.GET("project", handler.ReadProject)
	r.PUT("project", handler.UpdateProjectTitle)
	r.DELETE("project", handler.DeleteProject)

	//Project post CRUD
	r.POST("project-post", handler.PublishProjectPost)
	r.GET("project-post", handler.ReadProjectPost)
	r.PUT("project-post", handler.UpdateProjectPostDescription)
	r.DELETE("project-post", handler.DeleteProjectPost)

	//TODO: make handler return id for each entity
	r.GET("user-id", handler.GetUserID)
	r.GET("company-id", handler.GetCompanyID)
	r.GET("project-id", handler.GetProjectID)

	//Post intercations
	r.POST("comment", handler.AddComment)
	r.PUT("like-project-post", handler.LikeProjectPost)

	//TODO: donate to project
	//TODO: make test for all handlers

	r.Run(":8080")
}
