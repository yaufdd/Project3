package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yaufdd/project3/internal/handler"
	"github.com/yaufdd/project3/internal/middleware"
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

	repo := repository.NewRepository(db)
	service := service.NewServise(repo)
	handler := handler.NewHandler(service)

	pubBytes, _ := os.ReadFile(os.Getenv("JWT_PUBLIC_PATH"))
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		log.Fatalf("parse pubkey: %v", err)
	}

	jwtmw := middleware.NewJWTMW(pubKey, "project3")

	r := gin.Default()

	r.POST("register", handler.Registration)
	r.POST("auth", handler.Authentication)

	// return id for each entity
	r.GET("user-id", handler.GetUserID)
	r.GET("company-id", handler.GetCompanyID)
	r.GET("project-id", handler.GetProjectID)

	api := r.Group("api")
	api.Use(jwtmw.Handler())
	{
		//Company CRUD
		api.POST("company", handler.AddCompanyToFounder)
		api.GET("company", handler.ReadCompany)
		api.PUT("company", handler.UpdateCompanyname)
		api.DELETE("company", handler.DeleteCompany)

		//Project CRUD
		api.POST("project", handler.AddProject)
		api.GET("project", handler.ReadProject)
		api.PUT("project", handler.UpdateProjectTitle)
		api.DELETE("project", handler.DeleteProject)

		//Project post CRUD
		api.POST("project-post", handler.PublishProjectPost)
		api.GET("project-post", handler.ReadProjectPost)
		api.PUT("project-post", handler.UpdateProjectPostDescription)
		api.DELETE("project-post", handler.DeleteProjectPost)

		//Post intercations
		api.POST("comment", handler.AddComment)
		api.PUT("like-project-post", handler.LikeProjectPost)
	}

	admin := api.Group("/admin")
	admin.Use(jwtmw.RequiredRoles("admin"))
	{
		//TODO: add mw endpoints
		//User CRUD
		admin.POST("user", handler.CreateUser)
		admin.GET("user", handler.ReadUser)
		admin.PUT("user", handler.UpdateUsername)
		admin.DELETE("user", handler.DeleteUser)
	}

	//TODO: donate to project
	//TODO: make test for all handlers

	r.Run(":8080")
}
