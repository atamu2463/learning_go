package main

import (
	"net/http"

	"github.com/atamu2463/learning_go/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.ConnectDB()
	database.MigrateDB(db)
	database.CheckDBOrInsertSeed(db)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	router.Run(":8080")
}
