package main

import (
	"log"
	"net/http"
	_ "time"

  "github.com/gin-gonic/gin"
  "github.com/rs/cors"
	_ "github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/J-Graham22/BudgetBuddyServer/src/controllers"
	_ "github.com/J-Graham22/BudgetBuddyServer/src/db"
	_ "github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
)

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Welcome to Budget Buddy Server!",
    })
  })

  //households
  // router.HandleFunc("POST /households/{name}", controllers.AddHousehold)
  r.POST("/households/:name", controllers.AddHousehold)

  //users
  // router.HandleFunc("POST /users", controllers.AddUser)
  r.POST("/users", controllers.AddUser)

  //categories
  // router.HandleFunc("POST /categories/{id}", controllers.AddCategory)
  r.POST("/categories/:id", controllers.AddCategory)
  // router.HandleFunc("GET /categories/{id}", controllers.GetCategoriesForHousehold)
  r.GET("/categories/:id", controllers.GetCategoriesForHousehold)

  //transactions
  // router.HandleFunc("GET /transactions/{id}", controllers.GetTransactionsForHousehold)
  r.GET("/transactions/:id", controllers.GetTransactionsForHousehold)

  log.Println("Starting server on port :8080")
  r.Run(":8080")
}
