package controllers

import (
	"fmt"
	"net/http"
	"net/smtp"

	"wese/demo/models"
	"wese/demo/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Others(c *gin.Context) {

	// log.Println(users)
	// go services.LogToSlack("This is run in a separate process")
	fmt.Println("SMTP login test")

	// Sender data.
	from := "wese@hotel.inc.ug"
	password := ""

	// Receiver email address.
	to := []string{
		"rutatiina@riginem.org",
	}

	// smtp server configuration.
	smtpHost := ""
	smtpPort := ""

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

	c.JSON(http.StatusOK, gin.H{
		"message": "Users",
		"payload": "okay",
	})
}

// GET /books
// Find all books
func List(c *gin.Context) {
	var models []models.User
	// db.Find(&users)
	services.DB.Scopes(services.Paginate(c.Request)).Find(&models)

	// log.Println(users)

	// // go services.LogToSlack("This is run in a separate process")

	c.JSON(http.StatusOK, gin.H{
		"message": "Users",
		"payload": models,
	})
}

// GET /books/:id
// Find a book
func Show(c *gin.Context) {
	// Get model if exist
	var model models.User
	if err := services.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User",
		"payload": model,
	})
}

// POST /books
// Create new book
func Store(c *gin.Context) {
	/*
		// Validate input
		var input UserNew
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create book
		book := models.User{Title: input.Title, Author: input.Author}
		services.DB.Create(&book)

		c.JSON(http.StatusOK, gin.H{"data": book})
	*/

	//this link explains why we has to use form JSON
	//https://gin-gonic.com/docs/examples/bind-body-into-dirrerent-structs/

	// Validate input
	var input models.UserNew
	if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var webUser models.User
	c.ShouldBindBodyWith(&webUser, binding.JSON)
	services.DB.Create(&webUser)
	c.JSON(http.StatusOK, gin.H{
		"message": "User saved",
		"payload": webUser,
	})
}

// PATCH /books/:id
// Update
func Update(c *gin.Context) {
	// Get model if exist
	var model models.User
	if err := services.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UserUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services.DB.Model(&model).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": model})
}

// DELETE /books/:id
// Delete a book
func Destroy(c *gin.Context) {
	// Get model if exist
	var model models.User
	if err := services.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	services.DB.Delete(&model)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// PERMANENTLY DELETE /books/:id/void
// Permanently Delete a record
func VoidDestroyed(c *gin.Context) {
	// Get model if exist
	var model models.User
	if err := services.DB.Unscoped().Where("id = ?", c.Param("id")).First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if !model.DeletedAt.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found. Only deleted items can be voided"})
		return
	}

	//permanently delete record
	services.DB.Unscoped().Delete(&model)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
