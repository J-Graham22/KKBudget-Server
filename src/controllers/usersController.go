package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(c *gin.Context) {
	var user repository.User

	// Parse JSON body into struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not deserialize json body into user",
		})
		return
	}

	// Hash password
	saltedAndHashedPassword, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Unable to process password - error: %s", err),
		})
		return
	}

	// DB setup
	ctx := context.Background()
	ctx, dbConn, err := db.PrepareContext()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to open database - error: %s", err),
		})
		return
	}
	defer dbConn.Close(ctx)

	repo := repository.New(dbConn)

	// Insert user
	err = repo.AddUser(ctx, repository.AddUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: saltedAndHashedPassword,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to add user - error: %s", err),
		})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Added user %s with email %s", user.Name, user.Email),
	})
}
