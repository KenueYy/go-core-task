package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Writer struct {
	ID      uuid.UUID `gorm:"type:uuid; primaryKey; default:gen_random_uuid()"`
	Name    string    `json:"Name" binding:"required"`
	Surname string    `json:"Surname" binding:"required"`
	Email   string    `json:"Email" binding:"required,email" gorm:"uniqueIndex;not null"`
	Posts   []Post    `gorm:"foreignKey:WriterID"`
}

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid; primaryKey; default:gen_random_uuid()"`
	Title     string
	Info      string
	CreatedAt time.Time
	WriterID  uuid.UUID `gorm:"type:uuid"`
}

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	Port       int
}

var (
	DB     *gorm.DB
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)

func main() {
	DB = initDB()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.GET("/writers", getWriters)
	v1.POST("/create/writer", createWriter)
	v1.GET("/posts", getPosts)
	v1.POST("/create/:id/post", createPost)
	v1.GET("/writer/email/:email", getWriterByEmail)
	v1.GET("/writer/id/:id", getWriterById)
	v1.GET("/writer/:id/posts", getWriterPosts)
	r.Run(":8080")
}

func createPost(c *gin.Context) {
	var post Post
	var input struct {
		Title string `json:"Title" binding:"required"`
		Info  string `json:"Info" binding:"required"`
	}

	writerID, _ := uuid.Parse(c.Param("id"))

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	post = Post{
		ID:        uuid.New(),
		Title:     input.Title,
		Info:      input.Info,
		CreatedAt: time.Now(),
		WriterID:  writerID,
	}

	if err := DB.Create(&post).Error; err != nil {
		return
	}

	c.JSON(201, post)

}

func getPosts(c *gin.Context) {
	var posts []Post
	if err := DB.Find(&posts).Error; err != nil {
		return
	}

	c.JSON(http.StatusOK, posts)
}

func getWriterPosts(c *gin.Context) {
	id := c.Param("id")

	writerID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid writer id"})
		return
	}

	var posts []Post
	if err := DB.Where("writer_id = ?", writerID).Find(&posts).Error; err != nil {
		logger.Error("get writer posts failed", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func getWriters(c *gin.Context) {
	var writers []Writer
	if err := DB.Preload("Posts").Find(&writers).Error; err != nil {
		return
	}

	c.JSON(http.StatusOK, writers)
}

func getWriterByEmail(c *gin.Context) {
	var writer Writer
	email := c.Param("Email")

	if err := DB.Preload("Posts").First(&writer, "Email=?", email).Error; err != nil {
		return
	}

	c.JSON(http.StatusOK, writer)
}

func getWriterById(c *gin.Context) {
	var writer Writer
	id := c.Param("id")

	if err := DB.Preload("Posts").First(&writer, "id=?", id).Error; err != nil {
		return
	}

	c.JSON(http.StatusOK, writer)
}

func createWriter(c *gin.Context) {
	var writer Writer
	var input struct {
		Name    string `json:"Name" binding:"required"`
		Surname string `json:"Surname" binding:"required"`
		Email   string `json:"Email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	writer = Writer{
		ID:      uuid.New(),
		Name:    input.Name,
		Surname: input.Surname,
		Email:   input.Email,
	}

	if err := DB.Create(&writer).Error; err != nil {
		return
	}

	c.JSON(201, writer)

}

func initDB() *gorm.DB {
	cfg := &Config{
		DBHost:     "localhost",
		DBPort:     5433,
		DBUser:     "postgres",
		DBPassword: "password",
		DBName:     "walletdb_test",
		DBSSLMode:  "disable",
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to database connection",
			"host", cfg.DBHost,
			"port", cfg.DBPort,
			"db_name", cfg.DBName,
			"error", err.Error(),
		)
	}

	if err := db.AutoMigrate(Writer{}); err != nil {
		logger.Error("migration failed",
			"db_name", cfg.DBName,
			"error", err.Error(),
		)
	}

	if err := db.AutoMigrate(Post{}); err != nil {
		logger.Error("migration failed",
			"db_name", cfg.DBName,
			"error", err.Error(),
		)
	}

	logger.Info("database initialized",
		"db_name", cfg.DBName,
	)

	return db
}
