package main

import (
	"net/http"

	"github.com/atamu2463/learning_go/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDB()
	database.MigrateDB(database.ConnectDB())
	database.InsertSeedData(database.ConnectDB())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	router.Run(":8080")
}
