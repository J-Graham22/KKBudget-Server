package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
	"github.com/gin-gonic/gin"
)

// CREATE
func AddCategory(c *gin.Context) {
	householdId := c.Param("id")

	if householdId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Household ID not provided in request",
		})
		return
	}

	householdIdInt, err := strconv.ParseInt(householdId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse household_id into int",
		})
		return
	}

	var category repository.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not deserialize JSON body into category",
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
	err = repo.AddCategory(ctx, repository.AddCategoryParams{
		Name:             category.Name,
		HouseholdID:      int32(householdIdInt),
		Description:      category.Description,
		ParentCategoryID: category.ParentCategoryID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Encountered error while trying to add category - err: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully added category %s for household %d", category.Name, householdIdInt),
	})
}

// READ
func GetCategoriesForHousehold(c *gin.Context) {

}

// UPDATE
func UpdateCategoryName(c *gin.Context) {

}

// DELETE
func DeleteCategory(c *gin.Context) {
	//TODO

	//determine behavior for when a category is deleted
	// -- for all transactions that match that category
	// ---- if the category has a parent category, make that the new category
	// ---- if the category has no parent, leave the transactions uncategorized
}
