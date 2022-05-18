package controllers

import (
	"strconv"

	"github.com/G-Basilio/goApiRest/database"
	"github.com/G-Basilio/goApiRest/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID tem que ser um inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Livro não encontrado: " + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Sem ligação com JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Sem ligação com book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {

	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Sem ligação com  list book: " + err.Error(),
		})
		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Sem ligação com JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "	Livro não atualizado: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID tem que ser um inteiro",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "	Livro não deletado: " + err.Error(),
		})
		return
	}

	c.Status(204)

}
