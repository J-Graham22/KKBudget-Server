package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/J-Graham22/BudgetBuddyServer/src/db"
	"github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
	"github.com/gin-gonic/gin"
)

func AddTransaction() {
}

func UpdateTransaction() {
}

func GetTransactionsByBudget() {
}

func GetTransactionsForHousehold(c *gin.Context) {
  ctx := context.Background()
  ctx, dbConn, err := db.PrepareContext()
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": fmt.Sprintf("Encountered error while trying to open database - error: %s", err),
    })
    return
  }
  defer dbConn.Close(ctx)

  householdID := c.Param("id")
  if householdID == "" {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Household ID not provided in request",
    })
    return
  }

  repo := repository.New(dbConn)
  transactions, err := repo.GetTransactionsByHousehold(ctx, 0) // Replace `0` with parsed household ID if needed
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": fmt.Sprintf("Encountered error while retrieving transactions - error: %s", err),
    })
    return
  }

  c.JSON(http.StatusOK, transactions)
}

func GetTransactionsByCategory() {

}

func GetTransactionsByUser() {

}
