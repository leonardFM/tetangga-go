package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"tetangga-go/config" // Impor package config
)

func main() {
	// Inisialisasi Sentry dari package config
	config.InitSentry()

	r := gin.Default()

	// Tambahkan middleware Sentry ke Gin
	r.Use(sentrygin.New(sentrygin.Options{}))

	// Endpoint root
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to api tetangga GO",
			"port":    ":8080",
		})
	})

	// Endpoint ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Endpoint user
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ini user dari route baru by Leo"})
	})

	// Endpoint error (testing Sentry)
	r.GET("/error", func(c *gin.Context) {
		err := fmt.Errorf("ini adalah error contoh")
		config.CaptureError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	})

	// Jalankan server
	log.Println("🚀 Server running on port 8080")
	r.Run(":8080")
}
