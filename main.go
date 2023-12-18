package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"personal/sika/internal/database"
	"personal/sika/internal/database/seeders"
	"personal/sika/internal/routes"
	"time"
)

func main() {
	// TODO: Create a Config object that reads from an .ENV file

	// Creating and Passing DB Connection to each route for better testability (Dependency Injection)
	db := database.InitDB(database.DBConfig{
		Username:       "root",
		Password:       "secret",
		DBName:         "sika",
		MaxConLifetime: time.Minute * 3,
		MaxOpenIdleCon: 200,
	})

	if len(os.Args) != 2 {
		fmt.Println("Please specify serve or migrate argument.")
		return
	}

	switch os.Args[1] {
	case "serve":
		runServer(db)
	case "seed":
		fmt.Println("Seeding database, please wait...")
		seeders.SeedUsers(db)
		fmt.Println("Seeder ran successfully.")
	}
}

func runServer(db *sql.DB) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	users := r.Group("/users")
	{
		users.GET("/", func(context *gin.Context) {
			routes.GetUsers(context, db)
		})
		users.GET("/:id", func(context *gin.Context) {
			routes.GetUser(context, db)
		})
	}
	fmt.Printf("Web server is listening on port 8020...")
	err := r.Run(":8020")
	if err != nil {
		panic(err)
	}
}
