package main

import (
	"context"
	"errors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ilmedova/go-url-shortener/internal/handlers"
	"github.com/ilmedova/go-url-shortener/internal/repositories"
	"github.com/ilmedova/go-url-shortener/internal/services"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "github.com/ilmedova/go-url-shortener/docs"
	_ "github.com/lib/pq"
	_ "github.com/swaggo/http-swagger"
	_ "log"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found, using system environment variables")
	}
}

func InitDB() *sqlx.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	dbDriver := os.Getenv("DATABASE_DRIVER")

	db, err := sqlx.Connect(dbDriver, dbUrl)
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %v", err)
	}

	logrus.Info("Connected to PostgreSQL database")
	return db
}

func InitRedis() *redis.Client {
	redisAddr := os.Getenv("REDIS_URL")

	client := redis.NewClient(&redis.Options{Addr: redisAddr})
	logrus.Info("Connected to Redis")

	return client
}

// @title URL Shortener API
// @version 1.0
// @description This is a URL shortener API built with Golang and Gin.
// @termsOfService http://swagger.io/terms/
// @contact.name Mahri Ilmedova
// @contact.email ilmedovamahri2@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /
func main() {
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	LoadEnv()

	db := InitDB()
	redisClient := InitRedis()

	repo := repositories.NewURLRepository(db)
	cache := repositories.NewCacheRepository(redisClient)
	service := services.NewURLService(repo, cache)
	handler := handlers.NewURLHandler(service)

	r := gin.Default()
	r.POST("/shorten", handler.ShortenURL)
	r.GET("/:short", handler.ResolveURL)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		logrus.Info("Server running on port 8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server forced to shutdown: %v", err)
	}

	logrus.Info("Server exited properly")
}
