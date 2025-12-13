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

func AddUserHouseholdMapping(c *gin.Context) {
	householdId := c.Query("household_id")
	if householdId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "household_id not set",
		})
		return
	}
	householdIdInt, err := strconv.ParseInt(householdId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse household_id into int",
		})
		return
	}

	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user_id not set",
		})
		return
	}
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse user_id into int",
		})
		return
	}

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
	err = repo.AddUserHouseholdPair(
		ctx,
		repository.AddUserHouseholdPairParams{
			HouseholdID: int32(householdIdInt),
			UserID:      int32(userIdInt),
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Encountered exception while attempting to map user to household - error: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully mapped household %d to user %d", householdIdInt, userIdInt),
	})
}

// READ 
func GetUsersByHouseholdId(w http.ResponseWriter, r *http.Request) {
  //TODO
}
func GetHouseholdsByUserId(w http.ResponseWriter, r *http.Request) {
  //TODO
}

// UPDATE is unnecessary here

// DELETE
// use to remove a user from a household
func DeleteUserHouseholdMapping(w http.ResponseWriter, r *http.Request) {
  //TODO
}
// use when deleting a household
