package controllers

import (
	"adib2/praktek2/connection"
	"adib2/praktek2/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

var books []models.Book

func AddBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbName := "praktek2"
	collectionName := "adibpraktek2"

	client, err := connection.ConnectToMongoDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal terhubung ke database"})
		return
	}
	defer client.Disconnect(nil)

	collection := client.Database(dbName).Collection(collectionName)
	_, err = collection.InsertOne(context.TODO(), newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan buku ke database"})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}
