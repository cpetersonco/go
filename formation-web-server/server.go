package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := startDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	startWebServer(db)
}

func startDatabaseConnection() (*sql.DB, error) {
	connStr := "postgres://postgres:docker@localhost/user_tables?sslmode=disable"
	return sql.Open("postgres", connStr)
}

func startWebServer(db *sql.DB) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		table_name := c.PostForm("table_name")

		fmt.Printf("id: %s; page: %s; table: %s", id, page, table_name)

		createTable(table_name, db)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func createTable(table_name string, db *sql.D) {
	sanitized := pq.QuoteIdentifier(table_name)

	// Create a table.
	_, err := db.Exec(fmt.Sprintf("CREATE TABLE %s (id serial PRIMARY KEY)", sanitized))
	if err != nil {
		panic(err)
	}
}
