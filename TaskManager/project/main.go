package main

import (
	"log"
	"net/http"
	"project/config"
	"project/internal/handler"
	"project/internal/repository"
	"project/pkg/database"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	cfg := config.NewConfig()
	
	// Подключение к базе данных
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	
	// Инициализация репозитория
	taskRepo := repository.NewTaskRepository(db)
	
	// Инициализация обработчиков
	taskHandler := handler.NewTaskHandler(taskRepo)
	
	// Создание роутера
	router := gin.Default()
	
	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	
	// Группа API v1
	v1 := router.Group("/api/v1")
	{
		// Tasks endpoints (все 5 методов теперь реализованы)
		v1.GET("/tasks", taskHandler.GetTasks)
		v1.POST("/tasks", taskHandler.CreateTask)
		v1.GET("/tasks/:id", taskHandler.GetTask)
		v1.PUT("/tasks/:id", taskHandler.UpdateTask)
		v1.DELETE("/tasks/:id", taskHandler.DeleteTask)
		
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":    "ok",
				"timestamp": time.Now().Unix(),
				"service":   "Task Manager API",
			})
		})
	}
	
	// Существующий ping endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"time":    time.Now().Format(time.RFC3339),
		})
	})
	
	// Запуск сервера
	log.Printf("🚀 Server starting on %s", cfg.HTTPAddr)
	if err := router.Run(cfg.HTTPAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}