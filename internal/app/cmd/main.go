package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yaufdd/project3/internal/cache"
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", os.Getenv("REDIS_PORT")),
		Password: fmt.Sprintf("localhost:%s", os.Getenv("REDIS_PASSWORD")),
	})

	repo := repository.NewRepository(db)
	redis := cache.NewCache(rdb)
	service := service.NewServise(repo, redis)
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
	r.POST("refresh", handler.RefreshTokenAuth)

	// return id for each entity
	r.GET("user-id", handler.GetUserID)
	r.GET("company-id", handler.GetCompanyID)
	r.GET("project-id", handler.GetProjectID)

	api := r.Group("api")
	api.Use(jwtmw.Handler())
	{
		api.GET("company", handler.ReadCompany)

		api.GET("project", handler.ReadProject)

		api.GET("project-post", handler.ReadProjectPost)

		api.GET("user", handler.ReadUser)

		//Post intercations
		api.POST("comment", handler.AddComment)
		api.PUT("like-project-post", handler.LikeProjectPost)
	}

	me := api.Group("/me")
	me.Use()
	{
		me.PUT("user", handler.UpdateUsername)

		me.POST("company", handler.AddCompanyToFounder)
		me.PUT("company", handler.UpdateCompanyName)
		me.DELETE("company", handler.DeleteCompany)

		me.POST("project", handler.AddProject)
		me.PUT("project", handler.UpdateProjectTitle)
		me.DELETE("project", handler.DeleteProject)

		me.POST("project-post", handler.PublishProjectPost)
		me.PUT("project-post", handler.UpdateProjectPostDescription)
		me.DELETE("project-post", handler.DeleteProjectPost)
	}

	admin := api.Group("/admin")
	admin.Use(jwtmw.RequiredRoles("admin"))
	{
		admin.POST("user", handler.CreateUser)
		admin.DELETE("user", handler.DeleteUser)
	}

	//TODO: donate to project
	//TODO: make test for all handlers

	r.Run(":8080")
}
