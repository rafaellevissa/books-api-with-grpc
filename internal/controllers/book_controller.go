package controllers

import (
	"crud/internal/database"
	"crud/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON book: " + err.Error(),
		})
		return
	}

	//Criando
	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func ShowAllBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find books: " + err.Error(),
		})
		return
	}

	c.JSON(200, books)
}

func ShowBook(c *gin.Context) {
	bookId := c.Param("id")

	//Transformando para inteiro
	newBookId, err := strconv.Atoi(bookId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, newBookId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func EditBook(c *gin.Context) {
	bookId := c.Param("id")

	//Transformando para inteiro
	newBookId, err := strconv.Atoi(bookId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, newBookId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book by id: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	//Transformando para inteiro
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
