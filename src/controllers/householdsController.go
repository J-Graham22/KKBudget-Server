package controllers

import (
	"context"
	_ "errors"
	"fmt"
	_ "fmt"
	"net/http"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

// CREATE
func GetHouseholds(c *gin.Context) {
	ctx := context.Background()
	ctx, dbConn, err := db.PrepareContext()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to open database - err: %s", err),
		})
		return
	}
	defer dbConn.Close(ctx)

	repo := repository.New(dbConn)
	households, err := repo.GetAllHouseholds(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Encountered error when getting households - err: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, households)
}

func AddHousehold(c *gin.Context) {
	householdName := c.Param("name")

	if householdName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name of new household not provided in request",
		})
		return
	}

	ctx := context.Background()
	ctx, dbConn, err := db.PrepareContext()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to open database - err: %s", err),
		})
		return
	}
	defer dbConn.Close(ctx)

	repo := repository.New(dbConn)
	household, err := repo.AddHousehold(ctx, householdName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Encountered error when adding household - err: %s", err),
		})
		return
	}

	// Add the default categories
	var categoryDescription pgtype.Text
	categories := []struct {
		Name        string
		Description string
	}{
		{"Needs", "Use this category for any required expenses that arise, such as groceries and gas"},
		{"Wants", "Use this category for any discretionary expenses that are for enjoyment like going out to eat"},
		{"Unexpected", "Use this category for necessary but unexpected expenses, like a copay after getting sick or irregular car maintenance"},
		{"Cultural", "Use this category for expenses concerning personal cultural learning, such as education or admission to a museum"},
	}

	for _, category := range categories {
		categoryDescription.Scan(category.Description)
		err = repo.AddCategory(
			ctx,
			repository.AddCategoryParams{
				Name:        category.Name,
				Description: categoryDescription,
				HouseholdID: household.ID,
			},
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Encountered error when adding default categories - err: %s", err),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully added the %s household and added default categories", householdName),
	})
}
