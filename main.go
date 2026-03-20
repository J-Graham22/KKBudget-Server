package main

import (
	"log"
	"net/http"
	_ "time"

	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/rs/cors"

	"github.com/J-Graham22/BudgetBuddyServer/src/controllers"
	_ "github.com/J-Graham22/BudgetBuddyServer/src/db"
	_ "github.com/J-Graham22/BudgetBuddyServer/src/db/repository"
)

r := gin.Default()

func main() {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Budget Buddy Server!",
		})
	})

	addCRUDEndpoints()

	log.Println("Starting server on port :8080")
	r.Run(":8080")
}

func addCRUDEndpoints() {
	addHouseholdCRUD()
	addUserCRUD()
	addCategoryCRUD()
	addTransactionCRUD()
}

func addHouseholdCRUD() {
	// router.HandleFunc("POST /households/{name}", controllers.AddHousehold)
	r.POST("/households/:name", controllers.AddHousehold)
}

func addUserCRUD() {
	// router.HandleFunc("POST /users", controllers.AddUser)
	r.POST("/users", controllers.AddUser)
}

func addCategoryCRUD() {
	// router.HandleFunc("POST /categories/{id}", controllers.AddCategory)
	r.POST("/categories/:id", controllers.AddCategory)
	// router.HandleFunc("GET /categories/{id}", controllers.GetCategoriesForHousehold)
	r.GET("/categories/:id", controllers.GetCategoriesForHousehold)
}

func addTransactionCRUD() {
	// router.HandleFunc("GET /transactions/{id}", controllers.GetTransactionsForHousehold)
	r.GET("/transactions/:id", controllers.GetTransactionsForHousehold)
}
